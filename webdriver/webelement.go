package webdriver

import (
	"appium-go-client/jsonutils"
	"appium-go-client/property"
	"encoding/json"
	"fmt"
	"strconv"
)

// Element ...
type Element struct {
	Driver *Driver
	ID     string
}

// FindElement ...
func (d *Driver) FindElement(elBy string, elName string) *Element {
	reqBody := map[string]interface{}{
		"using": elBy,
		"value": elName,
	}

	appiumReq := &AppiumRequest{
		property.Post,
		reqBody,
		"/session/" + d.SessionID + "/element",
	}

	resp := doAppiumRequest(appiumReq, d.Client, "")

	if resp.StatusCode != 200 {
		statusCodeErrorHandler(resp.StatusCode, 404, "Driver: the driver was unable to find an element on the screen")
		statusCodeErrorHandler(resp.StatusCode, 400, "Driver: an invalid argument was passed to the findElement function")
	}

	mapBody := jsonutils.JSONToMap(resp.Body)
	value := map[string]string{}

	err := json.Unmarshal(*mapBody["value"], &value)

	if err != nil {
		panic(err)
	}
	return &Element{d, value["ELEMENT"]}
}

// FindElements ...
func (d *Driver) FindElements(elBy string, elName string) (elements []Element) {
	reqBody := map[string]interface{}{
		"using": elBy,
		"value": elName,
	}

	appiumReq := &AppiumRequest{
		property.Post,
		reqBody,
		"/session/" + d.SessionID + "/elements",
	}

	resp := doAppiumRequest(appiumReq, d.Client, "")

	if resp.StatusCode != 200 {
		statusCodeErrorHandler(resp.StatusCode, 404, "Driver: the driver was unable to find an element on the screen")
		statusCodeErrorHandler(resp.StatusCode, 400, "Driver: an invalid argument was passed to the findElement function")
	}

	mapBody := jsonutils.JSONToMap(resp.Body)

	var val []map[string]interface{}

	if err := json.Unmarshal([]byte(*mapBody["value"]), &val); err != nil {
		panic(err)
	}

	for _, el := range val {
		element := Element{d, fmt.Sprintf("%v", el["ELEMENT"])}
		elements = append(elements, element)
	}

	return
}

// Click ...
func (el *Element) Click() {
	appiumReq := &AppiumRequest{
		property.Post,
		nil,
		fmt.Sprintf("/session/%s/element%s/click", el.Driver.SessionID, el.ID)
	}
	resp := doAppiumRequest(appiumReq, el.Driver.Client, "")

	if resp.StatusCode != 200 {
		panic("ERROR WHEN TRY TO CLICK ELEMENT")
	}

}

// SendKeys ...
func (el *Element) SendKeys(keys string) {
	reqBody := map[string]interface{}{
		"value": keys,
	}

	appiumReq := &AppiumRequest{
		property.Post,
		reqBody,
		fmt.Sprintf("/session/%s/element%s/value", el.Driver.SessionID, el.ID)
	}

	resp := doAppiumRequest(appiumReq, el.Driver.Client, "")

	if resp.StatusCode != 200 {
		panic("ERROR WHEN TRY TO SEND KEYS")
	}
}

// Location ...
func (el *Element) Location() string {
	appiumReq := &AppiumRequest{
		property.Get,
		nil,
		"/session/" + el.Driver.SessionID + "/element/" + el.ID + "/location",
	}

	resp := doAppiumRequest(appiumReq, el.Driver.Client, "")

	if resp.StatusCode != 200 {
		panic("ERROR WHEN TRY TO GET ELEMENT LOCATION")
	}
	source := jsonutils.JSONToMap(resp.Body)
	return string(*source["value"])
}

// IsDisplayed ...
func (el *Element) IsDisplayed() bool {
	appiumReq := &AppiumRequest{
		property.Get,
		nil,
		"/session/" + el.Driver.SessionID + "/element/" + el.ID + "/displayed",
	}

	resp := doAppiumRequest(appiumReq, el.Driver.Client, "")

	if resp.StatusCode != 200 {
		panic("ERROR WHEN TRY TO GET ELEMENT DISPLAYED")
	}
	source := jsonutils.JSONToMap(resp.Body)
	result, _ := strconv.ParseBool(string(*source["value"]))
	return result
}

// GetText ...
func (el *Element) GetText() string {
	appiumReq := &AppiumRequest{
		property.Get,
		nil,
		"/session/" + el.Driver.SessionID + "/element/" + el.ID + "/text",
	}

	resp := doAppiumRequest(appiumReq, el.Driver.Client, "")

	if resp.StatusCode != 200 {
		panic("ERROR WHEN TRY TO GET TEXT")
	}
	source := jsonutils.JSONToMap(resp.Body)
	return string(*source["value"])
}

// GetAttribute ...
func (el *Element) GetAttribute(attributeName string) string {
	appiumReq := &AppiumRequest{
		property.Get,
		nil,
		"/session/" + el.Driver.SessionID + "/element/" + el.ID + "/attribute/" + attributeName,
	}

	resp := doAppiumRequest(appiumReq, el.Driver.Client, "")

	if resp.StatusCode != 200 {
		panic("ERROR WHEN TRY TO GET ATTRIBUTE")
	}
	source := jsonutils.JSONToMap(resp.Body)
	return string(*source["value"])
}
