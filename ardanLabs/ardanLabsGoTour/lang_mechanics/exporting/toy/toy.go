// Package toy contains support for managing toy inventory
package toy

// Declare a struct type named Toy with four fields
// Name string
// Weight int
// onHand int
// sold int
type Toy struct{
	Name string
	Weight int
	
	onHand int
	sold int
}

// Declare a func named New that accepts values for the exported fields. Return a pointer
// of type Toy that is initialized with the parameters
// NOTE I made an error here. the challenge said explicitly "for the exported fields" and I did
// args for all 4. That's also why in the func main we used the Update funcs for "initialize"
// the data as Bill said. I made the corrections below.
func New(name string,weight int)*Toy{
	return &Toy{
		Name:name,
		Weight:weight,
		// onHand:onhand,
		// sold:sold,
	}
}

// Declare a method named OnHand with a pointer receiver that returns the current on
// hand count
func(t *Toy)OnHand()int{
	return t.onHand 
}

// Declare a method named UpdateOnHand with a pointer receiver that updates and returns 
// the current on hand count
func(t *Toy)UpdateOnHand(c int)int{
	t.onHand = c
	return t.onHand
}

// Declare a method named Sold with a pointer recvr that returns the current sold count
func(t *Toy)Sold()int{
	return t.sold
}

// Declare a method named UpdateSold with a pointer recvr that updates and returns the 
// current sold count
func(t *Toy)UpdateSold(c int)int{
	t.sold = c
	return t.sold 
}


