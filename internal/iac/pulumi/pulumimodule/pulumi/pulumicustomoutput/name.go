package pulumicustomoutput

import "fmt"

// Name for custom outputs by prefixing output names with x_
func Name(name string, suffix ...string) string {
	customName := fmt.Sprintf("x_%s", name)
	for _, s := range suffix {
		customName = fmt.Sprintf("%s-%s", customName, s)
	}
	return customName
}
