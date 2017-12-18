package config

// Validator is used to ensure that data is in the correct form
type Validator interface {
	// Validate inspects the data and returns any errors that are present.
	// Nil if the data is in the correct form
	Validate() error
}
