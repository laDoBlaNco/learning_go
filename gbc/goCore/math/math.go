package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var nl = fmt.Println

func main() {
	fmt.Println("5 + 4 =", 5+4)
	fmt.Println("5 - 4 =", 5-4)
	fmt.Println("5 * 4 =", 5*4)
	fmt.Println("5 / 4 =", 5/4)
	fmt.Println("5 % 4 =", 5%4)
	inc := 1
	inc = inc + 1
	inc += 1
	inc++
	fmt.Println(inc)
	nl()
	fmt.Println("Float Precision increases with the size of values:")
	fmt.Println("Float Precision =", 0.1111111+0.1111)
	fmt.Println("Float Precision =", 0.1111111111111111111111111111+0.111111111111111111111111111111111)
	nl()
	fmt.Println("Random values:")

	seedVal := time.Now().Unix()
	rand.Seed(seedVal)
	randNum := rand.Intn(50) + 1
	fmt.Println("Random:", randNum)
	nl()
	fmt.Println("Builtin math funcs:")
	fmt.Println("Abs(-10) =", math.Abs(-10))
	fmt.Println("Pow(4,2) =", math.Pow(4, 2))
	fmt.Println("Sqrt(16) =", math.Sqrt(16))
	fmt.Println("Cbrt(8) =", math.Cbrt(8))
	fmt.Println("Ceil(4.4) =", math.Ceil(4.4))
	fmt.Println("Floor(4.4) =", math.Floor(4.4))
	fmt.Println("Round(4.4) =", math.Round(4.4))
	fmt.Println("Log2(8) =", math.Log2(8))
	fmt.Println("Log10(100) =", math.Log10(100))
	fmt.Println("Log(7.389) =", math.Log(7.389))
	fmt.Println("Max(5,4) =", math.Max(5, 4))
	fmt.Println("Min(5,4) =", math.Min(5, 4))
	// There are also functions for Cos,Tan,Acos,Asin,Atan,Asinh,Acosh,Atanh Atan2
	// Cosh,Sinh,Sincos,Htpot - trigonometry

	r90 := 90 * math.Pi / 180    // convert degrees into radians
	d90 := r90 * (180 / math.Pi) // convert radians into degrees
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	fmt.Println("Sin(90) =", math.Sin(r90))

	nl()
	fmt.Println("Exercise #2")
	//Enter Number 1: 5 (trim whitespace and convert)
	//Enter Number 2: 4 (same as above)
	// 5 + 4 = 9
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number 1: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	n1, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter Number 2: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	n2, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s + %s = %d\n", n1, n2, n1+n2)

}
