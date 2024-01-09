package logger

type KClogging interface {
	// The implementing type should implement the GenerateLogger method , within which the implemented logger should be assign to KClogger variable
	GenerateLogger()
}
