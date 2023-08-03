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
#include "../help/help.h"
#include "../ascii/ascii.h"
#include "history/history.h"
#include "command/command.h"
#include "helpers/helpers.h"

#define MAX_HISTORY_SIZE 1024
#define MAX_INPUT_SIZE 1024
#define MAX_NUM_TOKENS 1024

int shell(void) {
    char input[MAX_INPUT_SIZE];
    char *tokens[MAX_NUM_TOKENS];
    bool running = true;

    struct sigaction sa;
    memset(&sa, 0, sizeof(sa));
    sa.sa_handler = handleSignal;
    sigaction(SIGINT, &sa, NULL);

    system("clear");
    ascii();
    while (running) {
        char* prompt = getPowerlinePrompt();
        printf("%s", prompt);
        free(prompt);
        
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
        } else if (strcmp(tokens[0], "reload") == 0) {
            system("clear");
            ascii();
        } else if (strcmp(tokens[0], "help") == 0) {
            Help();
        } else if (strcmp(tokens[0], "iplookup") == 0) {
            system("./dist/ip");
        } else if (strcmp(tokens[0], "ping") == 0) {
            if (tokenCount < 2) {
                pingHelp();
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
                scannerHelp();
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
                dnslookupHelp();
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
            if (tokenCount > 2) {
                whoisHelp();
            } else {
                char command[MAX_INPUT_SIZE];
                snprintf(command, sizeof(command), "./whois %s", tokens[1]);
                int result = system(command);
                if (result == -1) {
                    perror("system");
                }
            }
        } else if (strcmp(tokens[0], "dirb") == 0) {
            if (tokenCount > 2) {
                dirbHelp();
            } else {
                char command[MAX_INPUT_SIZE];
                snprintf(command, sizeof(command), "./dirb %s", tokens[1]);
                int result = system(command);
                if (result == -1) {
                    perror("system");
                }
            }
        } else if (strcmp(tokens[0], "sniffer") == 0) {
            system("./dist/packet-sniffer");
        } else if (strcmp(tokens[0], "echoln") == 0) {
            echolnCommand(tokens);
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