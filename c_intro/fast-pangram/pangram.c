#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

bool ispangram(char *s) {

    // save init pointer.
    char *f = s;

    /*
        printf("int %d\n", s);
        printf("dereferenced int %d\n", *s);
        printf("dereferenced string %s\n", *s);
        printf("string %s\n", s);
    */

    printf("\n%s\n", f);

    for (char *p = s; *p > 0 ; p++) {
        // to lowercase
        // inspired by https://stackoverflow.com/a/2661917/3478120
        //
        *p =  *p > 0x40 && *p < 0x5b ? *p | 0x20 : *p;



    }

    // should be equal to 26

    printf("\n%s\n", f);

  // TODO implement this!
  return false;
}

int main() {
  size_t len;
  ssize_t read;
  char *line = NULL;

  while ((read = getline(&line, &len, stdin)) != -1) {
    if (ispangram(line)) {
      //printf("%s", line);
    }

    return 0;
  }

  if (ferror(stdin)) {
    fprintf(stderr, "Error reading from stdin");
  }

  free(line);

  fprintf(stderr, "ok\n");
}

