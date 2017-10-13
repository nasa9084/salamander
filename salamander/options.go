package salamander

type option func(*Server)

// Listen option
//     Example: NewServer(Listen(":8080"))
func Listen(listen string) func(*Server) {
	return func(s *Server) {
		s.listen = listen
	}
}
