#include <stdio.h>
#include <cgo/os/user.h>

int main()
{
  user_t* user = user_current();
  if (user->error != NULL)
  {
    printf("Something went wrong: %s", user->error);
    return 1;
  }
  user_group_ids_t* gids = user_group_ids(user);
  if (gids->error != NULL)
  {
    printf("Something went wrong: %s", gids->error);
    return 1;
  }
  for (size_t i = 0; i < gids->length; i++)
  {
    printf("%s\n", gids->buffer[i]);
  }
  user_group_ids_clean(gids);
  user_clean(user);
  return 0;
}
