#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

bool ispangram(char *s) {
    int res = 0;

    char ch = 0;

    for (char *p = s; *p > 0 ; p++) {
        // to lowercase
        // inspired by https://stackoverflow.com/a/2661917/3478120
        ch = *p > 0x40 && *p < 0x5b ? *p | 0x20 : *p;

        if (ch > 0x60 && ch < 0x7b) {
            res |=  1 << ch - 0x61;
        }
  }

  return res == 0b11111111111111111111111111;
}

int main() {
  size_t len;
  ssize_t read;
  char *line = NULL;

  while ((read = getline(&line, &len, stdin)) != -1) {
    if (ispangram(line)) {
      printf("%s", line);
    }

  }

  if (ferror(stdin)) {
    fprintf(stderr, "Error reading from stdin");
  }

  free(line);

  fprintf(stderr, "ok\n");
}

