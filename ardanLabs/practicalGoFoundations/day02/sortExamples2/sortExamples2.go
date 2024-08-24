package main

import(
	"fmt"
	"sort" 
)

// Demonstrating a technique for sorting a struct type using different sets of
// multiple fields in comparison. We chain together "Less" functions, each of which
// compares a single field.

// A Change is a record of the source code changes, recording user, language, and
// delta size
type Change struct{
	user string
	language string
	lines int
}

type lessFunc func(p1,p2 *Change)bool // apparently p1 p2 is for pointer 1 pointer 2

// multisorter implements the Sort interface, sorting the changes within
type multiSorter struct{// Note we aren't exporting everything.
	changes []Change
	less []lessFunc
}

// Sort method sorts the arg slice according to the less functions passed to
// OrderedBy
func(ms *multiSorter)Sort(changes []Change){
	ms.changes=changes
	sort.Sort(ms) 
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc)*multiSorter{
	return &multiSorter{
		less:less,
	}
}

// Finish implement sort.Interface - Len
func (ms *multiSorter)Len()int{
	return len(ms.changes)
}
// Swap
func(ms *multiSorter)Swap(i,j int){
	ms.changes[i],ms.changes[j]=ms.changes[j],ms.changes[i] 
}

// Less is part of the sort.Interface. Its implemented this time by looping along
// the less functions until it finds a comparison that discriminates between the 
// two items (one is less than the other). Note that it can call the less function
// twice per call. We could change the functions to return -1,0,1 and reduce the
// number of calls for greater efficiency: an excercise for the reader.
func(ms *multiSorter)Less(i,j int)bool{
	p,q:=&ms.changes[i],&ms.changes[j]
	// try all but the last comparison
	var k int// it's already zero default, not sure why we set as zero again???
	for k=0;k<len(ms.less)-1;k++{
		less:=ms.less[k]
		switch{
		case less(p,q):return true // p<q so we have a decision
		case less(q,p):return false // p>q so we have a decision
		}
		// if its not less or greater than its p==q so try the next comparison
	}
	// All comparisons to here said "equal", so just return whatever the final
	// comparison reports
	return ms.less[k](p,q) 
}

var changes = []Change{
	{"gri", "Go", 100}, 
	{"ken", "C", 150},
	{"glenda", "Go", 200},
	{"rsc", "Go", 200},
	{"r", "Go", 100},
	{"ken", "Go", 200},
	{"dmr", "C", 100},
	{"r", "C", 150},
	{"gri", "Smalltalk", 80},	
}


// Here we start  with our example multi-keys demonstrating sorting a struct type
// using different sets of multiple fields in the comparison, chaining together
// Less functions which compare a single field.
func main(){
	
	// closures that order the change structure - just like the prev example
	user:=func(c1,c2 *Change)bool{
		return c1.user<c2.user 
	}
	language:=func(c1,c2 *Change)bool{
		return c1.language<c2.language
	}
	increasingLines:=func(c1,c2 *Change)bool{
		return c1.lines<c2.lines
	}
	decreasingLines:=func(c1,c2 *Change)bool{
		return c1.lines>c2.lines // reversed > order
	}
	
	fmt.Println("The Original:")
	fmt.Println(changes)
	fmt.Println() 
	
	// Simple use: Sort by user
	OrderedBy(user).Sort(changes)// I need to understand more about this chaining stuff
	fmt.Println("By users:",changes) 
	
	// More examples:
	OrderedBy(user,increasingLines).Sort(changes) 
	fmt.Println("By user, <lines:",changes) 
	
	OrderedBy(user,decreasingLines).Sort(changes)
	fmt.Println("By user,>lines:",changes)
	
	OrderedBy(language,increasingLines).Sort(changes)
	fmt.Println("By language,<lines:",changes) 
	
	OrderedBy(language,decreasingLines).Sort(changes)
	fmt.Println("By language,>lines:",changes) 
	
	OrderedBy(language,increasingLines,user).Sort(changes)
	fmt.Println("By language,<lines,user:",changes) 
}
