package response

import (
	"encoding/json"
	"net/http"
)

// Struct for API Response.
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Set the API code response. Should always call JSON() function.
func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

// Set the API message response. Should always call JSON() function.
func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

// Set the API data response. Should always call JSON() function.
func (r *Response) SetData(data any) *Response {
	r.Data = data
	return r
}

// Render the JSON response. Always call this function after set up
// code, message, data.
func (r *Response) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		panic(err)
	}
}

func (r *Response) Error(err error) *Response {
	return &Response{
		Message: err.Error(),
	}
}
