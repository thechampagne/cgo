#include <stdio.h>
#include <stdlib.h>
#include <cgo/path/path.h>

int main()
{
    char *path1 = path_dir("/a/b/c");
    char *path2 = path_dir("a/b/c");
    char *path3 = path_dir("/a/");
    char *path4 = path_dir("a/");
    char *path5 = path_dir("/");
    char *path6 = path_dir("");
    printf("%s\n", path1);
    printf("%s\n", path2);
    printf("%s\n", path3);
    printf("%s\n", path4);
    printf("%s\n", path5);
    printf("%s\n", path6);
    free(path1);
    free(path2);
    free(path3);
    free(path4);
    free(path5);
    free(path6);
    return 0;
}