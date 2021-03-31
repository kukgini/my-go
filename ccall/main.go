package main

/*
#include <stdio.h>
int testc(int a, int b) {
    printf("hi, from c code: result = %d\n", a+b);
}
*/
import "C"

func main() {
    C.testc(1,3)
}
