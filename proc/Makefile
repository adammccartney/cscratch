.POSIX:
.SUFFIXES:
CC := gcc
CFLAGS := -g3
LDFLAGS :=
LDLIBS := -ladio

EXECUTABLES = cmdwait procuall proclos

all: $(EXECUTABLES)

cmdwait: cmdwait.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
cmdwait.o: cmdwait.c cscratch_common.h
	$(CC) -c $(CFLAGS) $<

proclos: proclos.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
proclos.o: proclos.c cscratch_common.h
	$(CC) -c $(CFLAGS) $<

procuall: procuall.o ugid_info.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
procuall.o: procuall.c cscratch_common.h ugid_info.h
	$(CC) -c $(CFLAGS) $<
ugid_info.o: ugid_info.c
	$(CC) -c $(CFLAGS) $<

.PHONY: clean

clean:
	rm -rf $(EXECUTABLES) *.o
