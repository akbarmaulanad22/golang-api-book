package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(result)
	PanicIfError(errDecode)
}

func WriteToResponse(w http.ResponseWriter, response interface{})  {
	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)
	errEncode := encoder.Encode(response)
	PanicIfError(errEncode)
}