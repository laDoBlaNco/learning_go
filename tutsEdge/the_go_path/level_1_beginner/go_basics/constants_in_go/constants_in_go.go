package main

func main() {
	const welcome string = "hello world" // as a constant welcome will never be changed
	const pi  = 3.14              // leaving this untyped will allow go to infer the type
	//change the type later if we need to use this in an expression where another usable type is
	// needed.

	var size float64 = 1.0

	println(welcome)
	println(pi+size)
}
