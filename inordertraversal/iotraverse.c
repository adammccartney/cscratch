#include <error.h>
#include <stdlib.h>
#include <stdio.h>
#include "abinarytree.h"

#define STACKSIZE sizeof(binarytree) * 16

static stack* initStack(int size) {
    stack* s = malloc(sizeof(stack));
    if (!s) {
        perror("malloc");
        exit(1);
    }
    s->size = size;
    s->top = -1;
    
    s->tree = malloc(s->size * sizeof(s->tree));
    if (!s->tree) {
        perror("malloc");
        exit(1);
    }
    return s;
}

void freeStack(stack* s) {
    free(s->tree);
    free(s);
}

static binarytree* initTree(char c) {
    binarytree* t = malloc(sizeof(binarytree));
    if (!t) {
        perror("malloc");
        exit(1);
    }
    t->c = c;
    t->left = NULL;
    t->right = NULL;
    return t;
}


int main(void) {
    stack* s = initStack(STACKSIZE);

    binarytree* A = NULL;
    binarytree* B = NULL;
    binarytree* C = NULL;
    binarytree* D = NULL;
    binarytree* E = NULL;
    binarytree* F = NULL;
    binarytree* G = NULL;
    binarytree* H = NULL;
    binarytree* J = NULL;

    A = initTree('A');
    B = initTree('B');
    D = initTree('D');
    C = initTree('C');
    E = initTree('E');
    G = initTree('G');
    F = initTree('F');
    H = initTree('H');
    J = initTree('J');

    A->left = B;
    A->right = C;
    B->left = D;
    C->left= E;
    C->right = F;
    E->right = G;
    F->left=H;
    F->right=J;
    
    traverseInorder(A, s);
    free(A);
    free(B);
    free(C);
    free(D);
    free(E);
    free(F);
    free(G);
    free(H);
    free(J);
    freeStack(s);

    return 0;
}

