BIN=eny
ENTRY=main.go

install: cli
	sudo mv $(BIN) /usr/local/bin

cli: main.go
	go build -o $(BIN) $(ENTRY)