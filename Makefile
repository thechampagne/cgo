SRCBUILD := build/cgo
SRCBIN := bin
CC := gcc
CFLAGS := -Wall -Wextra -std=c99 -pthread
CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := $(SRCBUILD)/path/libpath.a $(SRCBUILD)/path/libpath.so


all: $(LIBS)

$(SRCBUILD)/path/libpath.a: path/path.go
	$(CGO) $(STATIC) -o $(SRCBUILD)/path/libpath.a $<

$(SRCBUILD)/path/libpath.so: path/path.go
	$(CGO) $(SHARED) -o $(SRCBUILD)/path/libpath.so $<