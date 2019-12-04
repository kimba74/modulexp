package modulexp

import "fmt"

import "strings"

// Hello to the name provided or "you" if name is an empty string
func Hello(name string) string {
	n := "you"
	if name = strings.TrimSpace(name); name != "" {
		n = name
	}
	return fmt.Sprintf("Hello, %s!", n)
}

// Version of the application
func Version() string {
	return "v1.1.0"
}
