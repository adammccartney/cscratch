SHELL := /usr/bin/bash

CC := gcc
CFLAGS := -O2 -g

SRC := ./src
BUILD := ./build

all : $(BUILD)/echo $(BUILD)/client

$(BUILD)/echo: $(SRC)/echo.c builddir
	$(CC) $(CFLAGS) -o $@ $<

$(BUILD)/client: $(SRC)/client.c builddir
	$(CC) $(CFLAGS) -o $@ $<

builddir:
	mkdir -p $(BUILD)

.PHONY: all builddir

clean:
	rm -rf $(BUILD)
