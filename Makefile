SRCBUILD := build/cgo
SRCBIN := bin
CC := gcc
CFLAGS := -Wall -Wextra -std=c99 -pthread
CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := $(SRCBUILD)/path/path.a $(SRCBUILD)/path/path.so \
		$(SRCBUILD)/os/user.a $(SRCBUILD)/os/user.so \
        $(SRCBUILD)/math/bits.a $(SRCBUILD)/math/bits.so


.PHONY: all

all: $(LIBS)

$(SRCBUILD)/path/path.a: path/path.go
	$(CGO) $(STATIC) -o $(SRCBUILD)/path/path.a $<

$(SRCBUILD)/path/path.so: path/path.go
	$(CGO) $(SHARED) -o $(SRCBUILD)/path/path.so $<

$(SRCBUILD)/os/user.a: os/user.go
	$(CGO) $(STATIC) -o $(SRCBUILD)/os/user.a $<

$(SRCBUILD)/os/user.so: os/user.go
	$(CGO) $(SHARED) -o $(SRCBUILD)/os/user.so $<

$(SRCBUILD)/math/bits.a: math/bits.go
	$(CGO) $(STATIC) -o $(SRCBUILD)/math/bits.a $<

$(SRCBUILD)/math/bits.so: math/bits.go
	$(CGO) $(SHARED) -o $(SRCBUILD)/math/bits.so $<

clean:
	find $(SRCBUILD) -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
