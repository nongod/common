package code

// NilCode represents an empty value for the Code type.
const (
	NilCode Code = ""
)

// Code is a general-purpose identifier type used for defining various targets (e.g., error codes, action codes, etc...).
type Code string

// Equal compares two Code values for equality.
// Returns true if both Code values are the same.
func (c Code) Equal(v Code) bool {
	return c == v
}

// IsNil checks if the Code is an empty value (NilCode).
func (c Code) IsNil() bool {
	return c == NilCode
}

// String returns the string representation of the Code.
func (c Code) String() string {
	return string(c)
}

// Provider is an interface for types that can provide a Code.
type Provider interface {
	Code() Code
}

// From extracts a Code from a type that implements the Provider interface.
// If the type does not implement Provider, it returns NilCode.
func From(v any) Code {
	if coder, ok := v.(Provider); ok {
		return coder.Code()
	}

	return NilCode
}
