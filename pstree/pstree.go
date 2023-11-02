package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"syscall"
)

type Process struct {
	pid  uint64
	ppid uint64
}

const PROC_BASE = "/proc"
const MAXFILE = 4096

// readProc is modeled on the original read_proc from pstree
// written by Werner Almesberger and Craig Small
func readProc() {
	pfiles, err := os.ReadDir(PROC_BASE)
	if err != nil {
		log.Fatal(err)
	}
	exp, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err) // this should never happen!
	}

	for _, file := range pfiles {
		if file.IsDir() {
			name := file.Name()

			res := exp.Match([]byte(name))
			if res {
				path := PROC_BASE + "/" + name + "/stat"
				file, fopen_err := os.Open(path)
				if fopen_err != nil {
					log.Print(fopen_err, "process disappeared")
					continue
				}
				path = PROC_BASE + "/" + name
				var st syscall.Stat_t
				stat_err := syscall.Stat(path, &st)
				if stat_err != nil {
					log.Print(stat_err, "process disappeared")
					continue
				}

				// then read the stat file
				buffer := make([]byte, MAXFILE)
				// take care with this read!
				// TODO: cleanup
//				rread, read_err := file.Read(buffer)
				scanner := bufio.NewScanner(file)
				if read_err != nil {
					log.Print(read_err)
				}
				fmt.Print(string(buffer[:rread]))
			}
		}
	}
}

// syscall: open "/proc" O_RDONLY|O_NONBLOCK|O_CLOEXEC|
// openat(AT_FDCWD, "/proc", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 5
// newfstatat(5, "", {st_mode=S_IFDIR|0555, st_size=0, ...}, AT_EMPTY_PATH) = 0
// mmap(NULL, 135168, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f4f296fe000
// mmap(NULL, 135168, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f4f296dd000
// getdents64(5, 0x556aa26441a0 /* 441 entries */, 32768) = 12648
// newfstatat(AT_FDCWD, "/proc/1", {st_mode=S_IFDIR|0555, st_size=0, ...}, 0) = 0

// read processes

// get ppid (column 4 of stat)

// add proc (command, pid, ppid)

func main() {
	readProc()
}
