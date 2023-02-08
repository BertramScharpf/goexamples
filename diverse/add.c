/*
 *  diverse/add.c  --  Ganzen addieren
 */

#include <stdio.h>

int main( int argc, char **argv)
{
    double i = 99;
    i += argc;
    printf( "%f\n", i);
    return 0;
}
