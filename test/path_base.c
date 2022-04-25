#include <stdio.h>
#include <stdlib.h>
#include <cgo/path/path.h>

int main()
{
    char *path1 = path_base("/a/b");
    char *path2 = path_base("/");
    char *path3 = path_base("");
    printf("%s\n", path1);
    printf("%s\n", path2);
    printf("%s\n", path3);
    free(path1);
    free(path2);
    free(path3);
    return 0;
}