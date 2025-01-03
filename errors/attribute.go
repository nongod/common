package errors

// attributeKey is an enum type for identifying error attributes.
type attributeKey uint

const (
	attributeKeyCode       attributeKey = 1 // Attribute key for error code.
	attributeKeyHttpStatus attributeKey = 2 // Attribute key for HTTP status code.
)
