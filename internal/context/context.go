package context

// Key is a custom type for context keys to avoid collisions.
type Key string

// String returns the string representation of the context key.
func (k Key) String() string {
	return "myapp context key " + string(k)
}

// Exported context keys.
const (
	UserKey Key = "user"
)
