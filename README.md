# win-supervisor
> A supervisord-like daemon for Windows

## Download the Latest Release!
TODO

## Installation Instructions
1. Download the latest release
2. Place the executable in `C:\winsupervisor`
3. Create a config file in `C:\winsupervisor` named `winsupervisor.conf`
```
# winsupervisor.conf
[program:my-program]
Path=C:\path\to\myprogram.exe
Exe=myprogram.exe
Name=myprogram
```
Check out the example file in this repo (example.conf)

## Configuration
Configuration is done in `C:\winsupervisor\winsupervisor.conf`

### Program Header
Each program is defined using a program header.  This is defined by the pattern
```
[program:<program-name>]
```
Where `<program-name>` is the name of the program to be defined.

### Program attributes
Each program should contain the following attributes, expressed as `Key=value`. Do not place empty lines before, after, or in between definitions.
* Name `string` - This should represent the name of the program to be supervised
* Path `string` - This should be expressed as the full Windows path using forward slashes
* Exe `string` - This should be the name of the executable to be supervised

## Running Tests
TODO

## Building the Project
1. Clone the project
2. Run `go build`
