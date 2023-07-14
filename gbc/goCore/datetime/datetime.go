package main

import(
	"fmt"
	"time"
	"log" 
)
var pl = fmt.Println

func main(){
	
	// current time
	now:=time.Now() 
	pl(now.Year(),now.Month(),now.Day())
	pl(now.Hour(),now.Minute(),now.Second())

	pl()
	// we can change our location as well.
	loc,err:=time.LoadLocation("America/New_York")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("Time in New York %s\n",now.In(loc)) 
	loc,_ =time.LoadLocation("Asia/Shanghai") 
	fmt.Printf("Time in Shanghai %s\n",now.In(loc))
	
	pl()
	locEST,_ :=time.LoadLocation("EST")
	locUTC,_ :=time.LoadLocation("UTC")
	locMST,_ :=time.LoadLocation("MST")
	fmt.Printf("EST: %s\n",now.In(locEST))
	fmt.Printf("UTC: %s\n",now.In(locUTC))
	fmt.Printf("MST: %s\n",now.In(locMST))
	pl()
	
	bDay:=time.Date(1976,time.December,
		25,12,00,00,0,time.Local)
	diff:=now.Sub(bDay) 
	fmt.Printf("Days Alive: %d days\n",int(diff.Hours()/24)) 
	fmt.Printf("Hours Alive: %d days\n",int(diff.Hours())) 
	
	  

	
}
