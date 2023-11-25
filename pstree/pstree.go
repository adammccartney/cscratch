package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

const PROC_BASE = "/proc"
const MAXFILE = 4096
const (
	sym_empty_2  = "  "
	sym_branch_2 = "|-"
	sym_vert_2   = "| "
	sym_last_2   = "`-"
	sym_single_3 = "---"
	sym_first_3  = "-+-"
)

type Proc struct {
	pid      int
	comm     string
	uid      int
	ppid     int
	children *Child
	parent   *Proc
	next     *Proc
}

type Child struct {
	child *Proc
	next  *Child
}

// Walk the tree and find the process with pid
func (list *Proc) findProc(pid int) *Proc {
	for list != nil {
		if list.pid == pid {
			return list
		}
		list = list.next
	}
	return nil
}

// Create a new proc by inserting it into the list
func (list *Proc) newProc(comm string, pid int, uid int) *Proc {
	var pnew *Proc
	pnew.comm = comm
	pnew.pid = pid
	pnew.uid = uid
	pnew.children = nil
	pnew.parent = nil
	pnew.next = list
	list = pnew
	return list
}

// walk the nodes in the parent graph.
// find the point at which to insert the child process
// insert child into graph
func (parent *Proc) addChild(child *Proc) {
	var cnew *Child
	var walk **Child
	cnew.child = child
	for walk = &parent.children; *walk != nil; walk = &(*walk).next {
		// insert "by pid"
		// break as soon as the pid of the child in the currently focused node
		// of the walk is greater than the pid of the child being inserted
		if (*walk).child.pid > child.pid {
			break
		}
	}
	cnew.next = *walk
	*walk = cnew
}

func (list *Proc) addProc(comm string, pid int, ppid int, uid int) *Proc {
	var this *Proc
	var parent *Proc

	if list == nil {
		return &Proc{pid, comm, uid, ppid, nil, nil, nil}
	}

	this = list.findProc(pid)
	if this == nil {
		this = list.newProc(comm, pid, uid)
	} else {
		this.comm = comm
		this.uid = uid
	}
	if pid == ppid {
		ppid = 0
	}
	parent = list.findProc(ppid)
	if parent == nil {
		parent = list.newProc("?", ppid, 0)
	}
	parent.addChild(this)
	this.parent = parent
	return this
}

// readProc is modeled on the original read_proc from pstree
// written by Werner Almesberger and Craig Small
// this function reads proc and gathers all useful information
func (list *Proc) readProc() {

	pfiles, err := os.ReadDir(PROC_BASE)
	if err != nil {
		log.Fatal(err)
	}

	exp, err := regexp.Compile("[0-9]+") // for pid
	if err != nil {
		log.Fatal(err) // this should never happen!
	}

	for _, file := range pfiles {
		if file.IsDir() {
			name := file.Name()
			res := exp.Match([]byte(name))
			if res {
				pid, err_pid := strconv.Atoi(name)
				if err_pid != nil {
					log.Print(err_pid)
					continue
				}
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
				rread, read_err := file.Read(buffer)
				if read_err != nil {
					log.Print(read_err)
				}
				text := string(buffer[:rread])
				// parse command -- may include spaces, "(" or ")"
				c_start := strings.Index(text, "(")
				c_end := strings.LastIndex(text, ")")
				comm := text[c_start : c_end+1]
				// find the ppid
				tmp := text[c_end+1:]
				fields := strings.Split(tmp, " ")
				ppid, err_ppid := strconv.Atoi(fields[2])
				if err_ppid != nil {
					log.Print(err_ppid)
					continue
				}
				//fmt.Println("pid: ", pid)
				//fmt.Println("command: ", comm)
				//fmt.Println("ppid: ", ppid)
				//fmt.Print(string(buffer[:rread]))
				list.addProc(comm, pid, ppid, int(st.Uid))
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

func (list *Proc) dumpTree(pid int) {
	fmt.Println("will print tree starting at pid: ", pid)
}

func (list *Proc) dumpByUser(pid int, uid int) {
	fmt.Printf("tree start at pid %s matching uid %s\n", pid, uid)
}

func main() {
	args := os.Args

	if len(args) > 2 {
		log.Fatal("Too many arguments. Usage: pstree [username]")
	}

	var list *Proc
	list.readProc()

	// take an optional username, then restrict for processes
	// belonging to that user
	var uid int
	pw := false
	if len(args) == 2 {
		user, user_err := user.Lookup(args[1])
		if user_err != nil {
			log.Print(user_err)
		}
		pw = true
		uid, _ = strconv.Atoi(user.Uid)
	}
	pid := 1 // root of tree to display
	// note, could optionally find subtrees by taking pid as cmdline arg
	init := list.findProc(pid)
	if !pw {
		init.dumpTree(pid)
	} else {
		init.dumpByUser(pid, uid)
	}
}
