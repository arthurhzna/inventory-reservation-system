package logging

import (
	"fmt"
	"strings"
)

func argsToString(args ...any) string {
	var sb strings.Builder

	for _, arg := range args {
		sb.WriteString(
			fmt.Sprintf("%v ", arg),
		)
	}

	return strings.TrimSpace(sb.String())
}
