package apierrors

// ClientError is an interface for errors that should be returned to the client.
// They generally start with a 4xx HTTP status code.
type ClientError interface {
	error
	HTTPStatusCode() int
}
