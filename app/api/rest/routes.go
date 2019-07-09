package rest

import (
	"net/http"

	"github.com/bnkamalesh/webgo"
	"github.com/bnkamalesh/webgo/middleware"
)

// Routes returns a list of handlers defined for webGo
func GetRoutes() []*webgo.Route {
	return []*webgo.Route{
		&webgo.Route{
			Name:     "AtmPinChangeResendOtp",                                              // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodPost,                                           // request type
			Pattern:  "/card/pin/otp/resend/sms",                                // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), PinChangeResendOtp}, // route handler
		},
		&webgo.Route{
			Name:     "AtmPinChange",                                            // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodPost,                                         // request type
			Pattern:  "/card/pin/otp/request",                                 // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), PinChangeRequest}, // route handler
		},
		&webgo.Route{
			Name:     "PhoneConfirm",                                            // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodPost,                                         // request type
			Pattern:  "/card/pin/otp/resend/phone",                                 // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), PhoneConfirm}, // route handler
		},
		&webgo.Route{
			Name:     "ActivateCard",                                            // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodPost,                                         // request type
			Pattern:  "card/pin/confirm",                                 // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), ActivateCard}, // route handler
		},
	}
}