BINDIR=$(HOME)/.local/bin
BINARY=cm2secret

.PHONY: default
default: build

.PHONY: build
build:
	go build -o $(BINARY)  main.go

.PHONY: clean
clean:
	rm -f $(BINARY)

.PHONY: install
install: build
	test -d  $(BINDIR) ||  mkdir -p $(BINDIR)
	cp $(BINARY) $(BINDIR)/.

.PHONY: uninstall
uninstall:
	rm -f $(BINDIR)/$(BINARY)

