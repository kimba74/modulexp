package modulexp

import "fmt"

import "strings"

type stringer interface {
	String() string
}

// Hello to the name provided or "you" if name is an empty string
func Hello(name interface{}) string {
	var vname string
	if name != nil {
		switch s := name.(type) {
		case string:
			vname = s
		case stringer:
			vname = s.String()
		default:
			vname = "you"
		}

	}
	if vname = strings.TrimSpace(vname); vname == "" {
		vname = "you"
	}
	return fmt.Sprintf("Hello, %s!", vname)
}

// HelloWorld greeting
func HelloWorld() string {
	return Hello("world")
}

// Version of the application
func Version() string {
	return "v2.0.0"
}
