package commands

import (
	"fmt"
	"intcode-discord-bot/arguments"
	"intcode-discord-bot/intcode"
)

func Run(args arguments.CommandArguments) (string, error) {
	var err error
	var memory intcode.Memory
	var input intcode.SimpleIntReader
	var output intcode.SimpleIntWriter

	nargs := len(args.Arguments)

	switch nargs {
	case 1:
		memory, err = intcode.ParseMemory(args.Arguments[0])
		if err != nil {
			return "", fmt.Errorf("could not parse memory string: %w", err)
		}
		input = intcode.NewSimpleIntReader()
	case 2:
		memory, err = intcode.ParseMemory(args.Arguments[0])
		if err != nil {
			return "", fmt.Errorf("could not parse memory string: %w", err)
		}

		ins, err := intcode.ParseInputString(args.Arguments[1])
		if err != nil {
			return "", fmt.Errorf("could not parse list of inputs: %w", err)
		}

		input = intcode.NewSimpleIntReader(ins...)
	default:
		return "", fmt.Errorf("wrong number of arguments in command 'run', expected 1 or 2, got %v", nargs)

	}
	output = intcode.NewSimpleIntWriter()

	vm := intcode.VM{
		Memory:           memory,
		Input:            &input,
		Output:           &output,
		MemoryLimit:      1_000_000,
		InstructionLimit: 1_000_000,
	}

	err = vm.Run()
	if err != nil {
		return "", fmt.Errorf("running intcode vm: %w", err)
	}

	return fmt.Sprintf("Outputs: %v", output.Values()), nil
}
