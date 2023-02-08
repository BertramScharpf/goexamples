/*
 *  diverse/endrek.c  --  Endrekursion
 */

#include <stdio.h>

void showit( int i)
{
    printf( "%d\n", i);
    if (i > 0) {
        i--;
        showit( i);
    }
}

int main( int argc, char **argv)
{
    showit( argc);
    return 0;
}
