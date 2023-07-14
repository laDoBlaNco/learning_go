package notes

/*
	What are the rules for optimizing for correctness. Integrity is about becoming very serious about
	reliability. There are 2 driving factors (micro-level and macro-level).  The only thing code does is
	1. allocate memory, 2. read memory 3. write to memory. Even more specific all its doing is reading and
	writing integers all day (floating points for GPUs).

	- At a micro-level  integrity means that everyone of those reads and writes of memory need to be accurate,
	  consistent and effiencient. The type system is critical to making sure we have this micro level of
	  integrity. If we lose any of those things, we have integrity issues and 'corruption'

	- At a macro-level of integrity, everything we do and every problem we solve is a data-problem, its a
	  data transformation problem. Go is aa data-oriented (not object oriented) language. Data comes in, we
	  translate it to another type and then it goes out. If we don't understand the data then there's no way
	  we can even begin to understand the problem or solve it. WE ARE ALL DATA SCIENTISTS at the end of the day.

	  So these data transfomations also have to be accurate, consistent, and efficient. Writing less code
	  and error handling is critical to making sure we have this macro level of integrity.

	INTEGRITY must be first. Nothing trumps integrity. So again the two things that help us is writing less
	code (loC) and handling the errors.
	- researching of bugs shows that the industry average is around 15 to 50 bugs per 1000 LoC.So one simple
	  way to reduce bugs and increase integrity is to WRITE LESS CODE. This is what Go is trying to do. Do
	  the math. Less code less bugs.
	  - We should also only introduce code that is needed in that version of the product. The minimal amount
	    of code needed for that version and keeping it clean.

	- error handling is tedious, but the best security is tedious. Go's error handling mechanics are tedious
	  but that's what's needed when optimizing for correctness. Example:
	  - There were 48 crictial failures found in a study looking at a couple hundred bugs in Cassandra, HBase,
	    MapReduce, and Redis. Of those 48:
		- 92%: Failures from BAD ERROR HANDLING
		  - 35%: Incorrect handling
		    - 25%: Simply ignoring an error
			- 8%: Catching the wrong exception
			- 2%: Incomplete TODOs
	  - "Failure is expected, failure is not an odd case. Design systems that help you identify failure.
	    Design systems that can recover from failure." - JBD
	  - "Product excellence is the difference between something that only works under certain conditions,
	    and something that only breaks under certain conditions." - Kelsey Hightower

	INTEGRITY MUST BE TAKEN SERIOUSLY. Its the number 1 thing. The cost of integrity is typically performance.
	We are going to lose a little bit of performance (maybe just nanoseconds) for ensuring integrity, but its
	worth it in the end.

	After integrity comes READIBILITY. We must design our systems to be more comprehensible. There are 2 parts
	to this.
	- Subjective - The average developer on the team should have a full mental model of the codebase and everyone
	  should be able to read  the language clearly.
	- Practical - We don't hide cost. We can add abstraction, but we can't  hide the cost of the decisions we
	  are making. When you read code that has all kinds of "features" to make you life easier, you really have
	  no idea what its doing in the background. You can't tell when things are being run, how they are being
	  executed, how many objects are being created, etc. Cost is hidden from us. Go doesn't having all these
	  "features" because they hide cost. We want to be able to look at code and know reasonable how well it's
	  going to run.

	The last thing to keep in mind with code reviews is SIMPLICITY.
	- "Making things easy to do is a false economy. Focus on making things easy to understand and the rest
	  will follow." - Peter Bourgon
	- Don't make things easy to do, make them easy to understand. Sure this might be more tedious but the point
	  isn't about writing easy code, its about maintaining it after. We handle the problems when things are
	  failing, not when they are good and all work. That's the important part.
	- Simplicity is about 'hiding' complexity without losing 'readability' that's the challenge.
	- Simplicity is something you have to refactor to. You can't achieve it day one because of this battle
	  between simplicity and readability (hiding complexity vs losing readability). The process day in and day
	  out is
	  - 1. get it to work
	  - 2. review readability
	  - 3. clean our comprehension of the code and clean up our mental models of the codebase
	  - 4. refactor to hide some of the complexity. (creating levels of decoupling, etc)

	INTEGRITY -> READABILITY -> SIMPLICITY




*/
