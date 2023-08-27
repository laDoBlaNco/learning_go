package main

import "fmt"

// Now let's look at real embedding. Its the same programming, but notice that now
// instead of creating a field of type 'user', we simply EMBED our user type directly
// into our 'admin' type

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user // user EMBEDDED in our admin struct directly without the use of a named field
	// here we are using the VALUE SEMANTICS for our embedding. we can also use pointer semantics
	// as seen below
	// and embed a pointer of the type '*user'. In either case, accessing the embedded value is done
	// through the user of the type's name as the field.

	level string
}
type admin2 struct {
	*user // pointer semantics for embedding, meaning we need to use an address with this later
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "Kevin Whiteside",
			email: "whitesidekevin@gmail.com",
		},
		level: "super",
	}

	// Now we can access the INNER TYPE's method with our direct way as before
	ad.user.notify()

	// or using it as a PROMOTED method or behavior on our 'admin' type
	ad.notify()

	ad2 := admin2{
		user: &user{ // notice the only change is that we need to use an address of our user	
			name:  "Odalis Whiteside",
			email: "odalislorenzo74@gmail.com",
		},
		level: "super",
	}

	ad2.user.notify() // the calling of the methods/fields are the same no matter the data semantics
	ad2.notify()
}

// The best way to thing about embedding is to view the user type here as an INNER TYPE and
// admin as the OUTER TYPE. Its this inner/outer type relationship that is magical because
// with embedding, everything related to the inner type (BOTH FIELDS AND METHODS) can be
// promoted up to the outer type.

// NOTE that although this looks like inheritance, IT IS NOT. This is about PROMOTING BEHAVIOR
// NOT REUSING STATE
// We can also use embedding with interfaces. Let's look at that next.
