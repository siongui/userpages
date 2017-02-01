#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  char ptr[100];
  strcpy(ptr, "abcd\n");
  printf("%s", ptr);
  removeStringTrailingNewline(ptr);
  printf("%s", ptr);
  return 0;
}
