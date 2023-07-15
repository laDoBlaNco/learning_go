package main

// Now let's talk a little bit more about the heap. Go uses a concurrent mark and sweep collector
// We don't need to know too much about implementation, but we'll get into it in more detail in 
// the addendum section later. 

// A GC's job is pretty much the same anywhere. It goes through and determines what values or 
// memory allocations are still in use and what can be thrown out. Also its concurrent, so we should
// be able to do work while this is going on. The Pacer has an important job here as its determines
// when to collect right before we run out of heap memory. 

// We have to 'stop the world' (freeze all go routines on the machine) while the GC runs. It then
// needs to also maintain integrity during that time. so while the GC is running, a write barrier
// is also running to ensure everything stays accurate. The scheduler is waiting to find the right
// time to run and its waiting for all of the goroutines to call functions to freeze everything. 
// This is about 10-30 second microseconds of time. But at the same time we also need to give  25%
// of our processing power (1 of our 4 threads if we have 4 running) to do its  work. This is the 
// cost of the GC or the internal latency. .

// When thinking about the GC we need to think about 2 things:
// 	1. How much STW time is required to get this garbage collection done?
//	   We want to minimize that and we can use certain algorithms to do so
//  2. Then look at the total amount of time it took to do the garbage collection. This is time
// 	   that we aren't running full throttle.
// So the question we need to ask is how can we by sympathatic to the garbage collection cuz its
// a need and it does cost us.
// 	1. How do we help it reduce the amount of STW time?
// 	2. How do we help reduce how long the GC is going to take per collection?
// 	3. How can we reduce the number of garbage collections that we need for any large amount of work?
//		- What this comes down to is that the amount of SWT time + Overall time for GC * number of 
// 		  of GCs it takes vs the pace we are running at, can tell us what the real impact is on our
// 		  work, o sea how much time we are not running at full throttle. Also the pace and number of 
// 		  GCs also tells us how quickly we are filling up the gap on the heap memory. So if we can 
// 		  reduce the number of GCs we can in turn reduce the impact.
// 			- But how do we do this? Just increasing the gap (or increasing the heap memory)?
// 			  No, cuz the small heap helps us. Smaller heap means less memory to check. If we make it
// 			  bigger, then it'll increase the time it takes for GC. So in reality less heap is less
// 			  work and thus quicker. 
// 			- We should worry about the pace. The pace keeps us at top possible throttle. What we 
// 			  want to to do is use the profilers to reduce allocations and reduce the amount of 
// 			  allocations collected and in turn less GCs. But this is all during profiling and 
// 			  optimizing for performance, after we have a correct working program. 
