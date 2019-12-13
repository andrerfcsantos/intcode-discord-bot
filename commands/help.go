package commands

import "fmt"

// Help performs the actions for the "help" command sent to the bot,
// which informs the user about the usage and commands available
func Help(prefix string) string {
	return fmt.Sprintf("usage: %s [command] [command_args...]\n", prefix) +
		`
Intcode is an assembly-like language defined in Advent of Code 2019. This bot runs Intcode programs and returns its outputs. 
Available commands:
    - **help** - shows this message
	- **run <intcode_program> [inputs]** - runs the intcode program and returns the outputs produced by it. Optionally, a list of inputs can be passed to the program. Both the intcode program and the input list should be a comma-separated list of integers.
`
}
