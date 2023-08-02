#pragma once

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

void handleSignal(int signal);
void executeCommand(char **tokens, bool background);
int tokenizeInput(char *input, char **tokens);
void addToHistory(const char *input);
void printHistory(void);
int shell(void);
void powerline(void);
void ehcolnCommand(char **tokens);