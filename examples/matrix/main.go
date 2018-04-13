package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andreipimenov/algebra/matrix"
)

func main() {
	// Create new Matrix with 2 rows and 3 columns
	m, _ := matrix.New(2, 3)

	// Input random data into the matrix
	rand.Seed(time.Now().Unix())
	m.Each(func(i, j int, v float64) float64 {
		return rand.Float64()
	})

	// Get new transposed matrix
	t := m.T()

	// Add random number to each element
	n := float64(rand.Intn(100))
	t.Addn(n)

	//Print string representation of matrix
	fmt.Printf("%s\n", t.String())

}
