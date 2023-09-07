package main

// Server is our exported implementation and not an interface of it
type Server struct {
	host string
	// more fields
}

// NewServer returns a pointer to our concrete implementation
func NewServer(host string) *Server {
	return &Server{host}
}

// Creating the API off our concrete implemenation
func (s *Server) Start() error {
	return nil
}
func (s *Server) Stop() error {
	return nil
}
func (s *Server) Wait() error {
	return nil
}

func main(){
	
	srv := NewServer("localhost") 
	
	srv.Start() 
	srv.Stop() 
	srv.Wait() 
}

// Same result as the first example, but without the pollution and much simpler to understand 
// and use. 
