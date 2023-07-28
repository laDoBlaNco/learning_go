package notes

/*

	Here are the design philosophies and guidelines that we will need to follow when it comes
	to interfaces and composition
	
		- Interfaces give programs structure
		- Interfaces encourage design by COMPOSITION
		- Interfaces enable and enforce clean divisions between components
			- The standardization of interfaces can set clear and consistent expectations
		- Decoupling means reducing the dependencies between components and the types they use
			- This leads to CORRECTNESS, QUALITY, and PERFORMANCE.
		- Interfaces allow me to group concrete types by what they do 
			- Don't group types by a common DNA but by a common behavior
			- Everyone can work together when we focus on what we do and not who we are.
		- Interfaces help my code decouple itself from change.
			- We must our best to understand what could change and use interfaces to decouple
			- Interfaces with more than one method have more than one reason to change
			- Uncetainty about change is not a license to guess but a directive to STOP and 
			  learn more.
		- I must distinguish between code that:
			- Defends against fraud vs protects against accidents
			
	Validation 
	We use interfaces when:
		- Users of an API need to provide an implemenation detail.
		- API's have multiple implementations they need to maintain internally.
		- Parts of the API that can change have been identified and require decoupling.
	Don't use an interface
		- Just for the sake of using one
		- To generalize an algorithm
		- When users can declare their own interfaces
		- If its not clear how the interface makes the code better. 

*/
