#include <stdio.h>
#include <string.h>
#include "history.h"

#define MAX_HISTORY_SIZE 1024
#define MAX_INPUT_SIZE 1024
#define MAX_NUM_TOKENS 1024

char history[MAX_HISTORY_SIZE][MAX_INPUT_SIZE];
int historyCount = 0;

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