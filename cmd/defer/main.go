package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	err := testFunction()
	EmitError(err)

	err = otherTestFunction()
	EmitError(err)

	printDeferStackOrder()
}

func testFunction() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w: touched", err)
		}
	}()

	err = fmt.Errorf("failed")
	return
}

func otherTestFunction() error {
	var err error
	defer func() {
		if err != nil {
			err = fmt.Errorf("%w: touched", err)
		}
	}()

	err = fmt.Errorf("failed 2")
	return err
}

func printDeferStackOrder() {
	defer fmt.Println("third")
	defer fmt.Println("second")
	defer fmt.Println("first")
}

func EmitError(err error) {
	fmt.Println(err)
}
