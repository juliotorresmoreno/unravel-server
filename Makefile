#/bin/sh

CC = go
CFLAGS = build
FROM = .
TO = ./bin
BIN = unravel-server
INSTALL = /home/ubuntu/unravel
SERVICE_FROM = ./etc/unravel.service
SERVICE_TO = /etc/systemd/system

all: $(OBJ)
	$(CC) $(CFLAGS) $(FROM)
	mv $(BIN) $(TO)

install: $(OBJ)
	cp $(TO)/$(BIN) $(INSTALL) -f
	
#cp $(SERVICE_FROM) $(SERVICE_TO) -f