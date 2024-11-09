package pulumioutput

import (
	"fmt"
	"reflect"
	"strings"
)

func Name(t reflect.Type, name string) string {
	ts := strings.ReplaceAll(t.String(), "*", "")
	ts = strings.ReplaceAll(ts, ".", "-")
	ts = strings.ToLower(ts)
	return fmt.Sprintf("%s_%s", ts, name)
}
