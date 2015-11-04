package main
/*
The follow list maybe need to care:
1. when post the article to the sever,make sure the content
   isnt contains the character sequence like "^^^"
2. before run the application, make sure inserted the administrator account
3. make sure check the consts, whether they are all suit for the environment which the server depending

*/
import (
	"fmt"
	"os"
	"os/signal"
	"web/log"
	"web/server"
)
const (
    // the log file path
	LOG_FILE = "app.log"
    // the address 
	ADDRESS   = "localhost:9091"
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
            server.Close()
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
	server.StartNewServer(ADDRESS)
}
// go run /home/psycho/go/src/web/main.go
