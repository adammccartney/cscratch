.POSIX:
.SUFFIXES:
CC := gcc
CFLAGS := -std=c99 -g3 -Wall -Werror -Wextra -O2
PREFIX := /usr/local
LDFLAGS := -L $(PREFIX)/lib
LDLIBS := -labinarytree

iotraverse: iotraverse.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
iotraverse.o: iotraverse.c
	$(CC) -c $(CFLAGS) $<

.PHONY: clean
clean:
	rm -rf iotraverse *.o
