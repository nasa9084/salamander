package salamander

import "github.com/nasa9084/salamander/salamander/log"

type option func(*Server)

// Listen option
//     Example: NewServer(Listen(":8080"))
func Listen(listen string) func(*Server) {
	return func(s *Server) {
		s.listen = listen
	}
}

// LoggingLevel option
//     Example: NewServer(LoggingLevel("debug"))
func LoggingLevel(loggingLv log.Level) func(*Server) {
	return func(s *Server) {
		s.loggingLevel = loggingLv
	}
}
