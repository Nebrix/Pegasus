#include <stdio.h>
#include <stdio.h>
#include <stdlib.h>
#include <dirent.h>
#include <unistd.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <time.h>
#include <ctype.h>
#include <stdbool.h>
#include "core.h"

const char* humanReadableSize(off_t size) {
    static char result[100];
    if (size >= 1024 * 1024) {
        snprintf(result, sizeof(result), "%.1f MiB", (float)size / (1024 * 1024));
    } else if (size >= 1024) {
        snprintf(result, sizeof(result), "%.1f KiB", (float)size / 1024);
    } else {
        snprintf(result, sizeof(result), "%ld B", (long)size);
    }
    return result;
}

const char* humanReadableTime(time_t mod_time) {
    static char result[100];
    time_t now;
    time(&now);
    double seconds_ago = difftime(now, mod_time);
    
    if (seconds_ago < 60) {
        snprintf(result, sizeof(result), "%.0f seconds ago", seconds_ago);
    } else if (seconds_ago < 3600) {
        snprintf(result, sizeof(result), "%.0f minutes ago", seconds_ago / 60);
    } else if (seconds_ago < 86400) {
        snprintf(result, sizeof(result), "%.0f hours ago", seconds_ago / 3600);
    } else if (seconds_ago < 604800) {
        snprintf(result, sizeof(result), "%.0f days ago", seconds_ago / 86400);
    } else {
        struct tm *tm_info;
        tm_info = localtime(&mod_time);
        strftime(result, sizeof(result), "%b %d, %Y %H:%M", tm_info);
    }
    
    return result;
}

void printTableRow(int number, const char *name, const char *type, off_t size, time_t modified) {
    printf("│ %3d │ %-16s │ %-4s │ %-10s │ %-20s │\n", number, name, type, humanReadableSize(size), humanReadableTime(modified));
}

void printTableHeader() {
    printf("╭─────┬──────────────────┬──────┬────────────┬──────────────────────╮\n");
    printf("│ %-3s │ %-16s │ %-4s │ %-10s │ %-20s │\n", "#", "Name", "Type", "Size", "Modified");
    printf("├─────┼──────────────────┼──────┼────────────┼──────────────────────┤\n");
}

void printTableFooter() {
    printf("╰─────┴──────────────────┴──────┴────────────┴──────────────────────╯\n");
}

void list_command() {
    DIR *dir; 
    struct dirent *entry;
    struct stat filestat;
    int number = 0;

    printTableHeader();

    dir = opendir(".");
    if (dir == NULL) {
        perror("opendir");
        exit(EXIT_FAILURE);
    }

    while ((entry = readdir(dir)) != NULL) {
        if (lstat(entry->d_name, &filestat) == 0) {
            const char *type = (S_ISDIR(filestat.st_mode)) ? "dir" : "file";
            printTableRow(number, entry->d_name, type, filestat.st_size, filestat.st_mtime);
            number++;
        }
    }

    closedir(dir);

    printTableFooter();
}