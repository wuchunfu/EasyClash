//go:build windows

package notify

func Notification(title string, subtitle string, message string, sound string) error {
	return nil
}
