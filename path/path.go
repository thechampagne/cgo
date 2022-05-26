package main

/*
#include <stdlib.h>
#include <string.h>
 */
import "C"
import (
	"unsafe"
	p "path"
)

//export path_base
func path_base(path *C.char) *C.char {
	return C.CString(p.Base(C.GoString(path)))
}

//export path_clean
func path_clean(path *C.char) *C.char {
	return C.CString(p.Clean(C.GoString(path)))
}

//export path_dir
func path_dir(path *C.char) *C.char {
	return C.CString(p.Dir(C.GoString(path)))
}

//export path_ext
func path_ext(path *C.char) *C.char {
	return C.CString(p.Ext(C.GoString(path)))
}

//export path_is_abs
func path_is_abs(path *C.char) C.int {
	if p.IsAbs(C.GoString(path)) {
		return C.int(1)
	} else {
		return C.int(0)
	}
}

//export path_join
func path_join(elem **C.char, length C.int) *C.char {
  slice := (*[1 << 30]*C.char)(unsafe.Pointer(elem))[:int(length):int(length)]
  array := []string{}
  for _ ,v := range slice {
    array = append(array, C.GoString(v))
  }
  return C.CString(p.Join(array...))
}

//export _path_match
func _path_match(pattern, name *C.char) (C.int, *C.char) {
	matched, err := p.Match(C.GoString(pattern), C.GoString(name))
	if err != nil {
		return C.int(-1), C.CString(err.Error())
	}

	if matched {
		return C.int(1), nil
	} else {
		return C.int(0), nil
	}
}

//export _path_split
func _path_split(path *C.char) (*C.char, *C.char) {
	dir, file := p.Split(C.GoString(path))
	return C.CString(dir), C.CString(file)
}

func main() {}