// Package matrix implements operations with matrices in golang
package matrix

import (
	"bytes"
	"fmt"
	"sync"
)

// Matrix is a basic type for 2-dimentional matrices
// which consists of rows, columns and slice of elements
type Matrix struct {
	rows int
	cols int
	data []float64
	sync.RWMutex
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

// Clone returns new cloned matrix
func (m *Matrix) Clone() *Matrix {
	c := &Matrix{
		rows: m.rows,
		cols: m.cols,
		data: make([]float64, m.rows*m.cols),
	}
	copy(c.data, m.data)
	return c
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

func (m *Matrix) checkEqualDimentions(x *Matrix) error {
	if m.rows != x.rows || m.cols != x.cols {
		return fmt.Errorf("Dimentions of two matrices %dx%d and %dx%d are not equal", m.rows, m.cols, x.rows, x.cols)
	}
	return nil
}

func (m *Matrix) get(i, j int) float64 {
	m.RLock()
	defer m.RUnlock()
	return m.data[m.cols*i+j]
}

func (m *Matrix) set(i, j int, v float64) {
	m.Lock()
	m.data[m.cols*i+j] = v
	m.Unlock()
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

// Add adds the matrix
func (m *Matrix) Add(x *Matrix) error {
	if err := m.checkEqualDimentions(x); err != nil {
		return err
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.set(i, j, m.get(i, j)+x.get(i, j))
		}
	}
	return nil
}

// Sub subtracts the matrix
func (m *Matrix) Sub(x *Matrix) error {
	if err := m.checkEqualDimentions(x); err != nil {
		return err
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.set(i, j, m.get(i, j)-x.get(i, j))
		}
	}
	return nil
}

// Addn adds number to every element in the matrix
func (m *Matrix) Addn(n float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.set(i, j, m.get(i, j)+n)
		}
	}
}

// Scale scales matrix with given factor
func (m *Matrix) Scale(n float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.set(i, j, m.get(i, j)*n)
		}
	}
}

// Dot returns dot product of matrices
func (m *Matrix) Dot(x *Matrix) (float64, error) {
	if err := m.checkEqualDimentions(x); err != nil {
		return 0, err
	}
	r := float64(0)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			r += m.get(i, j) * x.get(i, j)
		}
	}
	return r, nil
}
