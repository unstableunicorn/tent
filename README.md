[![Go](https://github.com/unstableunicorn/tent/actions/workflows/go.yml/badge.svg)](https://github.com/unstableunicorn/tent/actions/workflows/go.yml)

# Tent
Tent is a cli utility for printing the Nova Console Tent(like consultant...well i find it funny...nvm) but prints a tent in your console with the name of searched winner.
This is a totally very silly pointless project for a bit of fun.

## Installation
You can download from the releases: [latest releases](https://github.com/unstableunicorn/tent/releases/latest) or curl your version and architecture:
Available architectures are `linux|darwin|windows`:
```
# Basic:
curl -L https://github.com/unstableunicorn/tent/releases/download/<version>/tent-<architecture>-amd64 -o <binary output path>

# Windows
curl -L https://github.com/unstableunicorn/tent/releases/download/v0.1.3/tent-windows-amd64 -o tent.exe

# Darwin
curl -L https://github.com/unstableunicorn/tent/releases/download/v0.1.3/tent-darwin-amd64 -o tent

# Linux
curl -L https://github.com/unstableunicorn/tent/releases/download/v0.1.3/tent-linux-amd64 -o tent
```
Copy to your PATH and run it anywhere!

Alternatively you can use the Docker images, cause why not:
```
docker run --rm unstableunicorn/tent
```
Or install as a go module like this:
```
go install github.com/unstableunicorn/tent
```

Without any arguments it will print the help and exit.

## Usage
Just play around a bit with help, will only list the base help here:
```
>tent
So you want to know who the console tent was for a particular year?
Perhaps you want to know which year someone was a console tent?
Or you just want to print a tent in your console?
Then look no further!
Just type the year like so:
 tent 2021
or their first name:
 tent Phil
If you want a list, well I can give you a list:
 tent list
The rest of this amazing application you can explore by yourself.

Usage:
  tent [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  grab        Grab the person
  help        Help about any command
  justthetent Just show me the tent
  mumble      If said softly you might find what you need
  yell        Show me all the console tents

Flags:
  -h, --help      help for tent
  -t, --inatent   if you want to more tent, this is how you get more tent

Use "tent [command] --help" for more information about a command.
```

Lets leave it there for now
