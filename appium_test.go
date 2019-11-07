package driver

import (
	"appium-go-client/driver"
	"fmt"
	"os"
	"testing"
)

func TestDriver(t *testing.T) {
	currentDirectory, _ := os.Getwd()
	desiredCaps := map[string]interface{}{
		"deviceName":        "Android",
		"platformName":      "Android",
		"udid":              "emulator-5554",
		"newCommandTimeout": "3600",
		"app":               currentDirectory + "/apps/ApiDemos-debug.apk",
		"appPackage":        "com.example.android.apis",
		"automationName":    "UiAutomator2",
	}
	// APPIUM Server URL Example http://127.0.0.1:4723/wd/hub
	driver := driver.CreateDriver("http://127.0.0.1:4723/wd/hub", desiredCaps)

	driver.Init()
	defer driver.Close()

	driver.ImplicitWait(5)

	elements := driver.FindElements("id", "android:id/text1")
	for _, element := range elements {
		fmt.Println(element.Location())
	}
}
