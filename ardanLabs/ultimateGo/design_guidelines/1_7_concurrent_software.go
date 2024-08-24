package notes

/*

	When we talk about concurrency we are talking about out of order, undefined execution. Taking
	a set of instructions that would otherwise be executed in sequence and finding a way to execute
	them out of order and still produce the same result. When working on a problem we MUST be sure
	and it must be obvious that out of order execution would add value. And yes that out of order
	value may just be the speed.
	
	When we say adds value, we mean that it adds ENOUGH of a performance gain for the cost
	of the complexity. Remember as engineers its all about understanding the costs of our design
	decisions. Depending on our problem, out of order execution may not be possible or even make
	sense all the time.
	
	Its also important to understand that concurrency is NOT THE SAME AS parallelism. Parallelism
	means executing two or more instructions at the same time. This is a DIFFERENT CONCEPT from
	concurrency that typically gets me and many others confused. Parallelism is ONLY possible when
	we have at least 2 cores or hardware threads avaiable on our machine or environment AND we have
	at least 2 Goroutines, each executing instructions independently of each core/hardware thread.
	
	Both we and the runtim are responsible for managing any concurrency in our application. We 
	are responsible for managing 3 specific things when writing concurrent software:
	
		- 
	
	Design Philosophy:
		- The application MUST startup and shutdown with integrity
			- Know how and when every single  Goroutine we create terminates
			- All Goroutines we create should terminate before main returns
			- Applications should be capable of shutting down on demand, even under load, in
			  a controlled way. 
			  	- We must stop accepting new  requests and finish the requests we have (load
			  	  shedding) for example.
		- Identify and monitor critical points of back pressure that can exist in our application
			- Channels, mutexes and atomic functions can create back pressure when Goroutines are
			  required to wait.
			- A little back pressure is good, it means that there is a good balance of concerns
			- A lot of back pressure is bad, it means things are imbalanced. 
			- Back pressure that is imbalanced will cause:
				- Failures inside the software and across the entire platform
				- Our application will collapse, implode, or freeze.
			- Measuring back pressure is a way to measure the health of the application.
		- Rate limit to prevent overwhelming back pressure inside our applications.
			- Every system has a breaking point, and we must know what that point is for our apps
			- Applications should reject new requests as early as possible once they are overloaded
				- Don't take in more work than we can reasonably handle at a time
				- Push back when we are at critical mass. Create our own external back pressure.
			- Use an external system for rate limiting when its is reasonable and practical
			  to do so.
		- Use timeouts to release the back pressure inside our applications.
			- No request or task should be allowed to take forever
				- So the pressure is that latency or waste of energy on those unfinished tasks
				  The energy or electrons its taking to keep everything running. If they never end
				  then that pressure would continue on forever, eventually taking the system to
				  its brink.
			- Identify how long users are willing to wait
			- Higher-level calls should tell lower-level calls how long they have to run
			- At the top-level,  the user should decide how long they are willing to  wait
			- Use the Context package.
				- Functions that users wait for should take a Context
					- These functions should select on <-ctx.Done() when they would otherwise
					  block indefinitely
				- Set a timeout on a Context only when we have good reason to expect that a
				  function's execution has a real time limit.
				- Allow the upstream caller to decide when the Context should be cancelled
				- Cancel a Context whenever the user abandons or explicitly aborts a call.
		- Architect or design applications to:
			- Identify problems when they are happening
			- Stop the bleeding
			- Return the system back to a normal state.


*/
