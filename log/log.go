package log
import (
	"io"
	"os"
	"sync"
)
type Logger struct {
	mu  sync.Mutex
	out io.Writer
}
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}
func (l *Logger) OutPut(message string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, err := l.out.Write([]byte(message))
	return err
}
var std = New(os.Stderr)
func New(out io.Writer) *Logger {
	return &Logger{out: out}
}
func Fatalln(message string) {
    std.OutPut(message)
	os.Exit(1)
}
func Print(message string) {
	std.OutPut(message)
}
func SetOutput(w io.Writer) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.out = w
}

// go run /home/psycho/go/src/web/log/log.go
