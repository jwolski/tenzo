all:
	go build -o bin/collabedit-client collabedit/client/*.go
	go build -o bin/collabedit-server collabedit/server/*.go

clean:
	rm bin/*

collabedit:
	@[ -f collabedit-server.pid ] && \
		kill `cat collabedit-server.pid` && \
		rm collabedit-server.pid
	@(./bin/collabedit-server & echo $$! > collabedit-server.pid)

.PHONY: collabedit
