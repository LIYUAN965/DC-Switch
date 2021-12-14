package swagger

// RedocOpts configures the Redoc middlewares
type RedocOpts struct {
	// SpecURL the url to find the spec for
	SpecURL string
	// RedocURL for the js that generates the redoc site, defaults to: https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js
	RedocURL string
	// Title for the documentation site, default to: API documentation
	Title string
}

type DocsOpts struct {
	DocsURL string
	Title   string
}

func (d *DocsOpts) EnsureDefaults() {
	if d.DocsURL == "" {
		d.DocsURL = "/static/swagger.json"
	}
	if d.Title == "" {
		d.Title = "API documentation"
	}
}

// EnsureDefaults in case some options are missing
func (r *RedocOpts) EnsureDefaults() {
	if r.SpecURL == "" {
		r.SpecURL = "/static/swagger.json"
	}
	if r.RedocURL == "" {
		r.RedocURL = redocLatest
	}
	if r.Title == "" {
		r.Title = "API documentation"
	}
}
