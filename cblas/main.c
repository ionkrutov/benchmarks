#include <stdio.h>
#include "cblas.h"
#include <stdlib.h>
#include <time.h>
#include <memory.h>
#include <math.h>

// gcc -O2 main.c -lcblas -lm
// ./a.out

void generate_random_matrix(int n, int m, double* matrix) {
    for (int i = 0; i < n * m; i++) {
        matrix[i] = (rand() % 4) * pow(-1, rand() % 2);
    }
}

void print_matrix(int n, int m, double* matrix) {
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            printf("%f ", matrix[i * m + j]);
        }
        printf("\n");
    }
}


int main() {
    int size_array[] = {100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, \
        2000, 3000, 4000, 5000, 6000, 7000, 8000};
    
    printf("size, time\n");
    for (int i = 0; i < 17; i++) {
        int n = size_array[i];
        int m = n;
        double * a = (double *)malloc(n * m * sizeof(double));
        double * b = (double *)malloc(n * m * sizeof(double));
        double * c = (double *)malloc(n * m * sizeof(double));

        generate_random_matrix(n, m, a);
        generate_random_matrix(n, m, b);

        time_t start_t, stop_t;
        start_t = clock();
        cblas_dgemm(CblasRowMajor, CblasNoTrans, CblasNoTrans, n, m, m, 1.0, a, m, b, m, 0.0, c, m);    
        stop_t = clock();

        printf("%d, %f\n", n, (double)(stop_t - start_t) / CLOCKS_PER_SEC);
        free(a);
        free(b);
        free(c);
    }

    return 0;
}