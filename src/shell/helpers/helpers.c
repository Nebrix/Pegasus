#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <fcntl.h>
#include <stdbool.h>
#include <signal.h>

#define MAX_HISTORY_SIZE 1024
#define MAX_INPUT_SIZE 1024
#define MAX_NUM_TOKENS 1024

void handleSignal(int signal) {
    if (signal == SIGINT)
        printf("Shell > ");
}

char* getPowerlinePrompt() {
    FILE* powerlineScript = popen("perl src/shell/powerline/main.pl", "r");
    if (powerlineScript == NULL) {
        perror("popen");
        exit(EXIT_FAILURE);
    }

    char* prompt = (char*)malloc(MAX_INPUT_SIZE);
    if (fgets(prompt, MAX_INPUT_SIZE, powerlineScript) == NULL) {
        perror("fgets");
        exit(EXIT_FAILURE);
    }

    pclose(powerlineScript);
    prompt[strcspn(prompt, "\n")] = '\0'; // Remove newline character
    return prompt;
}