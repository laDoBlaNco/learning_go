package notes

/*

	We need to compute less to get the results we need. That's performance in a nutshell. Its about
	not wasting effort and achieving execution efficiency. Writing code that is MECHANICALLY SYMPATHETIC
	with the runtim, OS, and hardware. Achieving performance by writing less and more efficient code
	but staying within the idioms and framework of the language. 
	
	I need to follow these rules:
		- Never guess about performance
		- Measurements must be relevant
		- Profile before I decide that something is performance critical (if not then we are guessing)
		- Test to validate correctness
	
	We need to remember that performance is important but its not the first priority and its not a 
	priority until we can prove that its not running FAST ENOUGH. The only way to know this is to have
	a working program that we have validated. The industry places those who we think know how to write
	performant code on a pedestal, but in reality its those that write code that is optimized for
	correctness and performs FAST ENOUGH that should be on those pedestals.
	
	Micro-Optimizations aare about squeezing every ounce of performance out of te machine as possible.
	When code is written with this as a priority, it's very difficult to write code that is readable, 
	simple or idiomatic. 
		- "When we're computer programmers we're concentrating on the intricate little fascinating 
		   details of programming and we don't take a broad engineering point of view about trying to
		   optimize the who system. You try to optimize the bits and bytes." - Tom Kurtz (inventor of BASIC)

		   	
	If performance matters than what should matter to me as Go dev.

	Four reasons why my software isn't running as fast as it could be:

	- EXTERNAL LATENCY (milliseconds of latency typically). These are things like external system calls
	  or reading from disk or the network.

	- INTERNAL LATENCY (microseconds of latency) - these are things that are in our code. They are things
	  that we control and will definitely talk about during this training. Two big parts are the GC and
	  also concurrency (synchronization and orchestration)

	- DATA ACCESS ON THE MACHINE - the machine is our model, so how we access and read data on our machine
	  can impact our performance as well

	- ALGORITHMS EFFICIENCIES - typically won't hurt us, unless we are tight loops. Hardware is at a level
	  that this shouldn't hurt us too much and again readability is more important, so if we do have a less
	  efficient algo to be more readable, then its probably worth it.
	  
	DATA ORIENTED design is a core philosophy and concept in the Go language. I must embrace it and
	learn to love the fact that data oriented design has a prototype first approach.
	  		- "Data dominates. If you've chosen the right data structures and organized things well,
	  		  the algorithms will almost always be self-evident. Data strutures, not algorithms, are
	  		  central to programming." - Rob Pike (co-creator of Go) 
	
	- The Design Philosophies
		- If we don't understand the data, then we don't understand the problem.
		- All problems are unique and specific to the data we are working with.
		- Data transformations are at the heart of solving problems. Each function, method and
		  work-flow must focus on implementing the SPECIFIC data transformations required to solve
		  a SPECIFIC problem.
		- If our data is changing, then our problems are changing. If my problems are changing, then
		  the data transformations also need to change with them.
		- Uncertainty about the data is not a license to gues but a directive to STOP AND LEARN MORE
		- Solving problems that we don't have, creates more problems that we now need to solve.
		- If performance matters, then we MUST have mechanical sympathy for how the hardware and OS
		  work
		- Minimize, simplify and REDUCE the amount of code required to solve each problem. We do less
		  work by not wasting efforts
		- Code that can be reasoned about and does not hide execution costs can be better understood,
		  debugged and performance tuned. 
		- Coupling data together and writing code that produces predictable access patterns to the 
		  data, will be most performant.
		- Changing data layouts can yield more significant performance improvements than changing just
		  the algorithms
		- Efficiency is obtained through algorithms but performance is obtained through data structures
		  and layouts. 


*/
