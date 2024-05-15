package model

// LoadError represents an error that occurred during the loading process.
type LoadError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// LoadResult represents the result of a loading process, including any errors.
type LoadResult struct {
	Tests      []TestCase
	LoadErrors []LoadError
}
