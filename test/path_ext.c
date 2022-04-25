#include <stdio.h>
#include <stdlib.h>
#include <cgo/path/path.h>

int main() {
    char *path1 = path_ext("/a/b/c/bar.css");
    char *path2 = path_ext("/");
    char *path3 = path_ext("");
    printf("%s\n", path1);
    printf("%s\n", path2);
    printf("%s\n", path3);
    free(path1);
    free(path2);
    free(path3);
    return 0;
}