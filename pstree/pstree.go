package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	// "github.com/adammccartney/algorill/pkg/datastruct/chqueue"
)

type Process struct {
	pid  uint64
	ppid uint64
}

// get num active processes

func numActiveProcesses() (int, error) {
	psCmd := exec.Command("sh", "-c", "ps -A | wc -l")

	psOut, err := psCmd.Output()
	if err != nil {
		print(err)
	}
	nproc := strings.TrimSuffix(string(psOut), "\n")
	result, err := strconv.Atoi(nproc)
	if err != nil {
		print(err)
	}
	return result, nil
}


// syscall: open "/proc" O_RDONLY|O_NONBLOCK|O_CLOEXEC|
//openat(AT_FDCWD, "/proc", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 5
//newfstatat(5, "", {st_mode=S_IFDIR|0555, st_size=0, ...}, AT_EMPTY_PATH) = 0
//mmap(NULL, 135168, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f4f296fe000
//mmap(NULL, 135168, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f4f296dd000
//getdents64(5, 0x556aa26441a0 /* 441 entries */, 32768) = 12648
//newfstatat(AT_FDCWD, "/proc/1", {st_mode=S_IFDIR|0555, st_size=0, ...}, 0) = 0

// read processes

func main() {
	procs, err := numActiveProcesses()
	if err != nil {
		print(err)
	}
	fmt.Println(procs)
}
