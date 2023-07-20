package notes

/*
 PREPARE YOUR MIND

 - Somewhere along the line
	- We became impressed with programs that contain large amounts of code
		- For example linux with 25,000,000 loC, Go tries to reduce this
	- We strived to create large abstractions in our codebase
		- Focus should be on thin layers of price decoupling to help with readability and maintainability
	- We forgot that the hardware is the platform
		- If performance matters then its the hardware that matters. This is Go's model.
	- We lost the understanding that every decision comes with a cost
		- Engineering is about understanding the cost of the decisions we are making. If we don't do it,
		  then all we do are hacking around. We 'hack' first, but we need to then start to engineer

 - These days are gone
	- We can throw more hardware at the problem
	- We cn throw more developers at the problem

 - Open our minds
	- Technology changes quickly but people's minds change slowly
	- Easy to adopt new tech but hard to adopt new ways of thinking
		- Go really is different. Things that cost us a lot in other languages, don't cost as much in Go
		  But there are other things that are simple in other languages that cost us in Go. There's a balance
		  and a new way of thinking with Go.

		  We need to learn to read before we learn to write. We need to read code more than we write it. We
		  learn about writing code by reading code.

		  We need a MENTAL MODEL
	- If we can't maintain a mental model of our project then we can't really maintain it or understand it.
	  Once you get to a certain number of LoC you can't really maintain it all in your mind. Our mental model
	  has limits. So in Go we can do more in less lines of code, meaning that we can keep teams small and
	  wrap our minds around more of our code.

	  "The hardest bugs are those where you mental model of the situation is just wrong, so you can't see the
	  problem at all." - Brian Kernighan

	  "Everyone knows that debugging is twice as hard as writing a program in the first place. So if you're
	  as clever as you can be when you write it, how will you ever debug it?" -Brian Kernighan

	  If we just depend on our debuggers to find all the bugs than our mental models will never be complete.
	  If we need a debugger than there's a bigger problem, cuz we aren't going to be able to attach a debugger
	  in production. We'll need to rely on logs and our mental model.

	  "Debuggers don't remove bugs. They only show the in slow motion." - Unknown



 - Interesting Questions - What do they mean?
	- Is it a good program?
	- Is it an efficient program?
	- Is it correct?
    - Was it done on time?
    - What did it cost?

- Aspire To
  - Be a champion for quality, efficiency and simplicity.
  - Have a point a view about what you want to do and see in your code.
  - Value introspection and self-review. This isn't focused on my code reviews, but on me as an engeineer.

READING CODE:
  - Go is a language that focuses on code being readable as a first priciple.
    "Code is read many more times than its is written." - Dave Cheney
    "Programming is, among other things, a kind of writing. One way to learn writing is to write,
    but in all other forms of writing, one also reads. We read examples of both bood and bad to
    facilitate learning. But how many programmers learn to write programs by reading programs?" - Gerald M Wienberg

MENTAL models
  - We must constantly make sure that our mental model of the code we are writing and maintaining
    is clear. When we can't remember where a piece of logic is or we can't remember how something
    works, then we are losing our mental model of the code. This is a clear indication that we need
    to refactor. Focus time on structuring code that provides the best possible mental model and
    during code reviews validate our mental models are still intact.

    How much code do we think we can maintain in our heads? It was said that a single dev can maintain
    a mental model of a program of about 1 ream of paper (~10k LoC). If we do the math, that would mean
    that a millon LoC code base would talk around 100 people to maintain it. That's 100 people that need
    to be coordinated, grouped, tracked and in constant feedback loop of communication.
*/
