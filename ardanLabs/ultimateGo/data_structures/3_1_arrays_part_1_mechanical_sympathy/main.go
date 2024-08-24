package main

import "fmt"

/*
Here we are going to be getting into the data structures for Go. Arrays, slices, maps, etc
Coming from other langs, the first thing most people notice is the LACK of data structures

Others (university CS students) also say that you shouldn't use arrays a lot in
your code, but Go has 2 arrays (Array & Slice) and the slice is Go's MOST IMPORANT
DATA STRUCTURE.

So let's first understand why that is. To do this we need to understand a bit more about the machine
and the  mechanical sympathy that Go has for it.

CPU Caches -
There are lots of mechanical differences between processors and their design. In this section, we'll see at
a high level about processors and the semantics that are relatively the same between them all. This semantic
understanding will provide us with the needed mental model for the processor works and the sympathy that we
can provide it.

Each core inside the processor has its own local cache memory (l1 and l2) and a common cache of memory (l3)
used to store/access data and instructions. The hardware threads in each core can access their local l1 and
l2 caches. Daa from l3 or main memory needs to be copied inot the l1 or l2 cache for access. Main memory
is so slow to access. Performance today is about how efficiently data flows through the hardware. If every
piece of data the hardward needs (at any given time) exists only in main memory, our programs will run
slower as compared to the data already being present on the l1 or l2 caches.

So how do we write code that guarantees the data that is needed to execute an instruction is always present
in the l1 or l2 caches. We would need to write code that is mechanical sympathetic with the processor's
prefetcher. The PREFETCHER attempts to predict what data is needed before instructions requirest the data
so its already present in either the l1 or l2 cache in its cache line. The granularity of a 64 bit machine
is 64 bytes. Ths 64 byte block of memory is the cache line. The prefetcher works best when the instructions
being executed create predictable access patterns to memory. One way to create these access patterns is to
construct contiguous blocks of memory and then iterate over that memory performing linear traversaal with a
predictable stride.

The ARRAY IS THE MOST IMPORANT data structure to the hardware because it supports predictable access patterns.
However, the slice is the most important data structure in Go. Slices in Go use Arrays underneath.

We construct an array, every element is equally distant from the next or previous element. As we iterate over
an array, we begin to walk cache line by connected cache line in a preductable stride. The Prefetcher will pick
up on this pattern and begin to efficiently pull the data into the processor, thus reducing data access latency
costs.

For example, let's say we have a big square matrix of memory and a linked list of nodes that match the
number of elements in the matrix. If we perform a traversal across the linked list, and then traverse the
matrix in both directions (column and row), how will the performance of the different traversals compare?
*/

func main() {
	fmt.Println("this is a test")
}

// create a square matrix of 2meg by 2meg
const (
	rows = 2 * 1024
	cols = 2 * 1024
)

// matrix represents a matrix with a large number of columns per row
var matrix [rows][cols]byte

// data represents a data node for our linked list later
type data struct {
	v byte
	p *data
}

// list points to the head of the list
var list *data

func init() {
	var last *data

	// create a link list with the same number of elements
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			// create a new node and link it in
			var d data
			if list == nil {
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d

			// Add a value to all even elements
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}
}

func RowTraverse() int {
	var ctr int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

// Row traverse will have the best performance because it walks through memory, cache line by connected cache
// line, which creates a predictable access pattern. Cache lines can be prefetched and copied into the l1
// and l2 cache before the data is needed.
/////////////////////////////////////////////////////////////////////

func ColumnTraverse() int {
	var ctr int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// ColumnTraverse is the worst by an order of a magnitude because this access pattern crosses over OS
// page boundaries on each memory access. This causes no predictability for cache line prefetching and becomes
// essentially random access memory.
////////////////////////////////////////////////////////////////////////

// Now what about our linked list
func LinkedListTraverse() int {
	var ctr int
	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}
		d = d.p
	}
	return ctr
}

/*
In our examples our linked list is twice as slow as the row traversal mainly becasue there are cache
line misses but few TLB (Transition Lookaside Buffer) misses. A bulk of the nodes connected in the list
exist inside the same OS pages.

So what is TLB. Each running program is given a full memory map of the virtual memory by the OS and that
running program thinks they have all the physical memory on the machine. However, physical memory needs to be
shared with all the running programs. The operating system shares physical memory by breaking the physical
memory into pages and mapping pages to the virtual memory for any given running program. Each OS can decide
the size of a page, but 4k, 8k, 16k are reasonable and common sizes.

The TLB is a small cache inside the processor that helps to reduce latency on translating a virtual address
to a physical address within the scope of an OS page and offset inside the page. A MISS against a TLB can
cause large latencies because now the hardware has tso wait for the OS to scan its page table to locate the
right page for the virtual address in question. If the program is running on a virtual machine (like the cloud)
then the virtual machine paging tabl needs to be scanned first. 
*/
