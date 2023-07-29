package notes

/*
	Signaling and Channels
	Channels allow Goroutines to communicate with each other through the use of signaling
	semantics. As I learned in another one of Bill's posts, channeling is less about queues
	than they are about signals. Sending  back and forth signals about when work is complete
	and when we shoudl block and unblock, etc. I remember the different types of signals which
	boil down to signals with or without data. These same signal semantics tell us about
	state, orchestration, etc.  
	
	Channels achieve this signaling through the use of sending/
	receiving data or by identifying state changes on the individual channels. Don't architect
	software with the idea of channels being queues, focus on signaling and the semantics that
	simplify the orchestration required. 
	
	Depending on the problem we're solving, we may require different channel semantics. Depending
	on the semantics we need, different architectural choices must be taken.
	
	Language Mechanics
		- Use channels to orchestrate and coordinate Goroutines
			- Focus on the signaling semantics and not the sharing of data.
			- Signaling with data or without data
			- Question their use for synchronizing access to shared state.
				- There are cases where channels can be simpler for this but intially question it.
		- Unbuffered channels: (signals with data)
			- Receive happens before send (which is why send is blocked until it happens and you 
			  can't send over a channel that's not being received from)
			- Benefit: 100% guarantee the signal being sent has been received
			- Cost: Unknown latency on when the signal will be received
		- Buffered channels: (signals with data)
			- Send happens before the Receive (as the channel has space on it to store data and
			  so we can send up to its capacity without waiting for the first Receive to happen)
			- Benefit: Reduce blocking latency between signaling
			- Cost: No guarantee when the signal being sent has been received.
				- The larget the buffer, the less the guarantee
		- Closing channels: (signals without data)
			- Close happens before the Receive (like buffered)
			- Signaling without data
			- Perfect for signaling cancellations and deadlines
		- NIL channels:
			- Send and Receive block
			- Turn off signaling
			- Perfect for rate limiting or short-term stoppages
			
	Design Philosophy
		- If any given Send on a channel CAN cause the sending Goroutine to block:	
			- Be careful with Buffered channels greater than 1
				- Buffers larger than 1 must have reasons/measurements
			- Must know what happens when the sending Goroutine blocks
		- If any given Send on a channel WON'T cause the sending Goroutine to block:
			- We have the exact number of buffers for each send
				- Fan Out pattern
			- We have the buffer measured for max capacity
				- Drop pattern
		- Less is more with buffers
			- Don't think about performance when thinking about buffers. More buffers doesn't 
			  mean that we are going to get a faster program
			- Buffers can help to reduce blocking latency between signaling
				- Reducing blocking latency towards zero does not necessarily mean better 
				  throughput
				- If a buffer of one is giving us good enough throughput, then keep it.
				- Question buffers that are larger than one and measure for size
				- Find the smallest buffer that provides good enough throughput.




*/
