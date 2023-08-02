#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <fcntl.h>
#include <stdbool.h>
#include <signal.h>
#include "shell.h"
#define MAX_HISTORY_SIZE 1024
#define MAX_INPUT_SIZE 1024
#define MAX_NUM_TOKENS 1024

char history[MAX_HISTORY_SIZE][MAX_INPUT_SIZE];
int historyCount = 0;

void powerline() {
    // Get the current working directory
    char cwd[1024];
    if (getcwd(cwd, sizeof(cwd)) != NULL) {
        // Find the last directory in the path
        char* last_dir = strrchr(cwd, '/');
        if (last_dir != NULL) {
            last_dir++; // Move past the slash to get the last directory name
        } else {
            last_dir = cwd; // If there's no slash, use the entire path
        }

        // Print the Powerline-style prompt
        printf("\033[1;34m%s \033[1;32m➜\033[0m ", last_dir);
    } else {
        // Print a simple prompt if getting the current directory fails
        printf("\n\033[1;32m[➜]\033[0m ");
    }
}

void handleSignal(int signal) {
    if (signal == SIGINT)
        printf("Shell > ");
}

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

void addToHistory(const char *input) {
    if (historyCount == MAX_HISTORY_SIZE) {
        // Shift the history entries to make room for the new entry
        for (int i = 1; i < MAX_HISTORY_SIZE; i++) {
            strcpy(history[i - 1], history[i]);
        }
        strcpy(history[MAX_HISTORY_SIZE - 1], input);
    } else {
        strcpy(history[historyCount], input);
        historyCount++;
    }
}

void printHistory(void) {
    for (int i = 0; i < historyCount; i++) {
        printf("\t%d.%s\n", i + 1, history[i]);
    }
}

int shell(void) {
    char input[MAX_INPUT_SIZE];
    char *tokens[MAX_NUM_TOKENS];
    bool running = true;

    struct sigaction sa;
    memset(&sa, 0, sizeof(sa));
    sa.sa_handler = handleSignal;
    sigaction(SIGINT, &sa, NULL);

    while (running) {
        powerline();
        if (fgets(input, sizeof(input), stdin) == NULL) {
            break;
        }

        // Remove newline character if present
        input[strcspn(input, "\n")] = '\0';

        // Check for empty input
        if (strlen(input) == 0) {
            continue;
        }

        // Tokenize input
        int tokenCount = tokenizeInput(input, tokens);

        // Handle built-in commands
        if (strcmp(tokens[0], "exit") == 0) {
            running = false;
        } else if (strcmp(tokens[0], "ping") == 0) {
            if (tokenCount < 2) {
                printf("Usage: ping <destination IP address>\n");
            } else {
                char command[MAX_INPUT_SIZE];
                snprintf(command, sizeof(command), "./ping %s", tokens[1]);
                int result = system(command);
                if (result == -1) {
                    perror("system");
                }
            }
        } else if (strcmp(tokens[0], "scanner") == 0) {
            if (tokenCount < 4) {
                printf("Usage: scanner <IP Address> <startPort> <endPort>\n");
            } else {
                char command[MAX_INPUT_SIZE];
                snprintf(command, sizeof(command), "./scanner %s %d %d", tokens[1], tokens[2], tokens[3]);
                int result = system(command);
                if (result == -1) {
                    perror("system");
                }
            }
        } else if (strcmp(tokens[0], "dnslookup") == 0) {
            if (tokenCount < 2) {
                printf("Usage: DNS <domain>\n");
            } else {
                char command[MAX_INPUT_SIZE];
                snprintf(command, sizeof(command), "./dns %s", tokens[1]);
                int result = system(command);
                if (result == -1) {
                    perror("system");
                }
            }
        } else if (strcmp(tokens[0], "subnet") == 0) {
            system("./dist/subnet");
        } else if (strcmp(tokens[0], "whois") == 0) {
            char command[MAX_INPUT_SIZE];
            snprintf(command, sizeof(command), "./whois %s", tokens[1]);
            int result = system(command);
            if (result == -1) {
                perror("system");
            }
        } else if (strcmp(tokens[0], "dirb") == 0) {
            char command[MAX_INPUT_SIZE];
            snprintf(command, sizeof(command), "./drib %s", tokens[1]);
            int result = system(command);
            if (result == -1) {
                perror("system");
            }
        } else if (strcmp(tokens[0], "sniffer") == 0) {
            char command[MAX_INPUT_SIZE];
            snprintf(command, sizeof(command), "./sniffer");
            int result = system(command);
            if (result == -1) {
                perror("system");
            }
        } else if (strcmp(tokens[0], "history") == 0) {
            printHistory();
        } else {
            // Handle command execution
            addToHistory(input);
            bool background = false;

            // Check if the last token is '&'
            char *last_token = tokens[tokenCount - 1];
            if (strcmp(last_token, "&") == 0) {
                background = true;
                *last_token = '\0'; // Remove the '&' from tokens
            }

            // Check if input/output redirection is required
            int fd_in = 0, fd_out = 1;
            for (int i = 0; tokens[i] != NULL; i++) {
                if (strcmp(tokens[i], "<") == 0) {
                    tokens[i] = NULL;
                    fd_in = open(tokens[i + 1], O_RDONLY);
                    if (fd_in == -1) {
                        perror("open");
                        break;
                    }
                    dup2(fd_in, STDIN_FILENO);
                    close(fd_in);
                } else if (strcmp(tokens[i], ">") == 0) {
                    tokens[i] = NULL;
                    fd_out = open(tokens[i + 1], O_WRONLY | O_CREAT | O_TRUNC, 0644);
                    if (fd_out == -1) {
                        perror("open");
                        break;
                    }
                    dup2(fd_out, STDOUT_FILENO);
                    close(fd_out);
                }
            }
            executeCommand(tokens, background);
        }
    }

    return 0;
}
