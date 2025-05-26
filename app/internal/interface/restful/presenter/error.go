package presenter

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SourceResponse struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

type ErrorResponse struct {
	Code   int             `json:"code,omitempty"`
	Detail string          `json:"detail,omitempty"`
	Source *SourceResponse `json:"source,omitempty"`
}

const errMsg = "Code: '%d' Detail:'%s' - Source: '%+v'"

// Error method implement built-in error interface
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf(errMsg, e.Code, e.Detail, e.Source)
}

// ErrorResponses is the list of ErrorResponse
type ErrorResponses []ErrorResponse

// Append a new error response to the list of errors
func (e *ErrorResponses) Append(newE ErrorResponse) {
	*e = append(*e, newE)
}

// Error method implement built-in error interface
func (e ErrorResponses) Error() string {
	buff := bytes.NewBufferString("")

	var fe ErrorResponse

	for i := 0; i < len(e); i++ {
		fe = e[i]
		buff.WriteString(fe.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// FromValidationErrors converts from validator.ValidationErrors to presenter.ErrorResponse
func (e *ErrorResponses) FromValidationErrors(vldErrs validator.ValidationErrors) {
	for i := range vldErrs {
		*e = append(*e, ErrorResponse{
			Code:   http.StatusUnprocessableEntity,
			Detail: vldErrs.Error(),
			Source: &SourceResponse{
				Pointer:   vldErrs[i].StructNamespace(),
				Parameter: vldErrs[i].StructField(),
			},
		})
	}
}
