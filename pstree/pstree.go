package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "github.com/adammccartney/algorill/pkg/datastruct/chqueue"
)

type Process struct {
	pid  uint64
	ppid uint64
}

const PROC_BASE = "/proc"

func readProc() {
	pfiles, err := os.ReadDir(PROC_BASE)
	if err != nil {
		log.Fatal(err)
	}

	processes := make([]string, len(pfiles))
	for i, file := range pfiles {
		fmt.Println(file.Name())
		name := file.Name()
		res, err := filepath.Match(name, "[0-9]+")
		if err != nil {
			log.Print(err)
		} else {
			processes[i] = file.Name()
			fmt.Println(res)
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
