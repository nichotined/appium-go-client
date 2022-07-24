package property

import "errors"

// Post constant will return "POST"
const Post = "POST"

// Get constant will return "GET"
const Get = "GET"

var ErrMakingRequest = errors.New("error making request")
