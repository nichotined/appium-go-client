package driver

// ImplicitWait ...
func (d *Driver) ImplicitWait(seconds int) {
	reqBody := map[string]interface{}{
		"ms": seconds * 1000,
	}

	appiumReq := AppiumRequest{
		"POST",
		reqBody,
		"/wd/hub/session/" + d.SessionID + "/timeouts/implicit_wait",
	}

	resp := doAppiumRequest(&appiumReq, d.Client, "")

	if resp.StatusCode != 200 {
		panic("Implicit Wait Error")
	}
}
