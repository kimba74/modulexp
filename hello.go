package modulexp

import "fmt"

// Hello to the name provided
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Version of the application
func Version() string {
	return "v1.0.0"
}
