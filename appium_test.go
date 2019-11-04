package driver

import (
	"appium-go-client/driver"
	"testing"
)

func TestDriver(t *testing.T) {
	desiredCaps := map[string]interface{}{
		"deviceName":        "Android",
		"platformName":      "Android",
		"udid":              "emulator-5554",
		"newCommandTimeout": "3600",
		"app":               "",
		"appPackage":        "",
		"automationName":    "UiAutomator2",
	}
	driver := driver.CreateDriver("http://127.0.0.1:4723", desiredCaps)
	driver.Init()
}
