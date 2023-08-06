#pragma once

#include <sys/wait.h>
#include <sys/types.h>

int shell(void);
pid_t startServer();
void killServer(pid_t pid);