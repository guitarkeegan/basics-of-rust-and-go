package datatypes

// the string is metadata, like a comment. reflection api will read it
// lets you inspect types on the fly
type Rate struct {
	Currency string
	Price    float64
}
