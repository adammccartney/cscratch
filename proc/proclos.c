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
    char fname[MAXLINE];
    snprintf(fname, MAXLINE, "/proc/%s/status", pid);

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
    struct stat sb;

    dirp = opendir(PROC);
    if (dirp) {
        errno = 0;
        if ((dp = readdir(dirp)) != NULL) {
            sleep(time); /* Sleep while the process gets killed */

            if ((fp = fopen(fname, "r")) == NULL) {
                fprintf(stderr, "Error: fopen failed to complete with code %d\n", errno);
                goto err_file;
            }

            if ((fd = fileno(fp)) == -1) {
                fprintf(stderr, "Error fileno %d\n", errno);
                goto err_proc;
            }

            if (fstat(fd, &sb) == -1) {
                fprintf(stderr, "Error fstat %d\n", errno);
                goto err_proc;
            }

            if ((rread = getline(&line, &len, fp)) != -1) {
                fprintf(stdout, "%spid:\t%s\n", line, pid);
            }
        }
err_proc:
        fclose(fp);
err_file:
        closedir(dirp);
    }
    return 0;
}
