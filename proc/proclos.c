/* proclos: test what happens if a process disappears between the call to
 *
 * program will attempt to read /proc/$PID/status file, but the process with
 * $PID will be killed between the call to readdir and the call to fopen on the
 * /proc/$PID/status file.
 */
#include <dirent.h>
#include <stdarg.h>
#include <sys/stat.h>

#include "cscratch_common.h"

#define MAXLINE 512
#define LPID 5
#define PROC "/proc"

/* check that string s is a contiguous array of integer characters */
bool s_isinteger(const char* s) {
    bool result = (*s != '\0');
    while (*s != '\0') {
        if ((*s < '0') || (*s > '9')) {
            return false;
        }
        s++;
    }
    return result;
}

int main(int argc, char* argv[]) {
    if (argc != 3) {
        printf("usage: proclos <pid> <t>\n");
        exit(EPERM);
    }
    char* pid;
    pid = argv[1];
    if (!s_isinteger(pid)) {
        fprintf(stderr, "Error: invalid pid %s\n", pid);
        exit(EPERM);
    } 
    /* Otherwise set up the path we want to read */
    char fname[FILENAME_MAX];
    sprintf(fname, "/proc/%s/status", pid);

    int time;
    time = atoi(argv[2]);
    if (!time) {
        fprintf(stderr, "Error: %s not positive integer\n", argv[2]);
        exit(EPERM);
    }

    DIR* dirp;
    struct dirent* dp;
    FILE* fp;
    int fd;
    char* line = NULL;
    size_t len = 0;
    ssize_t rread = 0;
    char* out = NULL;
    struct stat sb;

    dirp = opendir(PROC);
    if (dirp) {
        errno = 0;
        if ((dp = readdir(dirp)) != NULL) {
            sleep(time); /* Sleep while the process gets killed */
            fp = fopen(fname, "r");
            if (fp == NULL) {
                fprintf(stderr, "Error: fopen failed to complete with code %d\n", errno);
                exit(ESRCH); /* Process not found */
            }
            fd = fileno(fp);
            if (fstat(fd, &sb) == -1) {
                return -1; /* just cheese it! */
            }
            out = (char*)malloc(MAXLINE);
            if (out == NULL) {
                fprintf(stderr, "Error: malloc\n");
                goto err_malloc;
            }
            if ((rread = getline(&line, &len, fp)) != -1) {
                sprintf(out, "%-24s pid:%-30.30s\n", line, pid);
                int llen = strlen(out);
                fwrite(out, llen, 1, stdout);
            }
        }
err_malloc:
        closedir(dirp);
    }
    return 0;
}
