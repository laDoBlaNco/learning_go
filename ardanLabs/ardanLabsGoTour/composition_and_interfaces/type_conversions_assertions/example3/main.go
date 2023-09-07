// in this program we see how method sets can also affect behavior
package main

import "fmt" 

type user struct{
	name string
	email string
}

// implement the Stringer interface
func(u *user)String()string{
	return fmt.Sprintf("My name is %q and my email is %q",u.name,u.email) 
}

func main(){
	
	u := user{
		name:"Kevin Whiteside",
		email:"whitesidekevin@gmail.com",
	}
	
	fmt.Println(u)
	fmt.Println(&u) 
}
// TODO: Not sure why this works this way yet. Let me take a look in the morning: So I think
// what's happening here goes back to the the way we implemented the interface. Since I used
// *user, printing u doesn't find the method or Go says it doesn't have the Stringer implemented
// and prints  the generic way for a  struct. But when I use the &u (pointer) it now sees the 
// Stringer interface and uses our version. So this falls right in line with:
// 		- value receiver 			value/pointer
// 		- pointer receiver			pointer only 
// I can test this just by changing the pointer receiver to a value receiver, and I should be able
// to see my interface on both calls for u and &u *** And it worked ***

