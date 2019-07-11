package rest

import (
	"encoding/json"
	"net/http"

	"github.com/aadi756/cardflow/app/pkg/transaction"
	"github.com/bnkamalesh/webgo"
)

func PinChangeRequest(w http.ResponseWriter, r *http.Request) {
	// wctx := webgo.Context(r)

	payload := make(map[string]interface{}, 0)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// h.ReturnError(err, w, r, nrTxn)
		return
	}
	result , err := transaction.PinChangeRequest()

	webgo.R200(w, result)
}

func PinChangeResendOtp(w http.ResponseWriter, r *http.Request) {
	// wctx := webgo.Context(r)

	payload := make(map[string]interface{}, 0)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// h.ReturnError(err, w, r, nrTxn)
		return
	}

	webgo.R200(w, payload)
}

func PhoneConfirm(w http.ResponseWriter, r *http.Request) {
	// wctx := webgo.Context(r)

	payload := make(map[string]interface{}, 0)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// h.ReturnError(err, w, r, nrTxn)
		return
	}

	webgo.R200(w, payload)
}

func ActivateCard(w http.ResponseWriter, r *http.Request) {
	// wctx := webgo.Context(r)

	payload := make(map[string]interface{}, 0)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// h.ReturnError(err, w, r, nrTxn)
		return
	}

	webgo.R200(w, payload)
}
