package http

type HTTPApi struct {
	UrlPrefix string
}

func New() *HTTPApi {
	httpapi := HTTPApi{}
	httpapi.UrlPrefix = "/api/v1"
	return &httpapi
}
