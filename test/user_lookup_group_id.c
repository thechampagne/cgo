#include <stdio.h>
#include <cgo/os/user.h>

int main()
{
  user_group_t* user = user_lookup_group_id("[ID]");
  if (user->error != NULL)
  {
    printf("Something went wrong: %s", user->error);
    return 1;
  }
  printf("gid: %s\n", user->gid);
  printf("name: %s\n", user->name);
  user_group_clean(user);
  return 0;
}
