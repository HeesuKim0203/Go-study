package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert.Equal(t, 81, square(9), "Square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert.Equal(t, 9, square(3), "Square(3) should be 9")
}

func TestSquare3(t *testing.T) {
	assert.Equal(t, 1, square(0), "Square(0) should be 9")
}
