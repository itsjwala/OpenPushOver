package notification
import (
	"os/exec"
)
func darwinPush(m *Message) (err error) {
	if len(m.Title) == 0 {
		m.Title = "pushover"
	}
	cmd := exec.Command("terminal-notifier","-title",m.Title,"-message",m.Body)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return &NotificationErr{Return: string(out), Err: err}
	}
	return nil
}
