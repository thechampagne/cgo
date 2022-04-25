/*
typedef struct {
  int is_match;
  char* error;
} path_match_t;

typedef struct {
  char* dir;
  char* file;
} path_split_t;

path_match_t* path_match(char* pattern, char* name)
{
  path_match_t* self = (path_match_t*) malloc(sizeof(path_match_t));
  struct _path_match_return match = _path_match(pattern, name);
  if (self == NULL)
  {
	free(match.r1);
	return NULL;
  }
  if (match.r1 != 0)
  {
    self->error = (char*) malloc((strlen(match.r1) + 1) * sizeof(char));
  	if (self->error == NULL)
  	{
		free(match.r1);
		return NULL;
  	}
    strncpy(self->error, match.r1, strlen(match.r1));
    free(match.r1);
    return self;
  }

  if (match.r0) {
    self->is_match = 1;
  } else {
    self->is_match = 0;
  }
  self->error = NULL;

  return self;
}

path_split_t* path_split(char* path)
{
  path_split_t* self = (path_split_t*) malloc(sizeof(path_split_t));
  struct _path_split_return s = _path_split(path);
  if (self == NULL)
  {
	free(s.r0);
	free(s.r1);
	return NULL;
  }
  self->dir = s.r0;
  self->file = s.r1;
  return self;
}

void path_split_clean(path_split_t* self)
{
	if (self != NULL)
	{
		if (self->dir != NULL)
		{
			free(self->dir);
		}
		if (self->file != NULL)
		{
			free(self->file);
		}
		free(self);
	}
}

void path_match_clean(path_match_t* self)
{
	if (self != NULL)
	{
		if (self->error != NULL)
		{
			free(self->error);
		}
		free(self);
	}
}
 */
package main

/*
#include <stdlib.h>
#include <string.h>
 */
import "C"
import (
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