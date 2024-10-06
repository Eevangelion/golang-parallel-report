package main

import (
	"fmt"
	"math/rand/v2"

	utils "github.com/Eevangelion/golang-parallel-report/utils"
)

type Line []float64
type Matrix []Line

func MultiplyMatrix(a Matrix, b Matrix) Matrix {
	defer utils.Timer("matrix multiplication")()
	ans := make(Matrix, len(a))
	for i := range ans {
		ans[i] = make(Line, len(b[0]))
	}
	for i := range a {
		for j := range b[0] {
			for k := 0; k < len(a); k++ {
				ans[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return ans
}

func MultiplyMatrixWithGoroutines(a Matrix, b Matrix) Matrix {
	defer utils.Timer("matrix multiplication with goroutines")()
	ans := make(Matrix, len(a))
	for i := range ans {
		ans[i] = make(Line, len(b[0]))
	}
	for i := range a {
		for j := range b[0] {
			go func() {
				val := 0.
				for k := 0; k < len(a); k++ {
					val += a[i][k] * b[k][j]
				}
				ans[i][j] = val
			}()
		}
	}
	return ans
}

func FillMatrix(m Matrix) {
	for i := range m {
		for j := range m[i] {
			m[i][j] = rand.Float64()
		}
	}
}

func PrintMatrix(m Matrix, name string) {
	fmt.Println("Matrix:", name)
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%d ", int(m[i][j]))
		}
		fmt.Println()
	}
}

func test1() {
	matrixA := Matrix{{1, 5, 3}, {5, 2, 3}, {4, 2, 2}}
	matrixB := Matrix{{4, 3}, {1, 3}, {5, 5}}
	PrintMatrix(matrixA, "A")
	PrintMatrix(matrixB, "B")
	ans := MultiplyMatrixWithGoroutines(matrixA, matrixB)

	PrintMatrix(ans, "A x B")
}

func test2() {
	matrixA := make(Matrix, 500)
	for i := range matrixA {
		matrixA[i] = make(Line, 500)
	}
	matrixB := make(Matrix, 500)
	for i := range matrixB {
		matrixB[i] = make(Line, 500)
	}
	FillMatrix(matrixA)
	FillMatrix(matrixB)
	MultiplyMatrix(matrixA, matrixB)
	MultiplyMatrixWithGoroutines(matrixA, matrixB)
}

func main() {
	test2()
}
