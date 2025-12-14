# QUT Language Interpreter

A simple interpreter for **QUT**, a minimalistic esoteric programming language implemented in Go.

## Features

* Execute `.qut` files or run commands interactively via the command line.
* Supports basic tape operations, input/output, and register manipulation.
* Handles loops with automatic jump table creation.
* Includes optional debug output via `DEBUG=true` environment variable.

## Instructions

The language uses 12 instructions:

| Instruction      | Code  | Description                               |
| ---------------- | ----- | ----------------------------------------- |
| Loop jump        | `qut` | Jump to matching `QUT`                    |
| Move left        | `qUt` | Move tape pointer left                    |
| Move right       | `quT` | Move tape pointer right                   |
| Execute          | `qUT` | Execute current tape value as instruction |
| Output/Input     | `Qut` | Output character or read input            |
| Decrement        | `QUt` | Decrement current cell                    |
| Increment        | `QuT` | Increment current cell                    |
| Conditional jump | `QUT` | Jump if current cell is zero              |
| Reset cell       | `UUU` | Set current cell to 0                     |
| Register swap    | `QQQ` | Swap with register                        |
| Print char       | `TUQ` | Print character from current cell         |
| Read char        | `Tuq` | Read character into current cell          |

## Usage

Run a `.qut` file:

```bash
go run main.go program.qut
```

Start interactive mode:

```bash
go run main.go
```

Set debug mode:

```bash
DEBUG=true go run main.go program.qut
```

## Requirements

* Go 1.20+
* Works on Linux, macOS, and Windows
