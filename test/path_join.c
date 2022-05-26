#include <stdio.h>
#include <stdlib.h>
#include <cgo/path/path.h>

int main()
{
   char* arr1[] = {"a", "b", "c"};
   char* arr2[] = {"a", "b/c"};
   char* arr3[] = {"a/b", "c"};
   char* arr4[] = {"a/b", "../../../xyz"};
   char* arr5[] = {"", ""};
   char* arr6[] = {"a", ""};
   char* arr7[] = {"", "a"};
   
   char* path1 = path_join(arr1, 3);
   char* path2 = path_join(arr2, 2);
   char* path3 = path_join(arr3, 2);
   char* path4 = path_join(arr4, 2);
   char* path5 = path_join(arr5, 2);
   char* path6 = path_join(arr6, 2);
   char* path7 = path_join(arr7, 2);
   
   printf("%s\n", path1);
   printf("%s\n", path2);
   printf("%s\n", path3);
   printf("%s\n", path4);
   printf("%s\n", path5);
   printf("%s\n", path6);
   printf("%s\n", path7);
   
   free(path1);
   free(path2);
   free(path3);
   free(path4);
   free(path5);
   free(path6);
   free(path7);
   return 0;
}