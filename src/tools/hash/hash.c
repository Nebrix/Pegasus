#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <openssl/md5.h>
#include <openssl/sha.h>

char* hashString(const char* input, const char* algorithm) {
    size_t outputSize = 0;

    if (strcmp(algorithm, "md5") == 0) {
        outputSize = MD5_DIGEST_LENGTH;
    } else if (strcmp(algorithm, "sha1") == 0) {
        outputSize = SHA_DIGEST_LENGTH;
    } else if (strcmp(algorithm, "sha256") == 0) {
        outputSize = SHA256_DIGEST_LENGTH;
    } else {
        return "Error: Unsupported algorithm";
    }

    unsigned char* hash = malloc(outputSize);
    if (hash == NULL) {
        perror("Failed to allocate memory");
        exit(1);
    }

    if (strcmp(algorithm, "md5") == 0) {
        MD5((unsigned char*)input, strlen(input), hash);
    } else if (strcmp(algorithm, "sha1") == 0) {
        SHA1((unsigned char*)input, strlen(input), hash);
    } else if (strcmp(algorithm, "sha256") == 0) {
        SHA256((unsigned char*)input, strlen(input), hash);
    }

    // Convert the hash to a hexadecimal string
    char* hexHash = malloc(outputSize * 2 + 1);
    if (hexHash == NULL) {
        perror("Failed to allocate memory");
        free(hash);
        exit(1);
    }

    for (size_t i = 0; i < outputSize; i++) {
        snprintf(hexHash + (i * 2), 3, "%02x", hash[i]);
    }

    free(hash);
    return hexHash;
}

int main(int argc, char* argv[]) {
    if (argc != 3) {
        printf("Usage: %s <string> <algorithm>\n", argv[0]);
        return 1;
    }

    const char* input = argv[1];
    const char* algorithm = argv[2];

    char* hashOutput = hashString(input, algorithm);

    printf("%s\n", hashOutput);

    free(hashOutput);

    return 0;
}
