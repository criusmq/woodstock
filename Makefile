ifeq ($(origin GOPATH), undefined)
	GOPATH:=$(CURDIR)
else
	GOPATH:=$(CURDIR):$(GOPATH)
endif


# Colors
RESETCOLOR="\033[0m"
RCOLOR="\033[31m"
GCOLOR="\033[32m"
BCOLOR="\033[34m"
CCOLOR="\033[36m"
YCOLOR="\033[33m"
MCOLOR="\033[35m"
KCOLOR="\033[30m"
WCOLOR="\033[37m"

 # PHONY tasks are tasks not tied to files
.PHONY: all build test clean env coverage doc no_targets__ list

all: build test
	
no_targets__:
list:
	@sh -c "$(MAKE) -p no_targets__ | awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | grep -v '__\$$' | sort"



doc:
	@printf '%bBuilding documentation%b\n' $(BCOLOR) $(RESETCOLOR)
	@printf '%b ... ... (NOT)%b\n' $(RCOLOR) $(RESETCOLOR)

build:
	@printf '%bBuilding software%b\n' $(BCOLOR) $(RESETCOLOR)
	@printf '%b ... ... (NOT)%b\n' $(RCOLOR) $(RESETCOLOR)
# go build

test: 
	@printf '%bTesting: %b' $(BCOLOR) $(RESETCOLOR)
	GOPATH=$(GOPATH) go test -v ./...

coverage:
	@printf '%bTest Coverage%b\n' $(BCOLOR) $(RESETCOLOR)
	@printf '%b ... ... (NOT)%b\n' $(RCOLOR) $(RESETCOLOR)
# go test -cover

clean:
	@printf '%bCleaning%b\n' $(BCOLOR) $(RESETCOLOR)
	@printf '%b ... ... (NOT)%b\n' $(RCOLOR) $(RESETCOLOR)

env:
	@printf '%bGo Environment:%b\n' $(BCOLOR) $(RESETCOLOR) 
	@go env

fmt:
	@printf '%bfmt:%b ' $(BCOLOR) $(RESETCOLOR)
	gofmt -w=true src

vet:
	@printf '%bvet: %b' $(BCOLOR) $(RESETCOLOR)
	GOPATH=$(GOPATH) go tool vet -v src/woodstock/

# go clean
