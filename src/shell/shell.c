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
    CMD_HISTORY,
    CMD_SERVER,
    CMD_PEGASUSEDIT,
    CMD_TRACEROUTE,
    CMD_WEBSERVER,
    CMD_REVSHELL,
    CMD_GETIP,
    CMD_MACSPOOF,
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
    if (strcmp(command, "history") == 0) return CMD_HISTORY;
    if (strcmp(command, "server") == 0) return CMD_SERVER;
    if (strcmp(command, "edit") == 0) return CMD_PEGASUSEDIT;
    if (strcmp(command, "traceroute") == 0) return CMD_TRACEROUTE;
    if (strcmp(command, "webserver") == 0) return CMD_WEBSERVER;
    if (strcmp(command, "revshell") == 0) return CMD_REVSHELL;
    if (strcmp(command, "getip") == 0) return CMD_GETIP;
    if (strcmp(command, "macspoof") == 0) return CMD_MACSPOOF;
    return CMD_UNKNOWN;
}

int shell(void) {
    char input[MAX_INPUT_SIZE];
    char *tokens[MAX_NUM_TOKENS];
    bool background = false;
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
                if (tokenCount > 2) {
                    return 1;
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/ip %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;

            case CMD_PING:
                addToHistory(input);
                if (tokenCount < 2) {
                    pingHelp();
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/ping %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;
            
            case CMD_SCANNER:
                addToHistory(input);
                if (tokenCount < 1) {
                    return 1;
                } else {
                   char command[MAX_INPUT_SIZE];
                   snprintf(command, sizeof(command), "dotnet run --project src/tools/PScan/PScan.csproj %s", tokens[1]);
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
                    snprintf(command, sizeof(command), "./dist/dns %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;

            case CMD_SUBNET:
                addToHistory(input);
                if (tokenCount > 3) {
                    return 1;
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/subnet %s %s", tokens[1], tokens[2]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;
            
            case CMD_WHOIS:
                addToHistory(input);
                if (tokenCount > 2) {
                    whoisHelp();
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/whois %s", tokens[1]);
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
                    snprintf(command, sizeof(command), "./dist/dirb %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;

            case CMD_SNIFFER:
                addToHistory(input);
                system("./sniffer");
                break;

            case CMD_HASHIDENT:
                addToHistory(input);
                if (tokenCount > 2) {
                    hashIdentHelp();
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/id %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;
            
            case CMD_HASH:
                addToHistory(input);
                if (tokenCount > 3) {
                    hashHelp();
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/hash %s %s", tokens[1], tokens[2]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;

            case CMD_LS:
                addToHistory(input);
                list_command();
                break;

            case CMD_HISTORY:
                addToHistory(input);
                printHistory();
                break;

            case CMD_REVSHELL:
                addToHistory(input);
                if (tokenCount >= 2 && strcmp(tokens[1], "kill") == 0) {
                    if (server_pid != -1) {
                        kill(server_pid, SIGTERM);
                        printf("Server killed");
                        server_pid = -1;
                    } else {
                        printf("Server not running\n");
                    }
                } else if (tokenCount >= 2 && strcmp(tokens[1], "status") == 0) {
                    if (server_pid != -1) {
                        printf("Server is running (PID: %d).\n", server_pid);
                    } else {
                        printf("Server not running.\n");
                    }
                } else {
                    if (server_pid != -1) {
                        printf("Server is already running (PID: %d).\n");
                    } else {
                        server_pid = fork();
                        if (server_pid == 0) {
                            system("./dist/revshell &");
                        } else if (server_pid > 0) {
                            printf("Server started (PID: %d).\n", server_pid);
                        } else {
                            perror("fork");
                        }
                    }
                }
                break;

            case CMD_WEBSERVER:
                addToHistory(input);
                if (tokenCount >= 2 && strcmp(tokens[1], "kill") == 0) {
                    if (server_pid != -1) {
                        kill(server_pid, SIGTERM);
                        printf("Server killed");
                        server_pid = -1;
                    } else {
                        printf("Server not running\n");
                    }
                } else if (tokenCount >= 2 && strcmp(tokens[1], "status") == 0) {
                    if (server_pid != -1) {
                        printf("Server is running (PID: %d).\n", server_pid);
                    } else {
                        printf("Server not running.\n");
                    }
                } else {
                    if (server_pid != -1) {
                        printf("Server is already running (PID: %d).\n");
                    } else {
                        server_pid = fork();
                        if (server_pid == 0) {
                            system("./dist/webserver &");
                        } else if (server_pid > 0) {
                            printf("Server started (PID: %d).\n", server_pid);
                        } else {
                            perror("fork");
                        }
                    }
                }
                break;

            case CMD_TRACEROUTE:
                addToHistory(input);
                if (tokenCount > 2) {
                    return 1;
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/traceroute %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;

            case CMD_PEGASUSEDIT:
                addToHistory(input);
                if (tokenCount > 2) {
                    return 1;
                } else {
                    char command[MAX_INPUT_SIZE];
                    snprintf(command, sizeof(command), "./dist/pegasusedit %s", tokens[1]);
                    int result = system(command);
                    if (result == -1) {
                        perror("system");
                    }
                }
                break;

            case CMD_SERVER:
                addToHistory(input);
                if (tokenCount >= 2 && strcmp(tokens[1], "kill") == 0) {
                    if (server_pid != -1) {
                        kill(server_pid, SIGTERM);
                        printf("Server killed");
                        server_pid = -1;
                    } else {
                        printf("Server not running\n");
                    }
                } else if (tokenCount >= 2 && strcmp(tokens[1], "status") == 0) {
                    if (server_pid != -1) {
                        printf("Server is running (PID: %d).\n", server_pid);
                    } else {
                        printf("Server not running.\n");
                    }
                } else {
                    if (server_pid != -1) {
                        printf("Server is already running (PID: %d).\n");
                    } else {
                        server_pid = fork();
                        if (server_pid == 0) {
                            system("./dist/server &");
                        } else if (server_pid > 0) {
                            printf("Server started (PID: %d).\n", server_pid);
                        } else {
                            perror("fork");
                        }
                    }
                }
                break;

            case CMD_GETIP:
                system("./dist/getIP");
                break;

            case CMD_MACSPOOF:
                system("./dist/mac");
                break;

            case CMD_UNKNOWN:
                executeCommand(tokens, background);
                break;
        }
    }

    return 0;
}
