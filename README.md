# Intcode Discord Bot

Discord bot that allows you to run Intcode in Discord!

Intcode an assembly-like language defined in the problems of [Advent of Code 2019](https://adventofcode.com/2019).


## Invite this bot to you discord

Invite the bot to your server clicking [here](https://discordapp.com/api/oauth2/authorize?client_id=653389730813313031&permissions=124992&scope=bot)

This is the easiest way to start using the bot.

## Bot commands

Available commands:

- **help** - shows an help message with the commands
- **run <intcode_program> [inputs]** - runs the intcode program and returns the outputs produced by it. Optionally, a list of inputs can be passed to the program. Both the intcode program and the input list should be a comma-separated list of integers.

## Examples

These are some examples of how to use the bot:

* **Echo** - program that outputs its input

```
<Discord user>: !intcode run 3,0,4,0,99 100
<Intcode VM Bot>: Outputs: [100]
```

* **Quine** - program that outputs its source

```
<Discord user>: !intcode run 109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99
<Intcode VM Bot>: Outputs: [109 1 204 -1 1001 100 1 100 1008 100 16 101 1006 101 0 99]
```


* **Print** - prints the number in the middle

```
<Discord user>: !intcode run 104,1125899906842624,99
<Intcode VM Bot>: Outputs: [1125899906842624]
```

**Note:** Make sure your programs terminate with an halt instruction (opcode 99). The bot won't put one for you, and this may cause the program to give you an "unknown opcode" error or it may run into one of the limitations explained below. 

## Limitations

Discord has a limit of around 2000 characters for any given message. This limits the size of the program you can send and the length of the replies the bot can give you. Currently i'm looking for a good way to get around this.

Additionally, to make sure your program doesn't run forever or takes up huge amounts of memory, the vm running a particular program will shutdown if it grows larger than 1 million of intcode memory positions or after it executes 1 million of intcode instructions.

## Bugs and suggestions

Feel free to report bugs and make suggestions to improve the bot in the [issues section](https://github.com/andrerfcsantos/intcode-discord-bot/issues) of this repository.