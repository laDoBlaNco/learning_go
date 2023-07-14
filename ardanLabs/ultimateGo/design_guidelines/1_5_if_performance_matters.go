package notes

/*
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


*/
