package errhandle

var defaultErrOccured bool
var customErrOccured bool
var dataPassed bool

type customErr struct{}

func (ce customErr) Error() string {
	return "custom error"
}
