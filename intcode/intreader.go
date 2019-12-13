package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// IntReader is an input reader for ints
type IntReader interface {
	ReadInt() (int, error)
}

// SimpleIntReader implements an IntReader that will provide to readers defined values. The values this reader provides
// are set when calling NewSimpleIntReader.
type SimpleIntReader struct {
	pos    int
	buffer []int
}

// NewSimpleIntReader returns a new SimpleIntReader that will provide to reader the values passed as arguments
func NewSimpleIntReader(values ...int) SimpleIntReader {
	var res SimpleIntReader
	res.buffer = append(res.buffer, values...)
	return res
}

// ReadInt read a value from this reader. Implements the IntReader interface.
func (r *SimpleIntReader) ReadInt() (int, error) {
	if r.pos >= len(r.buffer) {
		return 0, fmt.Errorf("no input to read")
	}
	value := r.buffer[r.pos]
	r.pos++
	return value, nil
}

// ParseInputString attempts to parse a string with a comma separated list of integers into a slice of inputs
func ParseInputString(str string) ([]int, error) {
	var res []int

	strValues := strings.Split(str, ",")

	for _, v := range strValues {
		strValue := strings.TrimSpace(v)
		v, err := strconv.Atoi(strValue)
		if err != nil {
			return nil, fmt.Errorf("converting %s to integer: %w", v, err)
		}
		res = append(res, v)
	}

	return res, nil
}
