//go:build darwin

package notify

import "github.com/wailsapp/wails/v2/pkg/mac"

func Notification(title string, subtitle string, message string, sound string) error {
	return mac.ShowNotification(title, subtitle, message, sound)
}
