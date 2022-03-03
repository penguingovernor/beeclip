package main

import (
	"log"
	"os"

	"github.com/penguingovernor/beeclip/pkg/beeclip"
)

func main() {
	// See what operation to run.
	binary := os.Args[0]
	op := beeclip.CopyToClipboard
	if binary == "beepaste" {
		op = beeclip.PasteFromClipboard
	}

	// Get the command.
	clipCmd, err := beeclip.WhichClipboardCommand(op)
	if err != nil {
		log.Fatalf("failed to get clipboard command: %s\n", err)
	}
	clipCmd.Stdin = os.Stdin
	clipCmd.Stderr = os.Stderr
	clipCmd.Stdout = os.Stdout

	// And execute it
	if err := clipCmd.Run(); err != nil {
		log.Fatalf("failed to paste to clipboard: %s\n", err)
	}
}
