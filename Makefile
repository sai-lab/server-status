GO_BUILDOPT := -ldflags '-s -w'

gom:
	@if [`which go` = "go not found"]; then \
		echo "Please install Golang"; \
	else \
		go get github.com/mattn/gom; \
		gom install; \
	fi

link:
	mkdir -p $(GOPATH)/src/github.com/sai-lab
	ln -sf $(CURDIR) $(GOPATH)/src/github.com/sai-lab/server-status
	ln -sf $(CURDIR)/vendor $(CURDIR)/vendor/src

fmt:
	gom exec goimports -w *.go lib/*/*.go

build: fmt
	gom build $(GO_BUILDOPT) -o bin/server-status main.go
