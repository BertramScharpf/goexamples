/*
 *  dbls.c  --  Rechnen mit Double
 */

#include <stdio.h>

int main( int argc, char **argv)
{
    double d;

    scanf( "%lg", &d);
    printf( "%lg^2 = %lg\n", d, d*d);
    return 0;
}

