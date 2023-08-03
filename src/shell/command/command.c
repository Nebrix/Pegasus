#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <fcntl.h>
#include <string.h>
#include <stdbool.h>
#include "command.h"

#define MAX_HISTORY_SIZE 1024
#define MAX_INPUT_SIZE 1024
#define MAX_NUM_TOKENS 1024

void executeCommand(char **tokens, bool background) {
    pid_t pid = fork();

    if (pid == -1) {
        perror("fork");
        return;
    } else if (pid == 0) {
        // Child Process
        if (background) 
            setpgid(0, 0); // Set the new child process in a new process group for background execution
        execvp(tokens[0], tokens);
        perror("execvp");
        exit(EXIT_FAILURE);
    } else {
        // Parent Process
        if (!background) {
            int status;
            waitpid(pid, &status, 0);
        } else {
            printf("Process %d running in the background.\n", pid);
        }
    }
}

int tokenizeInput(char *input, char **tokens) {
    int tokenCount = 0;
    char *token = strtok(input, " \t\n");

    while (token != NULL && tokenCount < MAX_NUM_TOKENS) {
        tokens[tokenCount] = token;
        tokenCount++;
        token = strtok(NULL, " \t\n");
    }
    tokens[tokenCount] = NULL;
    return tokenCount;
}