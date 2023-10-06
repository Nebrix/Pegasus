#include <stdio.h>
#include <stdlib.h>
#include <netdb.h>
#include <arpa/inet.h>
#include <string.h>
#include <resolv.h>

int main(int argc, char *argv[]) {
    if (argc != 2) {
        exit(1);
    }

    char *domain = argv[1];

    // Perform DNS lookup
    struct hostent *host;
    host = gethostbyname(domain);
    if (host == NULL) {
        printf("Error: Unable to resolve host\n");
        exit(1);
    }

    // Print IP addresses
    printf("IP addresses for %s:\n", domain);
    char **ip;
    for (ip = host->h_addr_list; *ip != NULL; ip++) {
        struct in_addr addr;
        memcpy(&addr, *ip, sizeof(struct in_addr));
        printf("%s\n", inet_ntoa(addr));
    }

    // Perform DNS lookup for other record types
    printf("\nOther DNS records for %s:\n", domain);

    // NS records
    struct hostent *nsHost;
    nsHost = gethostbyname(domain);
    if (nsHost != NULL && nsHost->h_addrtype == AF_INET) {
        char **ns;
        for (ns = nsHost->h_aliases; *ns != NULL; ns++) {
            printf("NS: %s\n", *ns);
        }
    }

    // MX records
    char mxHost[256];
    int mxPreference;
    if (res_query(domain, C_IN, T_MX, (unsigned char *)&mxHost, sizeof(mxHost)) != -1) {
        unsigned char *p = (unsigned char *)&mxHost;
        p += 12; // Skip the header
        while (*p != 0) {
            int len = *p;
            printf("MX (Raw): ");
            for (int i = 0; i < len + 2; i++) {
                printf("%02X ", *(p + i));
            }
            printf("\n");
            printf("Preference: %d\n", *(p + len + 3));
            p += len + 4;
        }
    }

    // SOA record (extracted from TXT records)
    struct hostent *soaHost;
    soaHost = gethostbyname(domain);
    if (soaHost != NULL && soaHost->h_addrtype == AF_INET) {
        char **txt;
        for (txt = soaHost->h_aliases; *txt != NULL; txt++) {
            if (strncmp(*txt, "v=spf1", 6) == 0) {
                char *soaParts[7];
                int i = 0;
                char *token = strtok(*txt, " ");
                while (token != NULL && i < 7) {
                    soaParts[i] = token;
                    token = strtok(NULL, " ");
                    i++;
                }
                if (i >= 7) {
                    printf("SOA: Primary NS: %s, Responsible person's mailbox: %s\n", soaParts[6], soaParts[1]);
                }
                break;
            }
        }
    }

    // TXT records
    struct hostent *txtHost;
    txtHost = gethostbyname(domain);
    if (txtHost != NULL && txtHost->h_addrtype == AF_INET) {
        char **txt;
        for (txt = txtHost->h_aliases; *txt != NULL; txt++) {
            printf("TXT: %s\n", *txt);
        }
    }

    return 0;
}
