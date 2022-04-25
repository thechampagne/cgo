#include <stdio.h>
#include <cgo/path/path.h>

int main()
{
    printf("%d", path_is_abs("/dev/null"));
    return 0;
}