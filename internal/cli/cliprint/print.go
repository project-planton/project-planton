package cliprint

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintDefaultSuccess() {
	fmt.Printf("success %s\n", GreenTick)
}

func PrintSuccessMessage(msg string) {
	fmt.Printf("%s %s\n", msg, GreenTick)
}

func PrintError(error string) {
	fmt.Printf("%s %s\n", error, RedTick)
}

// PrintStep prints a step in the process with a blue dot
func PrintStep(msg string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s %s\n", BlueDot, cyan(msg))
}

// PrintSuccess prints a success message with a green checkmark
func PrintSuccess(msg string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s %s\n", CheckMark, green(msg))
}

// PrintInfo prints an informational message with a package icon
func PrintInfo(msg string) {
	fmt.Printf("%s %s\n", Package, msg)
}

// PrintHandoff prints a handoff message when transitioning to external tools
func PrintHandoff(tool string) {
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	
	fmt.Println()
	fmt.Printf("%s %s\n", Handshake, cyan("Handing off to "+tool+"..."))
	fmt.Printf("   %s\n", yellow("Output below is from "+tool))
	fmt.Println()
}
