package utils

import (
	"bytes"
	"fmt"
)

func ConvertMapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\", ", key, value)
	}
	return b.String()
}
