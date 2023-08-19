package main

import(
	"fmt"
)


/*
	METHODS
	A function is called a method when that function has a receiver declared. But as many
	say its still function in the end. The receiver is the parameter that is declared
	between the keyword func and the function name. 
	
	Let's start with the Method Declarations:
	There are two types of receivers, value receivers for implementing value semantics and 
	pointer receivers for implementing pointer semantics.

*/

// First we create our type. Methods are always part of types. So they are also called 
// associative functions.
type user struct{
	name string
	email string
}

// From our user type we will create our first method 'notify' which will be implemented with
// a value receiver. This means the method operates under value semantics and will operate
// on its own copy of the value used to make the call
func (u user) notify(){ // again we see the the 'receiver' is the same as 'self or this' in other
// languages, more or less. Its what connects or associates our 'user' type to our method.
// and gives us access to everything inside our instance of user.
	fmt.Printf("Sending User Email to %s<%s>\n",u.name,u.email) 
}

// The next method 'changeEmail' is implemented as a pointer receiver. This means the method 
// operates under our pointer semantics and will operate on shared access to the value used
// to make the call. 
func (u *user)changeEmail(email string){
	u.email = email
	fmt.Printf("Changed User Email To %s\n",email)
}

// Outside of a few exceptions, a method set for a type should never contain a mix of value
// and pointer receivers. Data semantic consistency is critically important and this 
// includes declaring methods or associative functions. 

// So how do we now call these methods. 
func main(){
	
	// when making a method call, the compiler doesn't care if the value used to make 
	// the call matches the receiver's data semantics exactly. The compiler just wants
	// a value or pointer of the same type
	bill := user{"Bill","bill@email.com"} 
	bill.notify()
	bill.changeEmail("bill@hotmail.com") 
	fmt.Println(bill) 
	fmt.Println() 
	
	// We can see that a value of type user is constructed and assigned to the bill
	// variable. In the case of the notify call, the bill variable matches the receiver
	// receiver type which is using a value receiver. Int eh case of changeEmail call, 
	// the bill variable doesn't match the receiver type which is using a pointer receiver
	// However, the compiler accepts the method call and shares the bill variable with the 
	// method. Go will adjust to make the call. This is the case where we didn't have to do
	// the actual syntax of (&bill).changeEmail(...) 
	
	// This works the same when the variable used to maek the call is a pointer var
	bill2 := &user{"Bill","bill@email.com"} 
	bill2.notify() 
	bill2.changeEmail("bill@hotmail.com") 
	fmt.Println(bill2)
	fmt.Println() 
	
	
	// This time we used a pointer variable to a value of type user. Once again, Go adjusts
	// to make the method call when calling the notify method. If Go didn't adjust, again
	// we would have had to to the following to make these work.
	bill3 := user{"Bill","bill@email.com"}  
	(&bill3).changeEmail("bill@hotmail.com") 
	fmt.Println(bill3)
	fmt.Println()
	// and 
	bill4 := &user{"Bill","bill@email.com"}
	(*bill4).notify() 
	fmt.Println(bill3)
	fmt.Println() 
	// do the conversion ourselves explicitly. We can be glad we don't have to do that.
	
	/*
		DATA SEMANTIC GUIDELINES FOR INTERNAL TYPES
		As a guideline, if the data we are working is an internal type (slice,map,channel,
		function, or interface) then we should be using value semantics to move the data
		around our program boundaries. This includes fields on a  type. However, when we 
		are reading or writing  to something we need to remember that we are using pointer
		semantics. For example, the following is in the net pkg. They are declared with an
		underlying type which is a slice of bytes. Because of this, these types follow the
		guidelines we mention above
		
		type IP []byte
		type IPMask []byte
		
		func (ip IP) Mask(mask IPMask) IP{ // note the value veciever
			if len(mask) == IPv6len && len(ip) == IPv6len && allFF(mask[:12]) {
				mask = mask[12:] 
			}
			if len(mask) == IPv4len && len(ip) == IPv6len && bytesEqual(ip[:12], v4InV6Prefix){
				ip = ip[12:]	
			}
			n := len(ip)
			if n != len(mask){
				return nil
			}
			out := make(IP, n)
			for i:=0;i<n;i++{
				out[i] = ip[i] & mask[i] 
			}
			return out
		}
		
		Note that with the Mask method, value semantics  are in play for both the receiver,
		parameter, and return arguments. This method accepts its own copy of a Mask value,
		it mutates that value and then returns a copy of the mutation. This method is using 
		value semantic mutation. This is not an accident or random.
		
		A function c an decide what data input and output it needs. What it can't decide is the
		data semantics for how the data flows in and out. The data drives that decision and the
		function must comply. This is why Mask implements a value semantic mutation api. It 
		must respect how a slice is designed to be moved around our program.
		
		func ipEmptyString(ip IP) string{
			if len(ip) == 0{
				return ""
			}
			return ip.String()
		}
		
		This one is also using value semantics for the input and output. This function accepts
		its own copy of an IP value and returns a string value. No use of pointer semantics
		because the data dictates the data semantics and not the function. So we can't just do
		things willy nilly as the designer if we want mechanical sympathy and integrity. 
		
		NOTE:  One exception to using value semantics is when you need to share a slice or map
		with a function that performs UNMARSHALING or DECODING. 
	*/
	 
	
	
}




