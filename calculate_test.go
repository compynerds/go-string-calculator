package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateNoSubtraction(t *testing.T) {
	expr := "1+2+3"
	result, err := calculate(expr)
	assert.Nil(t, err)
	assert.Equal(t, 6, result)
}

func TestCalculateError(t *testing.T) {
	expr := "1+2+3+one"
	result, err := calculate(expr)
	assert.NotNil(t, err)
	assert.Equal(t, 0, result)
}

func TestCalculateNegatives(t *testing.T) {
	expr := "-1+-2+-3"
	result, err := calculate(expr)
	assert.Nil(t, err)
	assert.Equal(t, -6, result)
}

func TestCalculateBasicAdditionAndSubtraction(t *testing.T) {
	expr := "1-2+3"
	result, err := calculate(expr)
	assert.Nil(t, err)
	assert.Equal(t, 2, result)
}

func TestCalculateLong(t *testing.T) {
	expr := "1-2+3-2+10-12+16+1-1"
	result, err := calculate(expr)
	assert.Nil(t, err)
	assert.Equal(t, 14, result)
}

func TestCalculateExcessiveSpacing(t *testing.T) {
	expr := "1   -2   +3   -2   +10   -12   +16   +1   -1"
	result, err := calculate(expr)
	assert.Nil(t, err)
	assert.Equal(t, 14, result)
}

func TestOrderCalc(t *testing.T) {
	exp := "(1+(4+5+2)-3)+(6+8)"
	result, err := calculateOrderOfOperations(exp)
	assert.Nil(t, err)
	assert.Equal(t, 23, result)
}

func TestOrderCalcNegative(t *testing.T) {
	exp := "(1-(4+5+2)-3)+(6+8)"
	result, err := calculateOrderOfOperations(exp)
	assert.Nil(t, err)
	assert.Equal(t, 1, result)
}
