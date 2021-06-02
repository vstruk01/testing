#include <stdio.h>
 
enum TEST {
    GENERATE_ENUM("dd")
};

int main(void) {
    printf("-> %d", GENERATE_ENUM);
    return 0;
}