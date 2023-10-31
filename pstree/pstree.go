package main

import (
	"fmt"
	"log"
	"os"
	//"path/filepath"
	"regexp"
	// "github.com/adammccartney/algorill/pkg/datastruct/chqueue"
)

type Process struct {
	pid  uint64
	ppid uint64
}

const PROC_BASE = "/proc"
const MAXFILE = 4096

func readProc() {
	pfiles, err := os.ReadDir(PROC_BASE)
	if err != nil {
		log.Fatal(err)
	}
	//processes := make([]string, len(pfiles))
	for _, file := range pfiles {
		if file.IsDir() {
			name := file.Name()
			res, err := regexp.MatchString("[0-9]+", name)
			if err != nil {
				log.Print(err)
			} else if res != false {
				path := PROC_BASE + "/" + name + "/stat"
				file, err := os.Open(path)
				if err != nil {
					// possible disappearing process
					log.Print(err)
				} else {
					buffer := make([]byte, MAXFILE)
					rread, err := file.Read(buffer)
					if err != nil {
						log.Print(err)
					}
					fmt.Print(string(buffer[:rread]))
				}
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
