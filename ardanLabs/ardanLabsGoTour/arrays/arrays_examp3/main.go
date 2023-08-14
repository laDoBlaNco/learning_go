// CONTIGUOUS MEMORY CONSTRUCTION
/*
This code helps to prove that an array provides a contiguous layout of memory

Here we declare an array of 5 strings initialized with values, then use VALUE semantic iteration to display
the information about each string. The output as we can see shows each individual string value, the address
of the v variable and the address of each element in the array.

We can also see how the array is a contiguous block of memory and how a string is a two word or 16 byte data structure on a 64 bit arch and an 8 byte data structure on a 32 bit arch. The address for each element is 
distanced by that respective stride. 

The fact that the v variable has the same address on each iteration strengthens the understanding as well
that v is a local variable of type string which contains a copy of each string value during iteration
*/

package main

import "fmt"

func main(){
	
	five := [5]string{"Annie","Betty","Charley","Doug","Bill"}
	
	for i,v := range five{
		fmt.Printf("Value[%s]\tAddress[%p]  IndexAddr[%p]\n",v,&v,&five[i])
	}
		
}

/*
NOTES:
	- If you don't understand the data, you don't understand the problem
	- if you don't understand the cost of solving the problem, you can't reason about the problem
	- If you don't understand the hardware, you can't reason about the cost of solving the problem
	- Arrays are fixed length data structures that can't be changed
	- Arrays of different sizes are considered to be different types
	- Memory is allocated as a contiguous block
	- Go gives you control over spacial lcoality.

*/
