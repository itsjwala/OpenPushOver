package notification

import (
	"errors"
	"strings"
	"runtime"
	"github.com/sirupsen/logrus"
)

type Message struct {
	Title      string
	Body       string
	Icon       string
	Urgency    string
	ExpireTime int
	Category   string
	Hint       string
	Sound      string
}

const (
	LowPriority      = "low"
	NormalPriority   = "normal"
	CriticalPriority = "critical"
)

var log = logrus.New()

// Errors
var (
	ErrTitleMsg = errors.New("Notification: A title or message must be specified")
)

type NotificationErr struct {
	File   string
	Return string
	Err    error
}

func (e *NotificationErr) Error() string {

	// Usually return will have a newline character
	return e.File + " " + strings.TrimSpace(e.Return) + ": " + e.Err.Error()
}

func (message *Message) Push() (err error){
	switch runtime.GOOS {
		case "linux":
			return linuxPush(message)
		case "windows" : 
			log.Info("Windows's Notifier is not implemented, will just log")
		case "darwin" :
			return darwinPush(message)
		default:
			log.Info("I only know about darwin, linux, windows.. Please add more here and implement")
	}
	return nil
}
