package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	// Grab the operating system.
	osID := osString(runtime.GOOS)
	if osID == unspecifiedLinux {
		osID = whichLinux()
	}

	// Then look up the appropriate command for
	// the operating system.
	commandMap := map[osString][]string{
		windows:      {"clip.exe"},
		darwin:       {"pbcopy"},
		xorgLinux:    {"xclip", "-selection", "c"},
		waylandLinux: {"wl-copy"},
	}
	if _, ok := commandMap[osID]; !ok {
		log.Fatalf("unsupported operating system: %s\n", osID)
	}
	cmd := commandMap[osID]

	// Then build the command.
	clipCmd := exec.Command(cmd[0], cmd[1:]...)
	clipCmd.Stdin = os.Stdin
	clipCmd.Stderr = os.Stderr
	clipCmd.Stdout = os.Stdout

	// And execute it
	if err := clipCmd.Run(); err != nil {
		log.Fatalf("failed to paste to clipboard: %s\n", err)
	}
}

type osString string

const (
	windows          osString = "windows"
	darwin           osString = "darwin"
	xorgLinux        osString = "xLinux"
	waylandLinux     osString = "wLinux"
	unspecifiedLinux osString = "linux"
)

// whichLinux assumes that runtime.GOOS is linux.
// It will then return the proper osString corresponding to it.
// If none of the tests work, then runtime.GOOS is returned.
func whichLinux() osString {
	// Check for WSL.
	if isWSL() {
		return windows
	}
	// Check for X11 or Wayland.
	//
	// Ref: https://unix.stackexchange.com/questions/202891/how-to-know-whether-wayland-or-x11-is-being-used
	switch os.Getenv("XDG_SESSION_TYPE") {
	case "x11":
		return xorgLinux
	case "wayland":
		return waylandLinux
	default:
		return osString(runtime.GOOS)
	}
}

// isWSL determines if the the caller is running under the Windows Subsystem for Linux.
//
// Ref: https://github.com/microsoft/WSL/issues/423#issuecomment-221627364
func isWSL() bool {
	kernelInfo, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return false
	}
	return strings.Contains(string(kernelInfo), "microsoft")
}
