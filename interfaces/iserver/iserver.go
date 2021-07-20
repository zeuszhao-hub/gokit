package iserver

type Server interface {
	Run() error
	Shutdown() error
}
