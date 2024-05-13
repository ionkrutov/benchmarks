package main

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

var a []float64
var b []float64
var c []float64

func init() {
	a = GenerateMatrix(1000, 1000)
	b = GenerateMatrix(1000, 1000)
	c = make([]float64, 1000*1000)
}

func BenchmarkMatrixMult(bt *testing.B) {
	ag := blas64.General{
		Rows:   1000,
		Cols:   1000,
		Stride: 1000,
		Data:   a,
	}
	bg := blas64.General{
		Rows:   1000,
		Cols:   1000,
		Stride: 1000,
		Data:   b,
	}
	cg := blas64.General{
		Rows:   1000,
		Cols:   1000,
		Stride: 1000,
		Data:   c,
	}

	for i := 0; i < bt.N; i++ {
		blas64.Gemm(blas.NoTrans, blas.NoTrans, 1, ag, bg, 0, cg)
	}
}

func BenchmarkMatrixMultCustom(bt *testing.B) {
	for i := 0; i < bt.N; i++ {
		CustomMatrixMultiplication(a, b, c, 1000, 1000, 1000)
	}
}

func TestCompareMatrixMult(t *testing.T) {
	ag := blas64.General{
		Rows:   1000,
		Cols:   1000,
		Stride: 1000,
		Data:   a,
	}
	bg := blas64.General{
		Rows:   1000,
		Cols:   1000,
		Stride: 1000,
		Data:   b,
	}
	cg := blas64.General{
		Rows:   1000,
		Cols:   1000,
		Stride: 1000,
		Data:   c,
	}

	blas64.Gemm(blas.NoTrans, blas.NoTrans, 1, ag, bg, 0, cg)

	res1 := cg.Data

	res2 := make([]float64, 1000*1000)
	CustomMatrixMultiplication(a, b, res2, 1000, 1000, 1000)
	for i := range res1 {
		if res1[i] != res2[i] {
			t.Errorf("Expected %f, got %f", res2[i], res1[i])
		}
	}
}
