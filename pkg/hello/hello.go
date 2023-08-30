package hello

import "fmt"

func Hello(subject string) string {
	return fmt.Sprintf("Hello, %s!", subject)
}
