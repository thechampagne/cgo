#include <stdio.h>
#include <cgo/os/user.h>

int main()
{
  user_t* user = user_lookup("[USERNAME]");
  if (user->error != NULL)
  {
    printf("Something went wrong: %s", user->error);
    return 1;
  }
  printf("uid: %s\n", user->uid);
  printf("gid: %s\n", user->gid);
  printf("username: %s\n", user->username);
  printf("name: %s\n", user->name);
  printf("home_dir: %s\n", user->home_dir);
  user_clean(user);
  return 0;
}
