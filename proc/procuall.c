#include <dirent.h>
#include <stdarg.h>
#include <sys/stat.h>

#include "adio.h"
#include "cscratch_common.h"
#include "ugid_info.h"


/* procuall.c: all process being run by a user
 *
 * usage: procuall <username>
 * */

#define MAXLINE 512
#define MAXFNAME 128
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

char*
make_filename(const char* pid) {
    char* fname = (char*)malloc(MAXFNAME);
    if (s_isinteger(pid)) {
        sprintf(fname, "/proc/%s/status", pid);
        return fname;
    }
    return NULL;
}

int
main (int argc, char* argv[]) {
    if (argc != 2) {
        printf("usage: procuall <username>\n");
        exit(EPERM);
    }
    char* uname = argv[1];
    uid_t uid = uidFromName(uname);
    printf("user: %s\tuid: %d\n", uname, uid);

    DIR* dirp;
    struct dirent* dp;
    char* fname;
    FILE* fp;
    int fd;
    char* lone;
    struct stat sb;

    dirp = opendir(PROC);
    if (dirp) {
        errno = 0;
        while ((dp = readdir(dirp)) != NULL) {
            fname = make_filename(dp->d_name);
            if (access(fname, F_OK) == 0) {
                fp = fopen(fname, "r");
                if (fp == NULL) { /* Entered a bad state */
                    fprintf(stderr, "Error: fopen attempted read on %s, returned %d", fname, errno);
                } else {
                    fd = fileno(fp);
                    if (fstat(fd, &sb) == -1) {
                        return -1; /* just cheese it! */
                    }
                    if (uid == sb.st_uid) {
                        lone = fgetLine(MAXLINE, fp);
                        printf("%-24s pid:%-30.30s\n", lone, dp->d_name);
                    }
                }
            } else {
                fprintf(stderr, "Error: %s failed before fopen\n", dp->d_name);
            }
            free(fname);
        }
        closedir(dirp);
    }
    return 0;
}
