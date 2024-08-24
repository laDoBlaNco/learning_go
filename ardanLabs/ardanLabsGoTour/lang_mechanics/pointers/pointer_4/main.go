// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Before we spoke about how the compiler must know the the size of a value at compile
// time or it can't construct it on the stack but it'll need to be constructed on 
// the HEAP.
// The algo the compiler uses to determine if a value should be constructed on the
// stack or on the heap is called "escape analysis". The name the algorithm makes 
// it sound like values are constructed on the stack first and then escape (or move)
// to the heap when necessary. This is NOT the case. The construction of a value only
// happens once, and the escape analysis algorithm decides where that will happen
// on the stack or on the heap. Only construction on the heap is referred to an
// ALLOCATION in Go. 

// Our understanding of escape analysis is about understanding value ownership. The
// idea is, when a value is constructed within the scope of a function, then that 
// function owns the value. From there, we ask the question, does the value being 
// constructed still have to exist when the owning function returns? If the answer
// is no, the value can be constructed on the stack. If the answer is yes, the value 
// must be constructed on the heap.

// NOTE: The ownership rule is a good base rule for identifying code that causes
// allocations, and that will help us to understand even better the costs of the
// design decisions. However we must appreciate that escape analysis has flaws that can
// result in non-obvious allocations. Also, the algorithm takes opportunities to
// leverage compiler optimizations to save on allocations. 


// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	u1 := stayOnStack()
	u2 := escapeToHeap()

	println("u1", &u1, "u2", u2)
}

// stayOnStack is using value semantics to return a user value back to the caller.
// In other words the caller gets THEIR OWN COPY of the user value being constructed
// and that's why this one stays on the stack, because we aren't returning a pointer
// to anything. Once it returns the user value it constructs no longer needs to exist
// since the caller is getting their own copy of it. Therefore, the construction of
// the user v alue inside stayOnStack can happen on the stack. No allocation.
//go:noinline
func stayOnStack() user { // NOTE the 'user' being returned and not *user
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

	return u
}

// escapeToHeap() is using pointer semantics to return a user value back to the caller
// In this case, the caller gets shared access (an address) to the user value being
// constructed. So when the escapeToHeap func is called and returns, the user value that
// IT constructs does in fact still need to exist, since the caller is getting shared
// access to the value. Therefore, the construction of teh user  value inside of 
// escapeToHeap can't happen on the stack, but must happen on the heap. Yes allocation.

// Think about if it would be constructed on the stack. When execution returns to the 
// main frame with a *User and the value it points to is DOWN the stack, This is what we 
// call sharing up the stack. The memory that
// is holding that value will be reused on the very next function call. oops. Integrity
// would be lost. Once the control goes back to the calling function, the memory on the
// stack (self-cleaning) is reusable again. A new frame will be sliced and the memory
// will be overridden, destroying the shared value. 

// This again is why we say that the stack is self-cleaning. Zero value initialization
// helps every stack frame that we need, to be cleaned without the need of a GC.
// The stack is self cleaning since a frame is taken and initialized for the execution
// of each function call. The stack is cleaned during function calls and not on returns
// as we mentioned before because the compiler doesn't know if that memory on the stack
// will ever be needed again, so why waste the electrons.

// Escape analysis decides if a value is constructed on the stack (default) or the heap
// (the escape). With the stayOnStack function, we are passing a copy of the value 
// back to the caller, so its safe to keep the value on the stack. With escapeToHeap
// we are passing a copy of the value's address back to the caller 
// (sharing up the stuck) so its not safe to keep the value on the stack. There are 
// all kinds of details related to escape  analysis, so we'll see more about that
// later. 


//go:noinline
func escapeToHeap() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u) // NOTE the use of the pointer in println func

	return &u // NOTE the return of the pointer (address) from the func which fulfills
	// the return type of *user in the func sig.
}

// NOTES:
// 		- When a value could be referenced after the function that constructs the value
// 		  returns
// 		- When the compiler determines a value is too large to fit on the stack
// 		- When the compiler doesn't know the size of a value at compile time
// 		- When a value is decoupled through the use of a function or interface values. 
// ITS TO THE HEAP WITH YOU!

/*
// See escape analysis and inlining decisions.

$ go build -gcflags -m=2
# github.com/ardanlabs/gotraining/topics/go/language/pointers/example4
./example4.go:24:6: cannot inline createUserV1: marked go:noinline
./example4.go:38:6: cannot inline createUserV2: marked go:noinline
./example4.go:14:6: cannot inline main: function too complex: cost 132 exceeds budget 80
./example4.go:39:2: u escapes to heap:
./example4.go:39:2:   flow: ~r0 = &u:
./example4.go:39:2:     from &u (address-of) at ./example4.go:46:9
./example4.go:39:2:     from return &u (return) at ./example4.go:46:2
./example4.go:39:2: moved to heap: u

// See the intermediate representation phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
CALL	"".createUserV1(SB)
	0x0026 00038 MOVQ	(SP), AX
	0x002a 00042 MOVQ	8(SP), CX
	0x002f 00047 MOVQ	16(SP), DX
	0x0034 00052 MOVQ	24(SP), BX
	0x0039 00057 MOVQ	AX, "".u1+40(SP)
	0x003e 00062 MOVQ	CX, "".u1+48(SP)
	0x0043 00067 MOVQ	DX, "".u1+56(SP)
	0x0048 00072 MOVQ	BX, "".u1+64(SP)
	0x004d 00077 PCDATA	$1,

// See bounds checking decisions.

go build -gcflags="-d=ssa/check_bce/debug=1"

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) github.com/ardanlabs/gotraining/topics/go/language/pointers/example4/example4.go
  example4.go:15	0x105e281		e8ba000000		CALL main.createUserV1(SB)
  example4.go:15	0x105e286		488b0424		MOVQ 0(SP), AX
  example4.go:15	0x105e28a		488b4c2408		MOVQ 0x8(SP), CX
  example4.go:15	0x105e28f		488b542410		MOVQ 0x10(SP), DX
  example4.go:15	0x105e294		488b5c2418		MOVQ 0x18(SP), BX
  example4.go:15	0x105e299		4889442428		MOVQ AX, 0x28(SP)
  example4.go:15	0x105e29e		48894c2430		MOVQ CX, 0x30(SP)
  example4.go:15	0x105e2a3		4889542438		MOVQ DX, 0x38(SP)
  example4.go:15	0x105e2a8		48895c2440		MOVQ BX, 0x40(SP)

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
 105e340 T main.createUserV1
 105e420 T main.createUserV2
 105e260 T main.main
 10cb230 B os.executablePath
*/
