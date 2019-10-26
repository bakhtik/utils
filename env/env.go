// Package env provides means to manage environment variables.
package env

import (
	"bufio"
	"os"
	"strings"
)

// SetEnvironment sets environment variables from the provided filenames.
//
// Files with environment variables shall have the following format:
// VariableName1=VariableValue1
// VariableName1=VariableValue1
// ...
func SetEnvironment(filenames ...string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			text := scanner.Text()
			index := strings.Index(text, "=")
			if index == -1 {
				continue
			}
			err = os.Setenv(text[:index], text[index+1:])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
