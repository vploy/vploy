package core

type HttpClient interface {
	//
}

type xHttpClient struct {
	opts HttpOptions
}

func (x *xHttpClient) Request(method, uri string) {
	//
}

// =======================

type xHttpNS struct{}

func (x *xHttpNS) New(opts HttpOptions) (client HttpClient) {
	client = &xHttpClient{opts: opts}
	return
}

var HTTP xHttpNS
