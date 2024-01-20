package network

import "fmt"

// Ping - take in an IP an dping that IP before
// returning a response
func Ping(ip string) { // NOTE the capital Ping, this means we want this to be exportable
	// ping lowercase would be a private func only accessible in this network package.
	fmt.Println(ip)
}
