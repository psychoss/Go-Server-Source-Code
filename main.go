package main
import (
	//"flag"
	"fmt"
	"os"
	"os/signal"
	"web/log"
	"web/server"
)
const (
	LOG_FILE = "app.log"
	ADREES   = "localhost:9091"
)
var (
	logFile *os.File
)
func handleExitSignal() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
    go func() {
		for sig := range s {
			fmt.Println(sig)
			log.Print("Program is exiting.\n")
			logFile.Close()
			os.Exit(1)
		}
	}()
}
func init() {
	logFile, _ = os.OpenFile(LOG_FILE, os.O_RDWR|os.O_APPEND, 0660)
	log.SetOutput(logFile)
	handleExitSignal()
}
func main() {
	server.StartNewServer(ADREES)
}
// go run /home/psycho/go/src/web/main.go