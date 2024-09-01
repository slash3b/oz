# include <stdio.h>

typedef struct {
    char* str;
    int age;
} Abc;

int main() {
    printf("hell0\n");


    int a[] = {3, 2, 1, 0};

    int len = sizeof(a)/sizeof(int);

    printf("sizeof %d\n", len);

    for ( int i = 0; i < len; i++ ) {

        printf("%d\n", a[i]);


    }

    Abc abc = {"bye", 85};
    
    printf("%s\n", abc.str);

    Abc* foo = &abc;
    printf("%p\n", foo);
    printf("str: %s\n", foo->str);
    printf("age: %d\n", foo->age);

    return 0;
}
