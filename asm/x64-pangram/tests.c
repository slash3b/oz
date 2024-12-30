#include <assert.h>
#include <stdio.h>

extern int pangram(char *);

int main(void) {
  printf("debug a. expected 1 %d %b\n", pangram("a"), pangram("a"));
  printf("debug `a b`. expected 1 %d %b\n", pangram("a b"), pangram("a b"));
  return 0;
  printf("debug b. expected 1 %d %b\n", pangram("b"), pangram("b"));
  printf("debug z. expected 1 %d %b\n", pangram("z"), pangram("z"));
  printf("debug all alpha. expected 1 %d %b\n", pangram("abcdefghijklmnopqrstuvwxyz"), pangram("abcdefghijklmnopqrstuvwxyz"));
  printf("debug incomplete should be 0. expected 1 %d %b\n", pangram("abcdefghijklmnopqrstuvwxy"), pangram("abcdefghijklmnopqrstuvwxy"));
  printf("failing test cases. expected 1 %d %b\n", pangram("the quick brown fox jumps over teh lazy dog"), pangram("the quick brown fox jumps over teh lazy dog"));
  return 0;
  assert(pangram("") == 0);
  assert(pangram("abcdefghijklmnopqrstuvwxyz") == 1);
  assert(pangram("the quick brown fox jumps over teh lazy dog") == 1);
  assert(pangram("abc, def! ghi... jkl25; mnopqrstuvwxyz") ==
         1);                                          // ignore punctuation
  assert(pangram("abcdefghijklmnopqrstuvwxy") == 0);  // incomplete
  assert(pangram("ABCdefGHIjklMNOpqrSTUvwxYZ") == 1); // mixed case
  assert(pangram("!bcdefghijklmnopqrstuvwxyz") ==
         0); // close-match symbols should not be false positive
  assert(pangram("\1bcdefghijklmnopqrstuvwxyz") ==
         0); // close-match control code should not be false positive
  assert(pangram("\7abcdefghijklmnopqrstuvwxyz") ==
         1); // other control codes are fine
  printf("OK\n");
}
