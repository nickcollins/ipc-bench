#ifndef IPC_BENCH_PROCESS_H
#define IPC_BENCH_PROCESS_H

#include <sys/types.h>

char *find_build_path();

pid_t start_process(char *argv[]);

void copy_arguments(char *arguments[], int argc, char *argv[]);

pid_t start_child(char *name, int argc, char *argv[]);

void start_children(char *prefix, int argc, char *argv[]);

#endif /* IPC_BENCH_PROCESS_H */
