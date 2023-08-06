#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdbool.h>
#include <signal.h>

#include "shell.h"
#include "../help/help.h"
#include "../ascii/ascii.h"
#include "history/history.h"
#include "command/command.h"
#include "helpers/helpers.h"
#include "../core-util/core.h"

#define MAX_HISTORY_SIZE 1024
#define MAX_INPUT_SIZE 1024
#define MAX_NUM_TOKENS 1024

pid_t startServer() {
    pid_t pid = fork();
    if (pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (pid == 0) {
        execlp("nohup ./server &", "server", NULL);
        perror("exec");
        exit(EXIT_FAILURE);
    } else {
        return pid;
    }
}

void killServer(pid_t pid) {
    if (kill(pid, SIGKILL) == -1) {
        perror("kill");
    } 
}

typedef enum {
    CMD_EXIT,
    CMD_RELOAD,
    CMD_HELP,
    CMD_IPLOOKUP,
    CMD_PING,
    CMD_SCANNER,
    CMD_DNSLOOKUP,
    CMD_SUBNET,
    CMD_WHOIS,
    CMD_DIRB,
    CMD_SNIFFER,
    CMD_HASHIDENT,
    CMD_HASH,
    CMD_LS,
    CMD_ECHOLN,
    CMD_HISTORY,
    CMD_SERVER,
    CMD_UNKNOWN
} CommandType;

CommandType getCommandType(const char* command) {
    if (strcmp(command, "exit") == 0) return CMD_EXIT;
    if (strcmp(command, "reload") == 0) return CMD_RELOAD;
    if (strcmp(command, "help") == 0) return CMD_HELP;
    if (strcmp(command, "iplookup") == 0) return CMD_IPLOOKUP;
    if (strcmp(command, "ping") == 0) return CMD_PING;
    if (strcmp(command, "scanner") == 0) return CMD_SCANNER;
    if (strcmp(command, "dnslookup") == 0) return CMD_DNSLOOKUP;
    if (strcmp(command, "subnet") == 0) return CMD_SUBNET;
    if (strcmp(command, "whois") == 0) return CMD_WHOIS;
    if (strcmp(command, "dirb") == 0) return CMD_DIRB;
    if (strcmp(command, "sniffer") == 0) return CMD_SNIFFER;
    if (strcmp(command, "hashident") == 0) return CMD_HASHIDENT;
    if (strcmp(command, "hash") == 0) return CMD_HASH;
    if (strcmp(command, "ls") == 0) return CMD_LS;
    if (strcmp(command, "echoln") == 0) return CMD_ECHOLN;
    if (strcmp(command, "history") == 0) return CMD_HISTORY;
    if (strcmp(command, "server") == 0) return CMD_SERVER;
    return CMD_UNKNOWN;
}

int shell(void) {
    char input[MAX_INPUT_SIZE];
    char *tokens[MAX_NUM_TOKENS];
    bool running = true;
    pid_t server_pid = -1;

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
        CommandType commandType = getCommandType(tokens[0]);
        switch (commandType) {
            case CMD_EXIT:
                addToHistory(input);
                running = false;
                break;

            case CMD_RELOAD:
                addToHistory(input);
                system("clear");
                ascii();
                break;

            case CMD_HELP:
                addToHistory(input);
                Help();
                break;

            case CMD_IPLOOKUP:
                addToHistory(input);
                system("./dist/ip");
                break;

            case CMD_PING:
                addToHistory(input);
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
                break;
            
            case CMD_SCANNER:
                addToHistory(input);
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
                break;

            case CMD_DNSLOOKUP:
                addToHistory(input);
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
                break;

            case CMD_SUBNET:
                addToHistory(input);
                system("./dist/subnet");
                break;
            
            case CMD_WHOIS:
                addToHistory(input);
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
                break;

            case CMD_DIRB:
                addToHistory(input);
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
                break;

            case CMD_SNIFFER:
                addToHistory(input);
                system("./dist/sniffer");
                break;

            case CMD_HASHIDENT:
                addToHistory(input);
                system("./dist/hash");
                break;
            
            case CMD_HASH:
                addToHistory(input);
                system("./dist/genhash");
                break;

            case CMD_LS:
                addToHistory(input);
                list_command();
                break;

            case CMD_ECHOLN:
                addToHistory(input);
                echolnCommand(tokens);
                break;

            case CMD_HISTORY:
                addToHistory(input);
                printHistory();
                break;

            case CMD_SERVER:
                addToHistory(input);
                if (tokenCount >= 2 && strcmp(tokens[1], "kill") == 0) {
                    if (server_pid > 0) {
                        killServer(server_pid);
                        server_pid = -1;
                        printf("Server killed.\n");
                    } else {
                        printf("Server is not running.\n");
                    }
                } else if (tokenCount >= 2 && strcmp(tokens[1], "status") == 0) {
                    if (server_pid > 0) {
                        printf("Server is active (PID: %d).\n", server_pid);
                    } else {
                        printf("Server is down.\n");
                    }
                } else {
                    if (server_pid > 0) {
                        printf("Server is already running (PID: %d).\n", server_pid);
                    } else {
                        server_pid = startServer();
                        printf("Server started (PID: %d).\n", server_pid);
                    }
                }
                break;

            case CMD_UNKNOWN:
                bool background = false;

                char *last_token = tokens[tokenCount - 1];
                if (strcmp(last_token, "&") == 0) {
                    background = true;
                    *last_token = '\0';
                }

                int fd_in = 0, fd_out = 1;
                for (int i = 0; tokens[i] != NULL; i++) {
                    if (strcmp(tokens[i], "<") == 0) {
                        tokens[i] = NULL;
                        fd_in = open(tokens[i + 3], O_RDONLY);
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
                addToHistory(input);
                executeCommand(tokens, background);
                break;
        }
    }

    return 0;
}