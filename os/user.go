package main

/*
#include <stdlib.h>

typedef struct {
  char* gid;
  char* name;
  char* error;
} user_group_t;

typedef struct {
  char* uid;
  char* gid;
  char* username;
  char* name;
  char* home_dir;
  char* error;
} user_t;

typedef struct {
  size_t length;
  char** buffer;
  char* error;
} user_group_ids_t;
*/
import "C"
import (
  "unsafe"
  u "os/user"
)

//export user_lookup_group
func user_lookup_group(name *C.char) *C.user_group_t {
  self := (*C.user_group_t) (C.malloc(C.size_t(unsafe.Sizeof(C.user_group_t{}))))
  res, err := u.LookupGroup(C.GoString(name))
  if err != nil {
    self.gid = nil
    self.name = nil
    self.error = C.CString(err.Error())
    return self
  }

  self.gid = C.CString(res.Gid)
  self.name = C.CString(res.Name)
  self.error = nil
  return self
}

//export user_lookup_group_id
func user_lookup_group_id(gid *C.char) *C.user_group_t {
  self := (*C.user_group_t) (C.malloc(C.size_t(unsafe.Sizeof(C.user_group_t{}))))
  res, err := u.LookupGroupId(C.GoString(gid))
  if err != nil {
    self.gid = nil
    self.name = nil
    self.error = C.CString(err.Error())
    return self
  }

  self.gid = C.CString(res.Gid)
  self.name = C.CString(res.Name)
  self.error = nil
  return self
}

//export user_current
func user_current() *C.user_t {
  self := (*C.user_t) (C.malloc(C.size_t(unsafe.Sizeof(C.user_t{}))))
  res, err := u.Current()
  if err != nil {
  self.uid = nil
  self.gid = nil
  self.username = nil
  self.name = nil
  self.home_dir = nil
  self.error = C.CString(err.Error())
  return self
  }

  self.uid = C.CString(res.Uid)
  self.gid = C.CString(res.Gid)
  self.username = C.CString(res.Username)
  self.name = C.CString(res.Name)
  self.home_dir = C.CString(res.HomeDir)
  self.error = nil
  return self
}

//export user_lookup
func user_lookup(username *C.char) *C.user_t {
  self := (*C.user_t) (C.malloc(C.size_t(unsafe.Sizeof(C.user_t{}))))
  res, err := u.Lookup(C.GoString(username))
  if err != nil {
  self.uid = nil
  self.gid = nil
  self.username = nil
  self.name = nil
  self.home_dir = nil
  self.error = C.CString(err.Error())
  return self
  }

  self.uid = C.CString(res.Uid)
  self.gid = C.CString(res.Gid)
  self.username = C.CString(res.Username)
  self.name = C.CString(res.Name)
  self.home_dir = C.CString(res.HomeDir)
  self.error = nil
  return self
}

//export user_lookup_id
func user_lookup_id(uid *C.char) *C.user_t {
  self := (*C.user_t) (C.malloc(C.size_t(unsafe.Sizeof(C.user_t{}))))
  res, err := u.LookupId(C.GoString(uid))
  if err != nil {
  self.uid = nil
  self.gid = nil
  self.username = nil
  self.name = nil
  self.home_dir = nil
  self.error = C.CString(err.Error())
  return self
  }

  self.uid = C.CString(res.Uid)
  self.gid = C.CString(res.Gid)
  self.username = C.CString(res.Username)
  self.name = C.CString(res.Name)
  self.home_dir = C.CString(res.HomeDir)
  self.error = nil
  return self
}

//export user_group_ids
func user_group_ids(user *C.user_t) *C.user_group_ids_t {
  us := u.User{
    Uid: C.GoString(user.uid),
    Gid: C.GoString(user.gid),
    Username: C.GoString(user.username),
    Name: C.GoString(user.name),
    HomeDir: C.GoString(user.home_dir),
  }
  self := (*C.user_group_ids_t) (C.malloc(C.size_t(unsafe.Sizeof(C.user_group_ids_t{}))))
  res, err := us.GroupIds()
  if err != nil {
  self.buffer = nil
  self.error = C.CString(err.Error())
  return self
  }

  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))

  slice := (*[1<<30 - 1]*C.char)(array)

  for i, val := range res {
        slice[i] = C.CString(val)
  }
  
  self.buffer = (**C.char) (array)
  self.length = C.size_t(len(res))
  self.error = nil
  return self
}

//export user_clean
func user_clean(self *C.user_t) {
  if self != nil {
    if self.uid != nil {
      C.free(unsafe.Pointer(self.uid))
    }
    if self.gid != nil {
      C.free(unsafe.Pointer(self.gid))
    }
    if self.username != nil {
      C.free(unsafe.Pointer(self.username))
    }
    if self.name != nil {
      C.free(unsafe.Pointer(self.name))
    }
    if self.home_dir != nil {
      C.free(unsafe.Pointer(self.home_dir))
    }
    if self.error != nil {
      C.free(unsafe.Pointer(self.error))
    }
    C.free(unsafe.Pointer(self))
  }
}

//export user_group_clean
func user_group_clean(self *C.user_group_t) {
  if self != nil {
    if self.gid != nil {
      C.free(unsafe.Pointer(self.gid))
    }
    if self.name != nil {
      C.free(unsafe.Pointer(self.name))
    }
    if self.error != nil {
      C.free(unsafe.Pointer(self.error))
    }
    C.free(unsafe.Pointer(self))
  }
}

//export user_group_ids_clean
func user_group_ids_clean(self *C.user_group_ids_t) {
  if self != nil {
    slice := (*[1<<30 - 1]*C.char)(unsafe.Pointer(self.buffer))[:self.length:self.length]
    for i := 0 ; i < len(slice); i++  {
      if slice[i] != nil {
        C.free(unsafe.Pointer(slice[i]))
      }
    }
    C.free(unsafe.Pointer(self.buffer))
    if self.error != nil {
      C.free(unsafe.Pointer(self.error))
    }
    C.free(unsafe.Pointer(self))
  }
}

func main() {}