package tools

import (
	"fmt"
	"strings"
	"time"
)

func DateTimeMySQL() string {
	t := time.Now()
	/* Con 02 le indicamos la cantidad de digitos */
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}
