package main

import(
	"fmt"
	"math/rand"
	"time"
)

/*
	Assuming the program does declare two types named Car and Cloud that each implement
	the fmt.Stringer interface, we can construt a collection that allows us to store both
	values. Then 10 times, we randomly choose a number from 0 to 1, and perform a type assertion
	to see if the value at the random index contains a Cloud value. Since it's possible its 
	not a Cloud value, the second form of the type assertion is critical here (comma, ok) 
	
	This makes me think that where I would use an option in other langs, I would do this idiom
	in Go

*/

type car struct{}

func(car)String()string{
	return "Vroom!" 
}

type cloud struct{}

func(cloud)String()string{
	return "Big Data!" 
}

func main(){
	
	rand.Seed(time.Now().UnixNano()) 
	
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}
	
	for i:=0;i<10;i++{
		rn := rand.Intn(2) 
		
		if v,is := mvs[rn].(cloud);is{ // note how we didn't actually use 'ok' here, meaning we can
		// actually use whatever var makes sense
		fmt.Println("Got Lucky:",v) 
		continue 			
		}
		
		fmt.Println("Got Unlucky") 
	}
}
