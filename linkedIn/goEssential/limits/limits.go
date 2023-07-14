package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// with the fallecies of networking and distributed computing, we need to protect
// ourselves. We can do that with the io, context, and http packages.
func main() {
	// first start by creating a context that we want to work in. Remember a context
	// is like creating a constaint that we want our connection work within.
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
	if err != nil {
		log.Fatal(err) // not log.Fatalf. The diff is Fatalf uses a string, Fatal just the err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	const mb = 10 //1 << 20 // this 1 left bit shift 20 creates  1048576 or 1049 or 1mb
	r := io.LimitReader(resp.Body, mb)
	io.Copy(os.Stdout, r)
	fmt.Println() 

	//CHALLENGE
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	fmt.Printf("%#v\n", user)
}

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"` // from public_repos
}


//CHALLENGE:
// userInfo return information on github  user
func userInfo(login string) (*User, error) {
	/*TODO:
	Call teh github API for a given login
	e.g. https://api.github.com/users/tebeka
	e.g. https://api.github.com/users/

	*/
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close() 
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%#v - %s", url, resp.Status)
	}
	usr:=User{Login:login}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&usr); err != nil {
		log.Fatalf("error: can't decode - %s\n", err)
	}

	return &usr, nil
}

