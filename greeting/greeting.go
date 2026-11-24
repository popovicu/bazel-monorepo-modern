package greeting

import "fmt"

// For returns a greeting for name.
func For(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
