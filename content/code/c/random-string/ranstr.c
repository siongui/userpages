#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* alphabet: [a-z0-9] */
const char alphabet[] = "abcdefghijklmnopqrstuvwxyz0123456789";

/**
 * not a cryptographically secure number
 * return interger [0, n).
 */
int intN(int n) { return rand() % n; }

/**
 * Input: length of the random string [a-z0-9] to be generated
 */
char *randomString(int len) {
  char *rstr = malloc((len + 1) * sizeof(char));
  int i;
  for (i = 0; i < len; i++) {
    rstr[i] = alphabet[intN(strlen(alphabet))];
  }
  rstr[len] = '\0';
  return rstr;
}

int main(int argc, char **argv) {
  // the seed for a new sequence of pseudo-random integers
  // to be returned by rand()
  srand(time(NULL));

  char *p;
  p = randomString(10);
  printf("%s\n", p);
  free(p);

  p = randomString(11);
  printf("%s\n", p);
  free(p);

  p = randomString(12);
  printf("%s\n", p);
  free(p);
}
