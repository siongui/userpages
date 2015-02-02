#include <stdio.h>
#include <string.h>

void removeStringTrailingNewline(char *str) {
  if (str == NULL)
    return;
  int length = strlen(str);
  if (str[length-1] == '\n')
    str[length-1]  = '\0';
}
