 # PHONY tasks are tasks not tied to files

.PHONY: all build test clean no_targets__ list

no_targets__:
list:
	sh -c "$(MAKE) -p no_targets__ | awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | grep -v '__\$$' | sort"


all: build test

build:
	echo "Building... (NOT)"
# go build

test:
	echo "Testing... (NOT)"
# go test

coverage:
	echo "Coverage... (NOT)"
# go test -cover

clean:
	echo "Cleaning... (NOT)"
# go clean
