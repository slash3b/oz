#include <assert.h>
#include <stdio.h>

extern int fib(int n);
extern int fac(int n);

int main(void) {
  printf("fac(3) expected 6, got: %d\n", fac(3));
  printf("fic(10) expected 3_628_800 %d\n", fac(10));
  return 0;
  printf("fib(3) expected 6 %d\n", fib(3));
  return 0;

  assert(fib(0) == 0);
  assert(fib(1) == 1);
  printf("fib(3) expected 2 %d\n", fib(3));
  printf("fib(10) expected 55 %d\n", fib(10));
  return 0;
  assert(fib(2) == 1);
  assert(fib(3) == 2);
  assert(fib(10) == 55);
  assert(fib(12) == 144);
  printf("OK\n");
}
