package main

import "fmt"

// COPYING SLICES MANUALLY
// There is a built-in function as we already know named 'copy' which will allow for the shallow
// copying of slices. since a string has a backing array as well, but of immutable bytes, it
// can be used as a source but never a destination.

func main(){
	
	slice1 := []string{"a","b","c","d","e"}
	slice2 := make([]string,len(slice1)) 
	
	copy(slice2,slice1) 
	
	inspectSlice(slice1)
	inspectSlice(slice2)
	fmt.Println() 
	// as long as the dest slice has the right type and length, the built-in 'copy' can perform
	// its shallow copy.
	
	
	// POINTER SEMANTIC MUTATION IN SLICES
	// Its very important to remember that even though we are using vALUE SEMANTICS to move our
	// slice around the program, when reading and writing a slice, we are using POINTER SEMANTICS.
	// Sharing individual elements of a slice with different parts of our program can cause
	// unwanted side effects. 
	
	// Here we construct a slice of 1 user, set a pointer to that user, then use the
	// pointer to update likes. 	
	users := make([]user,1) 
	ptrUsr0 := &users[0] 
	ptrUsr0.likes++ 
	
	for i:=range users{
		fmt.Printf("User: %d Likes: %d\n", i,users[i].likes) 
	}
	fmt.Println()
	 	
	// See that a slice is used to maintain a collection of users. When a pointer is set to the
	// first user and used to update likes, The ouput shows that the user is working.
	// But when we want to add an new user to the collection and use the pointer again to update
	// note what happens
	users = append(users,user{}) 
	ptrUsr0.likes++  
	for i:=range users{
		fmt.Printf("User: %d Likes: %d\n", i,users[i].likes) 
	}
	// But since the append function replaced the backing array with a new one, the pointer is
	// updating the old backing array and the likes are lost. As well with a reference to the
	// old backing array still, we have created a memory leak as that backing array won't
	// be wiped. 
	
	// We must be very careful to know if a slice is going to be used in an append operation
	// during the couse of the running program. How we shared the slice needs to be considered
	// Sharing individual indexes might not be the best idea. Sharing an entire slice value
	// might not work either when appending is in operation. Probably using a slice as a field
	// in a struct, and sharing the struct value would be the better way to go.
}

type user struct{
	likes int
}


func inspectSlice(slice []string){
	fmt.Printf("Length[%d]  Capacity[%d]\n",len(slice),cap(slice))
	for i := range slice{ // this is the pointer semantic version
		fmt.Printf("[%d]  %p  %s\n",i, &slice[i],slice[i]) // note the 16 byte stride
	}
}


