package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivideByZero = errors.New("can't divide by 0")
var ErrNumTooBig = errors.New("num too large")

type OpError struct {
	Op string
	Code string
	Message string
	Time time.Time
}

func NewOpError(op, code, message string, time time.Time) *OpError {
	return &OpError{
		Op: op,
		Code: code,
		Message: message,
		Time: time,
	}
}

func (op OpError) Error() string { // implement the error interface :siuuuu:
	return string(op.Message)
}

func DoSomething() error {
	return NewOpError("do something", "do_something", "wtf do something", time.Now())
}

func divide(a, b int) (int, error) { // by convetion the error return value is the last one in the list of multiple return values
	if b == 0 {
		return 0, ErrDivideByZero
	}
	if a > 1000 {
		return 0, ErrNumTooBig
	}
	return a / b, nil
}

func main() {
	result, err := divide(100000, 5)
	if err != nil {
		fmt.Println(err)
		if err == ErrDivideByZero {
			fmt.Println("Come on! we all know that")
		}
	} else {
		fmt.Println(result)
	}
}
