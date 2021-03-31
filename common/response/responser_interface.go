package response

type Responser interface {
	Success() bool
	Errors() []error
	AddError(err error)
	AddErrors(errs []error)
	WriteJSONResponse() error
}
