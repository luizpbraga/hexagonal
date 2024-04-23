package arithmetic

import (
	"errors"
)

type Arith struct {
}

func New() *Arith {
	return new(Arith)
}

func (adap *Arith) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (adap *Arith) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (adap *Arith) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (adap *Arith) Division(a, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("Division by zero not defined")
	}
	return a / b, nil
}
