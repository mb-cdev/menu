package response

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	IsSucceed    bool
	Errs         []string
	ResponseData interface{}
}

func (r *JSONResponse) Success() bool {
	return r.IsSucceed
}

func (r *JSONResponse) Errors() []string {
	return r.Errs
}

func (r *JSONResponse) WriteJSONResponse(w http.ResponseWriter) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}

func (r *JSONResponse) AddError(e error) {
	r.Errs = append(r.Errs, e.Error())
}

func (r *JSONResponse) AddErrors(errs []error) {
	for _, e := range errs {
		r.AddError(e)
	}
}
