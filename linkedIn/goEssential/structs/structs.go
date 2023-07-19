package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"
)

// type struct
type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func main() {
	b := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b.TimeLeft())

	b.Update(10.5)
	fmt.Println(b.Balance)

	expires := time.Now().Add(7 * 24 * time.Hour)

	b1, err := NewBudget("puppies", 32.2, expires)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", b1)
	}

	b2, err := NewBudget("kittens", -32.2, expires)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", b2)
	}
	fmt.Println("________________________________________________")
	//CHALLENGE OUTPUT
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}
	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())

	fmt.Println("________________________________________________")
	// Interfaces:
	sq := Square2{20}
	fmt.Println(sq.Area())

	c := Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{sq, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)

	fmt.Println("________________________________________________")
	// CHALLENGE - WRITER INTERFACE
	c2 := &Capper{os.Stdout} // here we put osStdout which is a writer into our
	// type Capper as the osWriter field.
	fmt.Fprintln(c2, "Hello there") // using Fprintln since we are using io.Writer?

	fmt.Println("________________________________________________")
	// Generics example:
	fmt.Println(min([]float64{2, 1, 3}))
	fmt.Println(min([]string{"b", "a", "c"}))

}

// methods to our type, first is not changing the underlying struct so its value receiver
// the second is mutating our struct, so its a pointer receiver. But its not returning
// anything.
func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}
func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

// For Go's OOP we use new funcs - note we are returning a pointer to Budget
func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("empty campaignID")
	}
	if balance <= 0 {
		return nil, fmt.Errorf("balance must be bigger than 0")
	}
	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}
	return &b, nil //escape analysis will happen when compiled  to put the pointer in heap
}

// CHALLENGE - define a square struct with 2 methods
// Square is a square
type Square struct {
	X      int
	Y      int
	Length int
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}
	s := Square{
		X:      x,
		Y:      y,
		Length: length,
	}
	return &s, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

// Area returns the square area - doesn't change anything in the square so it can
// use a value receiver
func (s Square) Area() int {
	return s.Length * s.Length
}

// Another square
type Square2 struct {
	Length float64
}

// Another method Area for Square2
func (s Square2) Area() float64 {
	return s.Length * s.Length
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Area returns the circle of the square
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sumAreas return the sum of all areas in the slice
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// Shape is aa shape interface -- not a struct but an interface (job description)
type Shape interface {
	Area() float64
}

// CHALLENGE -- CAPPER  io.Writer
// Capper implements io.Writer and turns everything to uppercase
type Capper struct { // struct with only 1 field, an underlying io.Writer
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A') // getting diff between lower and capital 'a'

	out := make([]byte, len(p)) // create new slice so we don't change p
	for i, c := range p {
		if c >= 'a' && c <= 'z' { // if the char is lowercase (between 'a' & 'z')
			c -= diff // subract the diff between lower and cap to make it cap
		}
		out[i] = c
	}
	return c.wtr.Write(out) //calling underlying io.Writer's method
}

// in 1.18 generics is the new black. What can we do?
type Ordered interface { //in generics an interface can also be a SET of types
	int | float64 | string // not just a set of methods
}

func min[T Ordered](values []T) (T, error) { // the big difference are type constraints
	// and rather than calling the type to receive or return, we use T for general type
	if len(values) == 0 {
		var zero T                                    // zero of some type, generic zero default as we don't know the type
		return zero, fmt.Errorf("min of empty slice") // return zero,error
	}

	m := values[0]
	for _, v := range values[1:] { //Go idiom, set the first index and run through 2nd and  on
		if v < m {
			m = v
		}
	}
	return m, nil
}
