package main

import "fmt"

func sumOfVectors() {
	var v1 [20]float64
	var v2 [20]float64
	var v3 [20]float64

	for i := 0; i < 20; i++ {
		v1[i] = 2.0
		v2[i] = 3.0
	}
	for i := 0; i < 20; i++ {
		v3[i] = v1[i] + v2[i]
	}

	fmt.Println(v3)
}

/*
hadamard product is operation on two vectors of the same dimension,
which produces another vector of the same dimension and the result
is product of corresponding elements of the two vectors
for example for vectors v1 = [1,2,3] and v2 = [1,2,3] the result is
v3 = [1*1, 2*2, 3*3] = [1,4,9]
*/

func hadamardProduct() {
	var x1 []float64 = []float64{1, 2, 3, 4, 5}
	var x2 []float64 = []float64{1, 2, 3, 4, 5}
	var x3 []float64

	for i := 0; i < 5; i++ {
		x3 = append(x3, x1[i]*x2[i])
	}

	fmt.Println(x3)
}

func twoDimensionalMatrixSlice() {
	var x [][]float64 = [][]float64{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				fmt.Println(x[i][j])
			} else {
				fmt.Print(x[i][j], " ")

			}
		}
	}
}

func inverseMatrix() {

}

func addMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {
	rows1, cols1 := len(matrix1), len(matrix1[0])
	rows2, cols2 := len(matrix2), len(matrix2[0])

	if rows1 != rows2 || cols1 != cols2 {
		panic("Matrices are not the same size")
	}

	result := make([][]int, rows1)
	for i := 0; i < rows1; i++ {
		result[i] = make([]int, cols1)
		for j := 0; j < cols1; j++ {
			result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return result
}

func multiplyMatrices(m1 [][]int, m2 [][]int) [][]int {
	r1, c1 := len(m1), len(m1[0])
	r2, c2 := len(m2), len(m2[0])

	if r1 != c2 || c1 != r2 {
		panic("The matrices are wrong size")
	}

	res := make([][]int, r1)
	for i := 0; i < r1; i++ {
		res[i] = make([]int, c2)
		for j := 0; j < c2; j++ {
			sum := 0
			for k := 0; k < c1; k++ {
				sum += m1[i][k] * m2[k][j]
			}
			res[i][j] = sum
		}
	}

	return res
}

// check if multiplying matrices is commutative
func justChecking(m1 [][]int, m2 [][]int) bool {
	arr1 := multiplyMatrices(m1, m2)
	arr2 := multiplyMatrices(m2, m1)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if len(arr1[i]) != len(arr2[i]) {
			return false
		}

		for j := 0; j < len(arr1[i]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	sumOfVectors()
	hadamardProduct()
	twoDimensionalMatrixSlice()

	var a1 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var a2 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	var q1 [][]int = [][]int{{1, 2, 3}, {1, 2, 3}}
	var q2 [][]int = [][]int{{1, 2}, {1, 2}, {1, 2}}

	fmt.Println(addMatrices(a1, a2))
	fmt.Println(multiplyMatrices(a1, a2))
	fmt.Println(justChecking(a1, a2)) //true
	fmt.Println(justChecking(q1, q2)) //false

}
