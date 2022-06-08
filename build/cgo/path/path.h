#ifndef __CGO_PATH_H__
#define __CGO_PATH_H__


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "path.go"

#include <stdlib.h>

typedef struct {
    int is_match;
    char *error;
} path_match_t;

typedef struct {
    char *dir;
    char *file;
} path_split_t;
 
#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


/**
 * Base returns the last element of path. Trailing slashes are removed before extracting the last element.
 * If the path is empty, Base returns ".". If the path consists entirely of slashes, Base returns "/".
 *
 * Example:
 * * *
 * int main()
 * {
 *  char* path1 = path_base("/a/b");
 *  char* path2 = path_base("/");
 *  char* path3 = path_base("");
 *  printf("%s\n", path1);
 *  printf("%s\n", path2);
 *  printf("%s\n", path3);
 *  free(path1);
 *  free(path2);
 *  free(path3);
 *  return 0;
 *  }
 * * *
 * Output:
 * * *
 * b
 * /
 * .
 * * *
 *
 * @param path
 * @return dynamic string
 */
extern __declspec(dllexport) char* path_base(char* path);

/**
 * Clean returns the shortest path name equivalent to path by purely lexical processing.
 * It applies the following rules iteratively until no further processing can be done:
 *
 * 1. Replace multiple slashes with a single slash.
 * 2. Eliminate each . path name element (the current directory).
 * 3. Eliminate each inner .. path name element (the parent directory)
 *  along with the non-.. element that precedes it.
 * 4. Eliminate .. elements that begin a rooted path:
 *  that is, replace "/.." by "/" at the beginning of a path.
 *
 * The returned path ends in a slash only if it is the root "/".
 * If the result of this process is an empty string, Clean returns the string ".".
 *
 * Example:
 * * *
 * int main()
 * {
 *   char* paths[] = {
 * 		"a/c",
 * 		"a//c",
 * 		"a/c/.",
 * 		"a/c/b/..",
 * 		"/../a/c",
 * 		"/../a/b/../././/c",
 * 		"",
 * 	};
 *
 *   for (int i = 0; i < 7; i++)
 *   {
 *     char* path = path_clean(paths[i]);
 *     printf("Clean(\"%s\") = \"%s\"\n", paths[i], path);
 *     free(path);
 *   }
 *
 *   return 0;
 * }
 * * *
 * Output:
 * * *
 * Clean("a/c") = "a/c"
 * Clean("a//c") = "a/c"
 * Clean("a/c/.") = "a/c"
 * Clean("a/c/b/..") = "a/c"
 * Clean("/../a/c") = "/a/c"
 * Clean("/../a/b/../././/c") = "/a/c"
 * Clean("") = "."
 * * *
 *
 * @param path
 * @return dynamic string
 */
extern __declspec(dllexport) char* path_clean(char* path);

/**
 * Dir returns all but the last element of path, typically the path's directory.
 * After dropping the final element using Split, the path is Cleaned and trailing slashes are removed.
 * If the path is empty, Dir returns ".". If the path consists entirely of slashes followed by non-slash bytes,
 * Dir returns a single slash. In any other case, the returned path does not end in a slash.
 *
 * Example:
 * * *
 * int main()
 * {
 *  char* path1 = path_dir("/a/b/c");
 *  char* path2 = path_dir("a/b/c");
 *  char* path3 = path_dir("/a/");
 *  char* path4 = path_dir("a/");
 *  char* path5 = path_dir("/");
 *  char* path6 = path_dir("");
 *  printf("%s\n", path1);
 *  printf("%s\n", path2);
 *  printf("%s\n", path3);
 *  printf("%s\n", path4);
 *  printf("%s\n", path5);
 *  printf("%s\n", path6);
 *  free(path1);
 *  free(path2);
 *  free(path3);
 *  free(path4);
 *  free(path5);
 *  free(path6);
 *  return 0;
 * }
 * * *
 * Output:
 * * *
 * /a/b
 * a/b
 * /a
 * a
 * /
 * .
 * * *
 *
 * @param path
 * @return dynamic string
 */
extern __declspec(dllexport) char* path_dir(char* path);

/**
 * Ext returns the file name extension used by path.
 * The extension is the suffix beginning at the final dot in
 * the final slash-separated element of path; it is empty if there is no dot.
 *
 * Example:
 * * *
 * int main()
 * {
 *  char* path1 = path_ext("/a/b/c/bar.css");
 *  char* path2 = path_ext("/");
 *  char* path3 = path_ext("");
 *  printf("%s\n", path1);
 *  printf("%s\n", path2);
 *  printf("%s\n", path3);
 *  free(path1);
 *  free(path2);
 *  free(path3);
 *  return 0;
 * }
 * * *
 * Output:
 * * *
 * .css
 *
 *
 * * *
 *
 * @param path
 * @return dynamic string
 */
