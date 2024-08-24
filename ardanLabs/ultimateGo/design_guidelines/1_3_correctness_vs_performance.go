package notes

/*
	CORRECTNESS VS PERFORMANCE

	- Another issue is when we write code and we focus on performance first. When we do that, we are guessing.
	  We need to focus on correctness vs performance and then we can worry about performance later with all
	  the tools that Go gives us to handle this. We need to benchmark or profile to know if code
	  isn't fast enough. Then and ONLY thenshould we optimize for performance. This can't be done
	  until we have something that is working.
	  
	  Improvement comes from writing code and thinking about the code we've written. Then refactoring
	  the code to make it better. This requires the help of other people to also read the code we
	  are writing, prototype different ideas to validate them, try different approaches and  even 
	  asking others to attempt a solution.  Too many devs aren't prototyping the ideas first before 
	  they are already writing production code. We skip the learning process when we do this and 
	  typically go straight to bad production code. Refactoring them comes in as part of the dev
	  cycle. Its the process of IMPROVING the code from the things we are learning on a daily basis.
	  Without time to refactor, code will becoem impossible to manage and maintain over time. This
	  is the reason we have the legacy issues we see today in the industry.

	  "Make it correct, make it clear, make it concise, than make it fast. In that order." - Wes Dyer

	  "Good engineering is less about finding the 'perfect' solution and more about understanding the
	   tradeoffs and being able to explain them." - JBD

	   We are always drafting code. We don't write perfect code.

	- Everything that makes us get our code correct in the end has to do with READING code.

	- Simple is usually better
	  "Problems can usually be solved with simple, mundane solutions. That means there's no glamorous work.
	  You don't get to show off your amazing skills. You just build something and that gets the job done and
	  then move one. This approach may not earn you oohs and aahs, but it let's you get one with it" - Jason Fried

	  If you want the oohs and aahs be a front-end dev. Back-end devs 'build air conditionars'. No one thinks
	  about the ac until it breaks and all anyone cares about is who is going to come in and fix it. If people
	  know who you are its cuz your code is broke. 



*/
