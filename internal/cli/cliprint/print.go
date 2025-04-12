package cliprint

import "fmt"

func PrintDefaultSuccess() {
	fmt.Printf("success %s\n", GreenTick)
}

func PrintSuccessMessage(msg string) {
	fmt.Printf("%s %s\n", msg, GreenTick)
}

func PrintError(error string) {
	fmt.Printf("%s %s\n", error, RedTick)
}
