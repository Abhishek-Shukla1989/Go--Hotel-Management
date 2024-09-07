package constant

type ResponseStatus int
type Header int
type General int

const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorised
	ResourceAlreadyExists
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "RESOURCE_ALREADY_EXISTS"}[r-1]

}
func (r ResponseStatus) GetResponseMessage(resource string) string {

	switch r {
	case ResourceAlreadyExists:
		return resource + " already exist"
	default:
		return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
	}
}
