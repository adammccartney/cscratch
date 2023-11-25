CC := clang
CFLAGS = -Wall -Wextra -Werror -pedantic -std=c99 -g -O2 -I../lib/
PREFIX := /usr/local
LDFLAGS := -L$(DESTDIR)$(PREFIX)/lib

#GEN_EXE = mq_notify_sig mq_notify_sigwaitinfo mq_notify_thread \
#	  mq_notify_via_signal mq_notify_via_thread \
#	  pmsg_create pmsg_getattr pmsg_receive pmsg_send pmsg_unlink


SOURCES = pmsg_create.c pmsg_getattr.c pmsg_receive.c pmsg_send.c pmsg_unlink.c
OBJECTS = $(SOURCES:.c=.o)

EXE = pmsg_create pmsg_getattr pmsg_receive pmsg_send pmsg_unlink

all : ${EXE}

install: ${EXE}
	cp -f pmsg_create $(DESTDIR)$(PREFIX)/bin
	cp -f pmsg_getattr $(DESTDIR)$(PREFIX)/bin
	cp -f pmsg_receive $(DESTDIR)$(PREFIX)/bin
	cp -f pmsg_send $(DESTDIR)$(PREFIX)/bin
	cp -f pmsg_unlink $(DESTDIR)$(PREFIX)/bin

LDLIBS = -ladlpi
	# All of the programs in this directory need the
	# realtime library, librt.

clean :
	${RM} ${EXE} *.o

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/pmsg_create
	rm -f $(DESTDIR)$(PREFIX)/bin/pmsg_getattr
	rm -f $(DESTDIR)$(PREFIX)/bin/pmsg_receive
	rm -f $(DESTDIR)$(PREFIX)/bin/pmsg_send
	rm -f $(DESTDIR)$(PREFIX)/bin/pmsg_unlink

showall :
	@ echo ${EXE}


pmsg_create: pmsg_create.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
pmsg_create.o: pmsg_create.c
	$(CC) -c $(CFLAGS) $<

pmsg_getattr: pmsg_getattr.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
pmsg_getattr.o: pmsg_getattr.c
	$(CC) -c $(CFLAGS) $<

pmsg_receive: pmsg_receive.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
pmsg_receive.o: pmsg_receive.c
	$(CC) -c $(CFLAGS) $<

pmsg_send: pmsg_send.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
pmsg_send.o: pmsg_send.c
	$(CC) -c $(CFLAGS) $<

pmsg_unlink: pmsg_unlink.o
	$(CC) $(LDFLAGS) -o $@ $^ $(LDLIBS)
pmsg_unlink.o: pmsg_unlink.c
	$(CC) -c $(CFLAGS) $<
