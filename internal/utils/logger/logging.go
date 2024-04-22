package logger

type Log interface {
	// The implementing type should implement the GenerateLogger method , within which the implemented logger should be assign to KeyCloaklogger variable
	GenerateLogger()
}
