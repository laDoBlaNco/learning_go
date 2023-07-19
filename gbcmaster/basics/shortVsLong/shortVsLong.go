package shortVsLong


// version := 0 // this shouldn't be done
var version string // for package scoped variables you only use declarations

func main() {
	// score := 0 //Don't do this
	var score int // This is a declaration and alreay set to 0 by default

	// If you know the  value then use short, its the most used way for gophers
	// var width,height = 100,50 // don't do this
	width, height := 100, 50 // do this.

	//width = 50 // assigns 50 to width
	//color := "red" // new variable: color -- Dont do this
	width, color := 50, "red" // do this

	// When you want to group variables together for greater readability
	var (
		// related:
		video string

		// closely related:
		duration int
		current  int
	)
}
