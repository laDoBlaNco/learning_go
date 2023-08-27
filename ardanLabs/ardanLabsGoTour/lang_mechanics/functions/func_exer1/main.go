package main


/*
	Part A: Declare a struct type to maintain information about a user. Declare a function
	that creates value of and returns pointers of this type and an error value. Call this
	function from main and display the value.
	
	Part B: Make a second call to your function but this time ignore the value and just test
	the error value. 

*/

// add imports
import(
	"fmt"
)

// declare a type named user
type user struct{
	name string
	age int
}

// declare a function that creates user type values and returns a pointer to that value
// and an error value of nil
func newUser(n string,a int)(*user,error){
	
	if a > 80{
		return nil, fmt.Errorf("You are too old! Your're %v\n",a) 
	}
	u:=user{
		name:n,
		age: a,
	}
	return &u,nil
}


func main(){
	
	u,err:=newUser("Kevin",46) 
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("User:",u) 
	
	fmt.Println("===========================================")
	
	_,err=newUser("Kelen",101) 
	if err!=nil{
		fmt.Println(err) 
	}
	
}

