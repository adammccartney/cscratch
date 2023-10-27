#include <dirent.h>
#include <stdarg.h>
#include <sys/stat.h>

#include "cscratch_common.h"
#include "ugid_info.h"


/* procuall.c: all process being run by a user
 *
 * usage: procuall <username>
 * */

#define MAXLINE 4096
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
int
make_filename(char* fname, const char* pid) {
    if (s_isinteger(pid)) {
        sprintf(fname, "/proc/%s/status", pid);
        return 0;
    }
    return -1;
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
    FILE* fp;
    int fd;

    char* fname = NULL;
    char* line = NULL;
    char* out = NULL;
    ssize_t rread = 0;
    size_t len = 0;
    ssize_t lout = 0;
    struct stat sb;

    dirp = opendir(PROC);
    if (dirp) {
        errno = 0;
        while ((dp = readdir(dirp)) != NULL) {

            fname = (char*)malloc(FILENAME_MAX);
            int rc;
            if ((rc = make_filename(fname, dp->d_name)) == -1) {
                fprintf(stderr, "Error make_filename %d\n", errno);
                goto err_file;
            }

            if (access(fname, F_OK) != 0) {
                fprintf(stderr, "Error: %s failed before fopen\n", dp->d_name);
                goto err_file;
            }

            if ((fp = fopen(fname, "r")) == NULL) { /* Entered a bad state */
                fprintf(stderr, "Error: fopen attempted read on %s, returned %d", fname, errno);
                goto err_file;
            }

            if ((fd = fileno(fp)) == -1) {
                fprintf(stderr, "Error fileno %d", errno);
                fclose(fp);
                goto err_proc;
            }

            if (fstat(fd, &sb) == -1) {
                fprintf(stderr, "Error fstat %d", errno);
                goto err_proc;
            }

            if (uid == sb.st_uid) {
                if ((rread = getline(&line, &len, fp)) != -1) {
                    out = (char*)malloc(MAXLINE);
                    sprintf(out, "%spid:\t%s\n", line, dp->d_name);
                    lout = strlen(out);
                    fwrite(out, lout, 1, stdout);
                    free(out);
                } else {
                    fprintf(stderr, "Error letline %d", errno);
                }
            }
err_proc:
            fclose(fp);
err_file:
            free(fname);
        }
        /* cleanup directory stuff */
        free(dp);
        closedir(dirp);
    }
    return 0;
}
