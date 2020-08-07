# BeeClip 🐝📋

![](https://user-images.githubusercontent.com/13544676/89609229-19b9f780-d82c-11ea-8dca-349f6e07b1f2.png)

`BeeClip` is small abstraction layer over your operating system's clipboard CLI.

## Supported Operating Systems

| Operating System                     |              Command |
| ------------------------------------ | -------------------: |
| Windows / Windows Subsytem for Linux |           `clip.exe` |
| Darwin                               |             `pbcopy` |
| X11 Linux                            | `xclip -selection c` |
| Wayland Linux                        |            `wl-copy` |

## Installation

### Using GitHub Releases

Download the appropriate version of `BeeClip` for your operating system and architecture on the [release page](https://github.com/penguingovernor/beeclip/releases).

Don't forget to add `BeeClip` to your operating system's PATH variable.

### Using Go

If you have the `go` binary installed on your system, you can get `BeeClip` by simply running `go install github.com/penguingovernor/beeclip`.

## Usage

To use `BeeClip` pipe the output of your command to it.

Example:

```bash
# Pipe the output of a command.
$ echo "Hello world!" | beeclip

# Copy contents of a file to clipboard.
$ beeclip < file.txt

# Or for multiple files.
$ cat file1.txt file2.txt | beeclip
```

And this wouldn't be a complete `README.md` without a GIF example:

![](https://user-images.githubusercontent.com/13544676/89611200-6ce27900-d831-11ea-920a-d7040564edb6.gif)