extern __declspec(dllexport) char* path_ext(char* path);

/**
 * IsAbs reports whether the path is absolute.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d", path_is_abs("/dev/null"));
 *  return 0;
 * }
 * * *
 * Output:
 * * *
 * 1
 * * *
 *
 * @param path
 * @return 1 if it's true and 0 if it's not
 */
extern __declspec(dllexport) int path_is_abs(char* path);

/**
 * Join joins any number of path elements into a single path, separating them with slashes.
 * Empty elements are ignored. The result is Cleaned. However, if the argument list is empty
 * or all its elements are empty, Join returns an empty string.
 *
 * Example:
 * * *
 * int main()
 * {
 *   char* arr1[] = {"a", "b", "c"};
 *   char* arr2[] = {"a", "b/c"};
 *   char* arr3[] = {"a/b", "c"};
 *   char* arr4[] = {"a/b", "../../../xyz"};
 *   char* arr5[] = {"", ""};
 *   char* arr6[] = {"a", ""};
 *   char* arr7[] = {"", "a"};
 *   
 *   char* path1 = path_join(arr1, 3);
 *   char* path2 = path_join(arr2, 2);
 *   char* path3 = path_join(arr3, 2);
 *   char* path4 = path_join(arr4, 2);
 *   char* path5 = path_join(arr5, 2);
 *   char* path6 = path_join(arr6, 2);
 *   char* path7 = path_join(arr7, 2);
 *   
 *   printf("%s\n", path1);
 *   printf("%s\n", path2);
 *   printf("%s\n", path3);
 *   printf("%s\n", path4);
 *   printf("%s\n", path5);
 *   printf("%s\n", path6);
 *   printf("%s\n", path7);
 *   
 *   free(path1);
 *   free(path2);
 *   free(path3);
 *   free(path4);
 *   free(path5);
 *   free(path6);
 *   free(path7);
 *   return 0;
 * }
 * * *
 * Output:
 * * *
 * a/b/c
 * a/b/c
 * a/b/c
 * ../xyz
 *
 * a
 * a
 * * *
 *
 * @param elem array
 * @param length array length
 * @return dynamic string
 */
extern __declspec(dllexport) char* path_join(char** elem, int length);

/**
 * Match reports whether name matches the shell pattern. The pattern syntax is:
 *
 * pattern:
 * 	{ term }
 * term:
 * 	'*'         matches any sequence of non-/ characters
 * 	'?'         matches any single non-/ character
 * 	'[' [ '^' ] { character-range } ']'
 * 	            character class (must be non-empty)
 * 	c           matches character c (c != '*', '?', '\\', '[')
 * 	'\\' c      matches character c
 *
 * character-range:
 * 	c           matches character c (c != '\\', '-', ']')
 * 	'\\' c      matches character c
 * 	lo '-' hi   matches character c for lo <= c <= hi
 *
 * Match requires pattern to match all of name, not just a substring.
 * The only possible returned error is ErrBadPattern, when pattern is malformed.
 *
 * Example:
 * * *
 * int main()
 * {
 *  path_match_t* path = path_match("abc", "abc");
 *  if (path->error != NULL)
 *  {
 *    printf("error: %s", path->error);
 *    path_match_clean(path);
 *    return -1;
 *  }
 *   printf("%d", path->is_match);
 *   path_match_clean(path);
 *   return 0;
 * }
 * * *
 * Output:
 * * *
 * 1
 * * *
 *
 * @param pattern
 * @param name
 * @return path_match_t pointer
 */
extern __declspec(dllexport) path_match_t* path_match(char* pattern, char* name);

/**
 * Split splits path immediately following the final slash,
 * separating it into a directory and file name component.
 * If there is no slash in path, Split returns an empty dir and file set to path.
 * The returned values have the property that path = dir+file.
 *
 * Example:
 * * *
 * int main()
 * {
 *   path_split_t* path = path_split("static/myfile.css");
 *   printf("dir: %s\n", path->dir);
 *   printf("file: %s\n", path->file);
 *   path_split_clean(path);
 *   return 0;
 * }
 * * *
 * Output:
 * * *
 * dir: static/
 * file: myfile.css
 * * *
 *
 * @param path
 * @return path_split_t pointer
 */
extern __declspec(dllexport) path_split_t* path_split(char* path);

/**
 * function to free the memory after using path_match
 *
 * @param self pointer to path_match_t
 */
extern __declspec(dllexport) void path_match_clean(path_match_t* self);

/**
 * function to free the memory after using path_split
 *
 * @param self pointer to path_split_t
 */
extern __declspec(dllexport) void path_split_clean(path_split_t* self);

#ifdef __cplusplus
}
#endif

#endif // __CGO_PATH_H__