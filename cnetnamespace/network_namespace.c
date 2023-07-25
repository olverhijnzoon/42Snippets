// network_namespace_demo.c
#define _GNU_SOURCE
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>

#define STACK_SIZE (1024 * 1024)
#define CLONE_NEWNET 0x40000000

static char child_stack[STACK_SIZE];

char* const child_args[] = {
    "/bin/bash",
    NULL
};

int child_main(void* arg) {
    printf("###\n");
    printf("New shell in new network namespace!\n");
    printf("Network interfaces in new network namespace:\n");
    system("ip addr");
    return 0;
}

int main() {
    printf("42 Snippets - Network Namespaces\n\n");
    printf("Network interfaces in initial network namespace:\n");
    system("ip addr");
    printf("Starting a shell in new network namespace...\n");
    int flags = CLONE_NEWNET;
    pid_t pid = clone(child_main, child_stack + STACK_SIZE, flags | SIGCHLD, NULL);
    if (pid < 0) {
        perror("clone failed");
        exit(1);
    }
    waitpid(pid, NULL, 0);
    return 0;
}
