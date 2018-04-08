package hndlrs

// Loader is an interface used to attach an HTTP handler to a ServeMux
type Loader interface {
	// Load registers an HTTP handler into the provided ServeMux. On
	// success nil is returned, or an error if one occurs.
	Load(mux http.ServeMux) error
}
