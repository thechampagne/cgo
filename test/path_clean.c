#include <stdio.h>
#include <stdlib.h>
#include <cgo/path/path.h>

int main() {
    char *paths[] = {
            "a/c",
            "a//c",
            "a/c/.",
            "a/c/b/..",
            "/../a/c",
            "/../a/b/../././/c",
            "",
    };
    for (int i = 0; i < 7; i++) {
        char *path = path_clean(paths[i]);
        printf("Clean(\"%s\") = \"%s\"\n", paths[i], path);
        free(path);
    }
    return 0;
}
