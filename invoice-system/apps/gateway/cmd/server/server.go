package server

type StartupHandler interface {
	StartWebServer(port uint, terminateChan chan int) error
}
