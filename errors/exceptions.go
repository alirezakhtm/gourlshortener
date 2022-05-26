package errors

type UrlException struct {
	Message string
}

func (uex UrlException) Error() string  {
	return uex.Message
}