all:
	go build -o _ch/ch _ch/ch.go

install:
	mkdir -p ~/.local/share/cheatsheets
	cp -a [a-z]* ~/.local/share/cheatsheets/
