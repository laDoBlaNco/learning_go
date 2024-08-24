package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
	"time"
)

// ... for quick debugging
var p = fmt.Println

// switch statements express conditionals across many branches
func main() {

	// Here's a basic switch:
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		p("one")
	case 2:
		p("two")
	case 3:
		p("three")
	}
	
	// you can use commas to separate multiple expressions in the same 
	// case statement. We use the optional default case in this example
	// as well
	switch time.Now().Weekday(){
		case time.Saturday,time.Sunday:
			p("Its the weekend") 
		default:
			p("It's the weekday") 
	}
	
	// switch without an expression is an alternative way to express if/else
	// logic. here we also show how the case expressions can be non-constants
	t:=time.Now() 
	switch{
		case t.Hour()<12:p("It's before noon") 
		default:p("It's after noon") 
	}
	
	// A 'type' switch compares types instead of values. You can use this to
	// discover the type of an interface value. In this example, the var t 
	// will have the type corresponding to its clause.
	whatAmI:=func(i interface{}){
		switch t:=i.(type){ // note the obligatory syntax .(type) 
			case bool:p("I'm a bool")
			case int:p("I'm an int")
			default:fmt.Printf("Don't know type %T\n",t)  
		}
	}
	 whatAmI(true) 
	 whatAmI(69)
	 whatAmI("hey") 

}
