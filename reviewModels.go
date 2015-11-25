package main

// ReviewJSON holds the expected JSON request info
// for the POST /review endpoint which expect product
// information
type ReviewJSON struct {
	Supplier string `json:"supplier,omitempty"`
	Part     string `json:"part,omitempty"`
	Model    string `json:"model,omitempty"`
}
