#include <stdio.h>
#include <cgo/path/path.h>

int main()
{
    path_split_t *path = path_split("static/myfile.css");
    printf("dir: %s\n", path->dir);
    printf("file: %s\n", path->file);
    path_split_clean(path);
    return 0;
}