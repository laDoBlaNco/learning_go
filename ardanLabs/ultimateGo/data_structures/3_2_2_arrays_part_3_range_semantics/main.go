package main

import "fmt"

// before we talked about for/range has 2 forms (pointer and value semantics). Understanding this
// on a deeper level will help us to really understand the micro semantics about the decisions
// that we make and the costs we are managing.
// 	- Value semantics being cleaner but with the cost of inefficiencies
// 	- Pointer semantics being more efficient but with the cost of being more error prone since we are sharing
//    data now.

func main() {

	// using the pointer semantic form of the for/range.
	friends := [5]string{"annie", "betty", "charlie", "doug", "edward"}
	fmt.Printf("Bfr[%s] : ", friends[1]) // we print out the value at friends[1] "betty"

	for i := range friends { // pointer semantic form of the for/range
		friends[1] = "jack" // we change the value of Betty to Jack using pointer semantics

		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1]) // then we print out the new value of friends[1] "jack"
		}
	}

	// Using the value semantic form of the for/range.
	friends = [5]string{"annie", "betty", "charlie", "doug", "edward"} // we have the same array
	fmt.Printf("Bfr[%s] : ", friends[1])                               // here we print out the same value of friends[1] "betty"

	for i, v := range friends { // value semantic form of the for/range
		friends[1] = "jack" // we change betty our for Jack as we did before

		if i == 1 {
			fmt.Printf("v[%s]\n", v) // But when we go to print the new value at friends[1] it still say "betty"
			// Remember that value semantics works with its own copy. So this for/range using value semantics
			// is iterating over ITS OWN COPY of the friends array, not the original. so we do change out betty
			// to jack, but because its a copy, it gets thrown away after the change is made.
			// THESE ARE THE DATA SEMANTICS THAT ARE SO IMPORTANT TO UNDERSTAND as we will truly understand
			// the costs of what we are doing and the decisions we are making.
		}
	}

}

/*
So again as a guideline, if the data we are working on is a numeric, string, or
bool, we use value semantics to move the data around. This includes declaring
fields on a struct type. The nice thing about using value semantics is that we
are guaranteed that each function is operatoring on its own copy. This means that reads and writes
to the data are isolated to that function. This helps with integrity and identifying bugs related
to data corruption. 
*/
