package examples

type AlertCounter int

// The capital letter here being used means the type is exported and can be referenced
// directly by code outside of the counters package. 
// if it was 'alertCounter' then it would be unexported. This means that only code inside the
// counters package or folder, would be able to reference this type directly.

// But we could still access it indirectly if its methods or other associated functions are
// exported. For example
type alertCounter int

func New(value int) alertCounter{
	return alertCounter(value)
}

// Even though the above would legally compile, there's no value in it since the caller can't 
// reference the name directly if its in another package, which is more than likely. So this proves
// that export / non-export in Go doesn't make things public and private, as we could still access
// So its better not to do the above so as not to confuse others. 

type User struct{
	Name string
	ID int
	
	password string
}

type user2 struct{
	Name string
	ID int
}

type Manager struct{
	Title string
	
	user2 
}

