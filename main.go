package main

import (
	"appium-go-client/driver"
	"fmt"
)

func main() {
	desiredCaps := map[string]interface{}{
		"deviceName":        "Android",
		"platformName":      "Android",
		"udid":              "emulator-5554",
		"newCommandTimeout": "3600",
		"app":               "/Users/nicholaslagaunne/Workspace/Golife/alia/builds/golife.apk",
		"appPackage":        "com.gojek.life.staging",
		"automationName":    "UiAutomator2",
	}
	driver := driver.CreateDriver("http://127.0.0.1:4723", desiredCaps)

	driver.Init()
	defer driver.Close()

	driver.ImplicitWait(5)

	element := driver.FindElement("id", "input_field")
	element.Click()
	element.SendKeys("")
	fmt.Println(element.Location())
	fmt.Println(element.IsDisplayed())
	fmt.Println(element.GetText())
	fmt.Println(element.GetAttribute("text"))

	// fmt.Println(driver.GetPageSource())
	// fmt.Println(driver.GetSettings())

	// driver.StartActivity(desiredCaps["appPackage"].(string), "com.gojek.golife.presentation.feature.LifeSplashActivity")

}
