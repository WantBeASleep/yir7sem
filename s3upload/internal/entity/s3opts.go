package entity

type LoadOpts struct {
	ContentType string
}

type LoadOption func(o *LoadOpts)

var (
	WithContentType = func(contentType string) LoadOption {
		return LoadOption(func(o *LoadOpts) {
			o.ContentType = contentType
		})
	}
)

type GetOpts struct{}

type GetOption func(o *GetOpts)

var ()
