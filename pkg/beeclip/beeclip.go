package beeclip

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type ClipboardOperation uint

const (
	CopyToClipboard = ClipboardOperation(iota)
	PasteFromClipboard
)

type clipboardCommands struct {
	copyCommand  string
	pasteCommand string
}

func WhichClipboardCommand(op ClipboardOperation) (*exec.Cmd, error) {
	commands, ok := operatingSystemToClipboardCommands[determineOperatingSystemID(runtime.GOOS)]
	if !ok {
		return nil, fmt.Errorf("could not determine operating system")
	}

	commandInPath := func(cmd string) (*exec.Cmd, error) {
		fields := strings.Fields(cmd)
		if _, err := exec.LookPath(fields[0]); err != nil {
			return nil, err
		}
		return exec.Command(fields[0], fields[1:]...), nil
	}

	switch op {
	case CopyToClipboard:
		return commandInPath(commands.copyCommand)
	case PasteFromClipboard:
		return commandInPath(commands.pasteCommand)
	}

	return nil, fmt.Errorf("unknown operation: %d", op)
}

type operatingSystemID uint

const (
	operatingSystemWindows = operatingSystemID(iota)
	operatingSystemDarwin
	operatingSystemXorgLinux
	operatingSystemWaylandLinux
	operatingSystemUnknown
)

var operatingSystemToClipboardCommands = map[operatingSystemID]clipboardCommands{
	operatingSystemWindows: {
		copyCommand:  "powershell -command Get-Clipboard",
		pasteCommand: "powershell -command Set-Clipboard",
	},
	operatingSystemXorgLinux: {
		copyCommand:  "xclip -selection c",
		pasteCommand: "xclip -selection c -o",
	},
	operatingSystemWaylandLinux: {
		copyCommand:  "wl-copy",
		pasteCommand: "wl-paste",
	},
}

func determineOperatingSystemID(operatingSystem string) operatingSystemID {
	isWSL := func() bool {
		kernelInfo, err := os.ReadFile("/proc/sys/kernel/osrelease")
		if err != nil {
			return false
		}
		return strings.Contains(string(kernelInfo), "microsoft")
	}

	whichLinux := func() operatingSystemID {
		if isWSL() {
			return operatingSystemWindows
		}
		// Check for wayland display.
		if _, isWayland := os.LookupEnv("WAYLAND_DISPLAY"); isWayland {
			return operatingSystemWaylandLinux
		}
		// Check for xorg display.
		if _, isXorg := os.LookupEnv("DISPLAY"); isXorg {
			return operatingSystemWaylandLinux
		}
		return operatingSystemUnknown
	}

	switch operatingSystem {
	case "windows":
		return operatingSystemWindows
	case "darwin":
		return operatingSystemDarwin
	case "linux":
		return whichLinux()
	default:
		return operatingSystemUnknown
	}
}
