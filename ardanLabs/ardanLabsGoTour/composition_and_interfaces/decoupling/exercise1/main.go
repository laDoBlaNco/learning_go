package main

import "fmt"

/*
	Using the template, declare a set of concrete types that implement the set of
	predefined interface types. Then create values of these types and use  them
	to complete a set of predefined tasks below.

*/

// Our administrator reps a person or other entity capable of administering hardware
// and software infrastructures
type administrator interface {
	administrate(system string) // value semantics
}

// Developer reps a person or entity capable of writing software
type developer interface {
	develop(system string)
}

// ==========================================================================
// Now let's look at concrete types
// adminlist reps a group of administrators
type adminlist struct {
	list []administrator
}

// Enqueue adds an administrator to the adminlist
func (l *adminlist) Enqueue(a administrator) { // receives a value but uses pointer semantic receiver
	l.list = append(l.list, a)
}

// Dequeue removes an administrator from the adminlist and returns the value
func (l *adminlist) Dequeue() administrator { // returns a value, so still using value semantics
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

// devlist reps a group of developers
type devlist struct {
	list []developer
}

// Enqueue adds a developer to the devlist
func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

// Dequeue removes a developer from the devlist and returns the value
func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

// 1. Declare a concrete type named sysadmin with a name field of type string
type sysadmin struct {
	name string
}

//  2. Declare a method named administrate for the sysadmin type, implementing the
//     administrator interface. administrate should print out the name of the sysadmin,
//     as well as the system they are administering
func (sa *sysadmin) administrate(sys string) {
	fmt.Printf("System Admin: %s\nSystem: %s\n", sa.name, sys)
}

// 3. Declare a concrete type named programmer with a name field of type string
type programmer struct {
	name string
}

//  4. Declare a method named develop for the programmer type, implementing the
//     developer interface. develop should print out the name of the programmer,
//     as well as the system they are coding.
func (p *programmer) develop(sys string) {
	fmt.Printf("Programmer: %s\nSystem: %s\n", p.name, sys)
}

//  5. Declare a concrete type named company. Declare it as the composition of the
//     administrator and developer interface types.
type company struct { // the key here is that we are embedding interfaces as fields in a
	// concrete struct.
	administrator
	developer
}

//===================================================================================

func main() {

	// 6. Create a variable named admins of type adminlist
	var admins adminlist

	// 7. Create a variable named devs of type devlist
	var devs devlist

	// 8. Enqueue a new sysadmin onto admins
	admins.Enqueue(&sysadmin{"Carmen Sosa"})

	// 9. Enqueue two new programmers onto devs
	devs.Enqueue(&programmer{"Kevin Whiteside"})
	devs.Enqueue(&programmer{"Ana Bonifacio"})

	// 10. Create a variable named cmp of type company, and initialize it by hiring
	//     (dequeueing) an administrator from admins and a developer from devs.
	cmp := company{administrator: admins.Dequeue(), developer: devs.Dequeue()}

	// 11. Enqueue the company value on both lists since the company implements
	//     each interface
	admins.Enqueue(cmp) // NOTE: I did this without the pointer and it still worked. Not
	// sure why, but I'll move on and figure it out later. But since Enqueue uses pointer
	// semantics, I would have expected the same error I got when I tried to use
	// sysadmin and programmer values instead of pointers.
	devs.Enqueue(&cmp) // this should work since the value implements both interfaces
	// and I guess we should only be able to see the one that is implemented then???
	// I guess that's the point. It will dequeue the one that fills the interface so even
	// there are two in cmp, it only sees the one it needs to.

	// Now we have the tasks for administrators and developers to perform
	tasks := []struct { // this would then been a constant literal (non addressable) value
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	// Iterate over the tasks
	for _, task := range tasks {
		// 12. Check if the task needs an administrator else use a developer
		// NOTE in Go we try not to use 'else' explicitly
		if task.needsAdmin {
			admins.Dequeue().administrate(task.system)
			fmt.Println()
			continue
		}

		// 13. Dequeue a developer value from the devs list and call the develop method
		devs.Dequeue().develop(task.system)
		fmt.Println()
	}
}
