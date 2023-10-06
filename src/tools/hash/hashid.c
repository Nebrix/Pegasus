#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <openssl/md5.h>
#include <openssl/sha.h>

const char* identifyHash(const char* hashValue) {
    size_t hashLen = strlen(hashValue);

    if (hashLen == MD5_DIGEST_LENGTH * 2) {
        return "MD5";
    } else if (hashLen == SHA_DIGEST_LENGTH * 2) {
        return "SHA-1";
    } else if (hashLen == SHA256_DIGEST_LENGTH * 2) {
        return "SHA-256";
    } else {
        return "Unknown";
    }
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        return 1;
    }

    const char* hashValue = argv[1];

    const char* hashAlgorithm = identifyHash(hashValue);
    printf("Hash Algorithm: %s\n", hashAlgorithm);

    return 0; 
}