#include <stdio.h>
#include <unistd.h>
#include <sys/utsname.h>
#include <string.h>
#include "ascii.h"

char *version() {
    FILE *file = fopen(".version", "r");

    if (file == NULL) {
        perror("Error");
        return "Unknown";
    }

    static char version[20];  // Make it static to ensure it survives function scope
    if (fscanf(file, "VERSION=%s", version) != 1) {
        fclose(file);
        perror("Error reading version");
        return "Unknown";
    }
    fclose(file);

    return version;
}

void get_distribution_name(char *distro, size_t distro_size) {
    FILE *fp = fopen("/etc/os-release", "r");
    if (fp != NULL) {
        char line[256];
        while (fgets(line, sizeof(line), fp)) {
            if (strncmp(line, "NAME=", 5) == 0) {
                // Extract the distribution name from the line
                char *name_start = strchr(line, '=');
                if (name_start != NULL) {
                    name_start++; // Move to the start of the actual name

                    // Check if the name starts with a quotation mark
                    if (*name_start == '"') {
                        name_start++; // Move past the quotation mark
                        char *name_end = strchr(name_start, '"');
                        if (name_end != NULL) {
                            // Calculate the length of the name
                            size_t name_length = name_end - name_start;
                            strncpy(distro, name_start, name_length);
                            distro[name_length] = '\0'; // Null-terminate the string
                            fclose(fp);
                            return;
                        }
                    } else {
                        // The name doesn't start with a quotation mark
                        size_t name_length = strcspn(name_start, "\n");
                        strncpy(distro, name_start, name_length);
                        distro[name_length] = '\0'; // Null-terminate the string
                        fclose(fp);
                        return;
                    }
                }
            }
        }
        fclose(fp);
    }
    // If the distribution name cannot be found, use "Unknown"
    strncpy(distro, "Unknown", distro_size);
}

void ascii() {
    char *username = getlogin();

    char distro[256];
    get_distribution_name(distro, sizeof(distro));

    puts("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⡀⠀⠀⠀⠀⠀");
    puts("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣧⠃⡃⠀⠀⠀⠀⠀");
    puts("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⢿⠇⠏⡇⠀⠀⠀⠀");
    puts("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣟⠏⡜⠰⢣⠀⠀⠀⠀");
    puts("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠟⣡⠊⡠⢁⣎⠀⠀⠀⠀");
    puts("⠀⠀⢀⣶⣰⡂⠀⠀⠀⠀⣠⠖⠋⡔⠊⡠⢏⡠⢃⡜⠀⠀⠀⠀");
    puts("⠀⠀⡧⣩⠾⢹⣓⣤⠀⠘⣯⡀⠀⣗⣋⣤⢃⠔⢩⠆⠀⠀⠀⠀");
    puts("⠀⣸⣅⠁⢁⡞⠨⡍⢷⢂⡾⠁⣰⠡⠤⣊⠥⠒⠁⠀⠀⠀⠀⠀");
    puts("⠘⠿⠽⠞⢫⠀⠀⠀⢻⠟⣡⡾⠱⡑⠦⠥⣀⣀⣀⠀⠀⠀⠀⠀");
    puts("⠀⠀⠀⠀⡎⢀⠐⠀⠁⠈⠺⠷⠣⠃⠁⡀⠈⢿⡇⠨⢢⠀⠀⠀");
    puts("⠀⠀⠀⢀⡃⠀⠐⠀⠂⠐⠂⢀⠂⠂⠐⢀⠀⡾⢔⢄⠡⣃⠀⠀");
    puts("⠀⠀⣰⢏⡵⡄⢲⠶⠦⠴⠤⠦⠶⣶⠤⣀⡀⢧⡀⠉⠘⠚⠓⠂");
    printf("⠀⢸⡏⠃⠀⠙⣞⡆⠀⠀⠀⠀⠀⠀⣵⡚⠱⣎⣇⠀⠀   username: %s\n", username);
    printf("⠀⣿⣷⠄⠀⠀⢹⣽⠀⠀⠀⠀⠀⣼⣳⠃⠀⠹⣾⣤⠀   distro: %s\n", distro);
    printf("⠀⠉⠁⠀⠀⠰⠿⠟⠀⠀⠀⠀⠘⠛⠛⠀⠀⠀⠹⠋    version: %s\n", version());
}