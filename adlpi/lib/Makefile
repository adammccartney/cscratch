# Makefile to bs
#
# This Makefile relies on the assumption that each C file in this
# directory belongs in the library. It takes advantage of some features
# specific to GNU make. If you don't have GNU make, remove this file
# and rename "Makefile.std" in this directory to "Makefile".
#

CC := gcc
CFLAGS = -fPIC -shared -Wall -Wextra -Werror -pedantic -std=c99 -g -O2 -I ../
PREFIX := /usr/local

ADLPI_LIB = ../libadlpi.so

SOURCES=$(wildcard *.c)
HEADERS=$(wildcard *.h)
OBJECTS=$(patsubst %.c,%.o,$(wildcard *.c))

default: ${ADLPI_LIB}

${ADLPI_LIB} : ${OBJECTS}
	${RM} ${ADLPI_LIB}
	$(CC) -fPIC -shared $(CFLAGS) -o ${ADLPI_LIB} ${OBJECTS}

get_num.o: get_num.c get_num.h
	$(CC) -c $(CFLAGS) $<

error_functions.o : error_functions.c error_functions.h ename.c.inc
	$(CC) -c $(CFLAGS) $<

ename.c.inc :
	sh Build_ename.sh > ename.c.inc
	echo 1>&2 "ename.c.inc built"

install:
	mkdir -p $(DESTDIR)$(PREFIX)/lib
	cp -f $(ADLPI_LIB) $(DESTDIR)$(PREFIX)/lib
	cp -f adlpi_hdr.h $(DESTDIR)$(PREFIX)/include
	# note: if umask is set to something other than 0022, you will need
	# something like this:
	chmod 0755 $(DESTDIR)$(PREFIX)/lib/libadlpi.so

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/lib/libadlpi.so

clean :
	${RM} *.o ename.c.inc ${ADLPI_LIB}
