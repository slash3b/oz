#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#define STARTING_CAPACITY 8

typedef struct DA {
    void** items;
    int len;
    int cap;
} DA;


DA* DA_new (void) {
    DA* da = malloc(sizeof(DA));

    da->items = malloc(STARTING_CAPACITY * sizeof(void*));
    da->len = 0;
    da->cap = STARTING_CAPACITY;

    return da;
}

void DA_free(DA *da) {
  free(da->items);
  free(da);
}

int DA_size(DA *da) {
    return da->len;
}

void DA_push (DA* da, void* x) {
    if (da->len == da->cap) {
        da->cap <<= 1;

        da->items = realloc(da->items, (da->cap * sizeof(void*)));
        printf("resized to %d\n", da->cap);
    }

    da->items[da->len++] = x;
}

void* DA_pop(DA *da) {
    if (da->len == 0) {return NULL;}

    return da->items[--da->len];
}

void* DA_get(DA *da, int i) {
    if (i >= 0 && i <= da->len) {
        return da->items[i];
    }

    return NULL;
}

void DA_set(DA *da, void* x, int i) {
    if (i >= 0 && i <= da->len) {
        da->items[i] = x;
    }
}

int main() {
    DA* da = DA_new();

    assert(DA_size(da) == 0);

    // basic push and pop test
    int x = 5;
    float y = 12.4;
    DA_push(da, &x);
    DA_push(da, &y);
    assert(DA_size(da) == 2);

    assert(DA_pop(da) == &y);
    assert(DA_size(da) == 1);

    assert(DA_pop(da) == &x);
    assert(DA_size(da) == 0);
    assert(DA_pop(da) == NULL);

    // basic set/get test
    DA_push(da, &x);
    DA_set(da, &y, 0);
    assert(DA_get(da, 0) == &y);
    DA_pop(da);
    assert(DA_size(da) == 0);

    // expansion test
    DA *da2 = DA_new(); // use another DA to show it doesn't get overriden
    DA_push(da2, &x);

    // in one line variables i, n and arr are declared
    int i, n = 100 * STARTING_CAPACITY, arr[n];

    for (i = 0; i < n; i++) {
      arr[i] = i;
      DA_push(da, &arr[i]);
    }

    assert(DA_size(da) == n);
    for (i = 0; i < n; i++) {
      assert(DA_get(da, i) == &arr[i]);
    }

    for (; n; n--) {
      DA_pop(da);
    }

    assert(DA_size(da) == 0);
    assert(DA_pop(da2) == &x); // this will fail if da doesn't expand

    DA_free(da);
    DA_free(da2);
    printf("OK\n");
}
