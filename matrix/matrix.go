// Package matrix implements operations with matrices in golang
package matrix

import (
	"bytes"
	"fmt"
)

// Matrix is a basic type for 2-dimentional matrices
// which consists of rows, columns and slice of elements
type Matrix struct {
	rows int
	cols int
	data []float64
}

// New returns pointer to the new empty matrix with given dimentions
func New(rows, cols int) (*Matrix, error) {
	if rows < 0 || cols < 0 {
		return nil, fmt.Errorf("Dimetions %dx%d must not being negative", rows, cols)
	}
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}, nil
}

// String returns string representation of the matrix
func (m *Matrix) String() string {
	b := &bytes.Buffer{}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Fprintf(b, "%-15.3f", m.get(i, j))
		}
		fmt.Fprintf(b, "\n")
	}
	return b.String()
}

// Dimentions returns count of rows and columns of the matrix
func (m *Matrix) Dimentions() (int, int) {
	return m.rows, m.cols
}

func (m *Matrix) checkRange(i, j int) error {
	if i < 0 || j < 0 {
		return fmt.Errorf("Position (%d, %d) must not being negative", i, j)
	}
	if i >= m.rows || j >= m.cols {
		return fmt.Errorf("Position (%d, %d) is out of the range (0:%d, 0:%d)", i, j, m.rows-1, m.cols-1)
	}
	return nil
}

func (m *Matrix) get(i, j int) float64 {
	return m.data[m.cols*i+j]
}

func (m *Matrix) set(i, j int, v float64) {
	m.data[m.cols*i+j] = v
}

// Get returns the value of (i, j)
func (m *Matrix) Get(i, j int) (float64, error) {
	if err := m.checkRange(i, j); err != nil {
		return 0, err
	}
	return m.get(i, j), nil
}

// Set sets the value at (i, j)
func (m *Matrix) Set(i, j int, v float64) error {
	if err := m.checkRange(i, j); err != nil {
		return err
	}
	m.set(i, j, v)
	return nil
}

// Each applies function to every element in the matrix
func (m *Matrix) Each(f func(i, j int, v float64) float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.set(i, j, f(i, j, m.get(i, j)))
		}
	}
}

// T returns new transposed matrix
func (m *Matrix) T() *Matrix {
	t := &Matrix{
		rows: m.cols,
		cols: m.rows,
		data: make([]float64, m.rows*m.cols),
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			t.set(j, i, m.get(i, j))
		}
	}
	return t
}
