package main

import (
	"fmt"
)

var pl = fmt.Println

func sayHello(){
	pl("Hello gophers") 
}

func getSum(x,y int)int{
	return x+y
}

func getQuotient(x,y float64)(ans float64,err error){
	if y==0{
		return 0,fmt.Errorf("You can't divide by zero...")
	}
	return x/y,nil
}

func getTwo(x int)(int,int){
	return x+1,x+2
}

// Veradic functions are functions that receive an unknown amount of params
func getSum2(nums ...int)int{
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	return sum
}

func getSum3(nums []int)int{
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	return sum
}

func changeValue(f3 int)int{
	f3++
	return f3
}


func main() {
	// func funcName(parameters parameter Type) returntype {body} 
	sayHello() 
	pl(getSum(5,4)) 
	ans,err:=getQuotient(5,4)
	if err!=nil{
		pl(err)
	}
	pl(ans)
	
	f1,f2:=getTwo(5) 
	pl(f1,f2) 
	
	pl("Unknown sum:",getSum2(1,2,3,4,5)) 
	pl("Unknown sum with unexpected Slice Param:",getSum2([]int{1,2,3,4,5}...))
	pl("Unknown sum with expected slice",getSum3([]int{1,2,3,4,5})) 
	
	f3:=5
	pl("f3 before func:",f3) 
	changeValue(f3) 
	pl("f3 after func:",f3)
	
}
