// in this example we will see what was called out in our first example. Implementing a web
// request with a context that is used to timeout the request if it takes too long
package main

import(
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	
	// let's create our new request
	req,err:=http.NewRequest("GET","https://www.ardanlabs.com/blog/post/index.xml",nil) 
	if err!=nil{
		log.Println("ERROR:",err) 
		return
	}
	
	// let's create a timeout of 50 milliseconds
	ctx,cancel:=context.WithTimeout(req.Context(),50*time.Millisecond) 
	defer cancel() 
	
	// let's bind the new context into the request now
	req = req.WithContext(ctx) 
	
	// make the web call and return any error. Do will handle the context  level timeout.
	resp,err := http.DefaultClient.Do(req)
	if err!=nil{
		log.Println("ERROR:",err)
		return
	}
	
	// close the response body on the return
	defer resp.Body.Close() 
	
	// write the response to the stdout
	io.Copy(os.Stdout,resp.Body) 
	
	
}
