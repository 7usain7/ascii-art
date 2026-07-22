# ASCII-Art Generator

A simple Go application that converts input strings into horizontal ASCII art using font template banners.

## Features

- Renders printable ASCII characters (32 to 127).
- Supports multiline input using literal `\n` character sequences.
- Efficiently parses banner templates into memory for fast lookup.

## Usage

Run the program with a single string argument:

```bash
go run . "Your Text Here"
```

### Example

```bash
go run . "hello"
```

Output:
```text
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
```

To handle newline sequences, use `\n`:

```bash
go run . "hello\nworld"
```

## Banner Templates

The application loads font templates from the `models/` directory. By default, it uses `models/standard.txt`. You can customize the source path in `input_processing.go` to use other available styles like `shadow.txt` or `thinkertoy.txt`.
