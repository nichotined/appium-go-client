package webdriver

// TouchActionParams ...
type TouchActionParams struct {
	elementID string
	x         int
	y         int
	duration  int
}

// Tap ...
// func (d *Driver) Tap(touchActionParams TouchActionParams) {
// 	touchActionParams.elementID
// }

// func getOptions(touchActionParams TouchActionParams) {
// 	var opts interface{}

// 	if touchActionParams.elementID != nil {
// 		opts["element"] = touchActionParams.elementID
// 	}

// 	if touchActionParams.x != nil && touchActionParams.y != nil {
// 		opts['x'] = touchActionParams.x
// 		opts['y'] = touchActionParams.y
// 	}

// 	if touchActionParams.duration != nil {
// 		opts["duration"] = touchActionParams.duration
// 	}

// 	return opts
// }
