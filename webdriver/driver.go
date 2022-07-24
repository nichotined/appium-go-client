package webdriver

import (
	// "encoding/json"
	"encoding/json"

	"github.com/nichotined/appium-go-client/client"
	"github.com/nichotined/appium-go-client/jsonutils"
)

// Driver ...
type Driver struct {
	Client              *client.Client
	DesiredCapabilities map[string]interface{}
	SessionID           string
}

// AppiumRequest ...
type AppiumRequest struct {
	Method string
	Body   map[string]interface{}
	Path   string
}

// Create ...
func Create(url string, capabilities map[string]interface{}) (driver *Driver) {
	driver = &Driver{
		client.CreateClient(url),
		capabilities,
		"",
	}
	return driver
}

func doAppiumRequest(appiumReq *AppiumRequest, c *client.Client, name string) *client.Response {
	resp, err := c.MakeRequest(
		appiumReq.Method,
		jsonutils.StringMapToJSON(appiumReq.Body, name),
		appiumReq.Path,
	)

	if err != nil {
		panic(err)
	}

	return &resp
}

// Init ...
func (d *Driver) Init() {
	appiumReq := &AppiumRequest{
		"POST",
		d.DesiredCapabilities,
		"/session",
	}

	resp := doAppiumRequest(appiumReq, d.Client, "desiredCapabilities")

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"Appium-GO: Unable to create session. Please check if desired capabilities are correct.",
	)

	mapBody := jsonutils.JSONToMap(resp.Body)

	err := json.Unmarshal(*mapBody["sessionId"], &d.SessionID)
	if err != nil {
		panic(err)
	}
}

// Close ...
func (d *Driver) Close() {
	appiumReq := &AppiumRequest{
		"DELETE",
		nil,
		"/session/" + d.SessionID,
	}

	resp := doAppiumRequest(appiumReq, d.Client, "")

	statusCodeErrorHandler(
		resp.StatusCode, 500,
		"Appium-GO: Unable to close session",
	)
}

// GetPageSource ...
func (d *Driver) GetPageSource() string {
	appiumReq := &AppiumRequest{
		"GET",
		nil,
		"/session/" + d.SessionID + "/source",
	}

	resp := doAppiumRequest(appiumReq, d.Client, "")

	statusCodeErrorHandler(
		resp.StatusCode,
		500,
		"Appium-GO: Unable to get source",
	)
	source := jsonutils.JSONToMap(resp.Body)
	return string(*source["value"])
}

// UpdateSettings ...
func (d *Driver) UpdateSettings(settings map[string]interface{}) {
	reqBody := map[string]interface{}{
		"settings": settings,
	}
	appiumReq := &AppiumRequest{
		"POST",
		reqBody,
		"/session/" + d.SessionID + "/appium/settings",
	}

	resp := doAppiumRequest(appiumReq, d.Client, "")
	statusCodeErrorHandler(
		resp.StatusCode,
		500,
		"Appium-GO: Unable to update settings",
	)
}

// GetSettings ...
func (d *Driver) GetSettings() string {
	appiumReq := &AppiumRequest{
		"GET",
		nil,
		"/session/" + d.SessionID + "/appium/settings",
	}

	resp := doAppiumRequest(appiumReq, d.Client, "")
	statusCodeErrorHandler(
		resp.StatusCode,
		500,
		"Appium-GO: Unable to get settings",
	)
	source := jsonutils.JSONToMap(resp.Body)
	return string(*source["value"])
}

// StartActivity ...
func (d *Driver) StartActivity(appPackage string, appActivity string) {
	reqBody := map[string]interface{}{
		"appPackage":  appPackage,
		"appActivity": appActivity,
		// "appWaitPackage":          "",
		// "intentAction":            "",
		// "intentCategory":          "",
		// "intentFlags":             "",
		// "optionalIntentArguments": "",
		// "dontStopAppOnReset":      "",
	}

	appiumReq := &AppiumRequest{
		"POST",
		reqBody,
		"/session/" + d.SessionID + "/appium/device/start_activity",
	}

	res := doAppiumRequest(appiumReq, d.Client, "")

	if res.StatusCode != 200 {
		panic("Driver: Unable to start activity")
	}
}
