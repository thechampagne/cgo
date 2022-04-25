#include <stdio.h>
#include <cgo/path/path.h>

int main()
{
    path_match_t *path = path_match("abc", "abc");
    if (path->error != NULL) {
        printf("error: %s", path->error);
        path_match_clean(path);
        return -1;
    }
    printf("%d", path->is_match);
    path_match_clean(path);
    return 0;
}