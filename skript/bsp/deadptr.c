/*
 *  deadptr.c  --  Pointer to non-existing space
 */

#include <stdio.h>

int *gimme( int i)
{
    int l;
    int *p;

    p = &l;
    l = i*2;
    return p;
}

int main( int argc, char **argv)
{
    int *p, *q;

    p = gimme( 7);
    printf( "p=%p *p=%d\n", p, *p);
    q = gimme( 9);
    printf( "q=%p *q=%d\n", q, *q);
    printf( "p=%p *p=%d\n", p, *p);
    return 0;
}

