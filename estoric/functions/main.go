package main

import (
	"errors"
	"fmt"
)

// The only time you don't need a main func is if you are creating a package
// to be imported only

func main() {
	firstNum := 0
	secondNum := 4

	// mySum := calculateSum(firstNum, secondNum)

	fmt.Println(calculateSum(firstNum, secondNum))
	fmt.Println(calculateSumAndDiff(firstNum, secondNum))

	// Veriadic functions
	fmt.Println(calculateSums(1, 2, 3, 4, 5))
	fmt.Println(calculateSums([]int{1, 2, 3, 4, 5}...)) // to use a slice we need ...
	//fmt.Println([]string{"my", "name", "is", "ladoblanco"}...) //?????
	//Not sure why this gives an error if Println is variadic

	// Error handling
	remainder, err := calculateRemainder(secondNum, firstNum)
	fmt.Println(remainder, err)
	remainder, err = calculateRemainder(firstNum, secondNum)
	fmt.Println(remainder, err)

}

func calculateSum(num1, num2 int) int {
	return num1 + num2
}

func calculateSumAndDiff(num1, num2 int) (int, int) {
	return num1 + num2, num2 - num1

}
func calculateSumAndDiff2(num1, num2 int) (mySum, myDiff int) {
	mySum = num1 + num2
	myDiff = num2 - num1
	return // when using declared return values we can use "naked" return
	// we declare the return values in the open when we say how many to return

}

// Variadic funcion
func calculateSums(nums ...int) int { //...int makes nums a slice of params
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Error handling with functions
func calculateRemainder(numerator, denominator int) (int, error) { // return error
	if denominator == 0 { // error is defined in errors pack
		err := errors.New("cannot divide by zero") // need to use errors pack
		return 0, err                              // if err than return 0 or -1 with err
	}
	return numerator % denominator, nil // if all good return result with nil
}

/*
Go handles functions just as python and JS, as first class citizens. But also
Go has its own twist on some functions with some experimental items as well.
Functions give us the ability to write clean modular code that can be reused.
Kinda like the parts of a PC. As long as it has the right connections you can
pull one out and put another one in.

As an example from Go itself:
func Println(a,...interface{})(n int,err error){
	return Fprintln(os.Stdout,a...)
}
func Fprintln(w io.Writer,a ...interface{})(n int,err error){
	p := newPrinter()
	p.doPrintln(a)
	n,err = w.Write(p.buf)
	p.free()
	return
}
When we use fmt.Println we are actually using Fprint with addition params. So
rather than writing the full func def out for every implementation, Go uses its
own function modularity.

There are four key parts:
1. 'func'
2. function name
3. parameters with type - always in ()s
4. return type (if it needs to return a value). We can also return more than
one value - When returning more than one value, you typically don't want to go
over returning 2 but we can if needed.

Something else to keep in mind is that we can only put the amount of params that
we've defined. But functions like Println can take as many as you give it? This
is do to it being a veradic function, meaning that it will take all params that
you throw at it. To get this same functionality we use the ... to tell the
function that it may get multiple params in the end. Its the same ... we put on
a slice to break it up for a veriadic function. To do this you include the ...
in front of the type as in ...int

Another important part of functions, now that we are looking at returning
multiple values is Go's habit of having functions return errors to take care
of error handling. (Example above). The pattern of returning 0 or -1 with err
or the result with nil is because typical where you are using the function you
will build in some sort of conditional to handle an err if that's what the
function returns. As seen above we can declare and expect 2 results in vars
 If err is nil all good, if its not, then we do something else
since the result will be 0 or -1 and not what we expected.

*/
