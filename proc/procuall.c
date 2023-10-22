#include <dirent.h>
#include <stdarg.h>
#include <sys/stat.h>

#include "adio.h"
#include "cscratch_common.h"
#include "ugid_info.h"


/* procuall.c: all process being run by a user */

#define LFNAME 256
#define MAXLINE 512
#define LPID 5

int s_isdigit(const char* s) {
    int result = 0;
    while (*s != '\0') {
        if (('0' <= *s) && ('9' >= *s)) {
            result = 1;
        }
        s++;
    }
    return result;
}

char*
make_filename(const char* pid) {
    char* fname;
    if (s_isdigit(pid)) {
        sprintf(fname, "/proc/%s/status", pid);
        return fname;
    }
    return NULL;
}

int
main (int argc, char* argv[])
{
    char* uname = argv[1];
    uid_t uid = uidFromName(uname);
    printf("user: %s\tuid: %d\n", uname, uid);

    DIR* dirp;
    char* proc = "/proc";
    dirp = opendir(proc);
    struct dirent* dp;
    char* fname;
    int size = 0;
    FILE* fp;
    int fd;
    char* lone;

    struct stat sb;

    if (dirp) {
        errno = 0;
        while ((dp = readdir(dirp)) != NULL) {
            fname = make_filename(dp->d_name);
            if (fname) {
                fp = fopen(fname, "r");
                fd = fileno(fp);
                if (fstat(fd, &sb) == -1) {
                    return -1; /* just cheese it! */
                }
                if (uid == sb.st_uid) {
                    printf("pid :\t%s\n", dp->d_name);
                    lone = fgetLine(MAXLINE, fp);
                    printf("%s\n", lone);
                }
            }
        }
        closedir(dirp);
    }    
    return 0;
}
