package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

func GenerateMatrix(rows, cols int) []float64 {
	mat := make([]float64, rows*cols)
	for i := range mat {
		mat[i] = float64(rand.Int()%4) * math.Pow(-1, float64(rand.Int()%2))
	}
	return mat
}

func CustomMatrixMultiplication(a, b, c []float64, m, n, k int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for l := 0; l < k; l++ {
				c[i*n+j] += a[i*k+l] * b[l*n+j]
			}
		}
	}
}

func main() {
	// defer profile.Start(profile.ProfilePath("/tmp")).Stop()

	sizeArray := []int{100, 200, 300, 400, 500, 600, 700, 800, 900,
		1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000}

	fmt.Println("size, time")
	for _, size := range sizeArray {
		a := GenerateMatrix(size, size)
		b := GenerateMatrix(size, size)
		c := make([]float64, size*size)

		ag := blas64.General{
			Rows:   size,
			Cols:   size,
			Stride: size,
			Data:   a,
		}
		bg := blas64.General{
			Rows:   size,
			Cols:   size,
			Stride: size,
			Data:   b,
		}
		cg := blas64.General{
			Rows:   size,
			Cols:   size,
			Stride: size,
			Data:   c,
		}

		tic := time.Now()
		blas64.Gemm(blas.NoTrans, blas.NoTrans, 1, ag, bg, 0, cg)
		toc := time.Now()

		// print time in seconds
		fmt.Printf("%d, %v\n", size, float64(toc.Sub(tic).Microseconds())/1000000.0)
	}
}
