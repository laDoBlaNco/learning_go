package main

import(
	"encoding/json"
	"errors"
	"fmt"
)

// As we saw in the last example, Go functions can return multiple values and most functions
// in Go return error values. 

// We always check the error value as part of the programming logic, thus our if err!= nil {}

// As we'll see in this example, we can use the _ blank identifier if we don't care about one
// or  more of the results from our functions.


// user is a struct type that declares user information
type user struct{
	ID int
	Name string
}

// updateStates provides update stats
type updateStats struct{
	Modified int
	Duration float64
	Success bool
	Message string
}

func main(){
	
	// let's declare and initialize a value of type user
	u := user{
		ID:1432,
		Name:"Betty",
	}
	
	// Now let's update the user name. Don't care about the update stats
	if _,err := updateUser(&u);err!=nil{
		fmt.Println(err)
		return
	}
	
	// Display the udpate was successful
	fmt.Println("Updated user record for ID",u.ID) 
}

// updateUser updates the specified user document
func updateUser(u *user)(*updateStats,error){
	
	// response simulates a JSON response
	response := `{"Modified":1, "Duration":0.005, "Success":true, "Message":"updated"}`
	
	// unmarshal the json document into a value of the userStats struct type
	var us updateStats
	if err:=json.Unmarshal([]byte(response),&us);err!=nil{
		return nil,err
	}
	
	// Check the update status to verify the update was successful
	if us.Success !=true{
		return nil,errors.New(us.Message) 
	}
	
	return &us,nil
}

// we keep returning pointers and throwing stuff on the heap. I'm assuming we are doing this
// because we are simulating server responses maybe? 

// now in the next example, let's look at redeclaration or 'shadowing'
