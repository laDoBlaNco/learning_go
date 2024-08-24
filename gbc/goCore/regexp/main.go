package main

import(
	"fmt"
	"regexp" 
)
// ...for quick debugging
var pl = fmt.Println

// mainly regexp is for pattern matching with strings.

func main(){

	str := "The ape was at the apex"
	match,_ := regexp.MatchString("(ape[^ ]?)",str)
	pl(match)  
	
	str2 := "Cat rat mat fat pat"
	r,_ := regexp.Compile("([crmfp]at)")
	pl("MatchString:",r.MatchString(str2))   
	pl("FindString:",r.FindString(str2))   
	pl("Index:",r.FindStringIndex(str2))   
	pl("All String:",r.FindAllStringIndex(str2,-1))   
	pl("First 2 String:",r.FindAllStringIndex(str2,2))   
	pl("All submatch index:",r.FindAllStringSubmatchIndex(str2,-1))   
	pl(r.ReplaceAllString(str2,"Dog")) 
	

}
