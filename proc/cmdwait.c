#include "cscratch_common.h"

/* Simple test program for examining the contents of some files in the proc
 * filesystem. */

#define MAXLINE 256

int main(int argc, char* argv[]) {
    char* arg1 = argv[1];
    int pid = getpid();
    printf("%d\n", pid);
    sleep(90);
    printf("%s\n", arg1);
    return 0;
}
