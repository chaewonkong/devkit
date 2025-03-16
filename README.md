# 🧰 DevKit

**DevKit** is a Go-based CLI tool that provides developers with a collection of useful utilities to streamline common tasks. It helps developers perform various operations quickly and efficiently, boosting productivity.

## Features

### Ready

- **UUID Generator**: Generate unique UUIDs.
- **URL Encoder/Decoder**: Encode or decode URLs.
- **Base64 Encoder/Decoder**: Encode or decode strings in Base64 format.


### Coming Soon

- **CSV to JSON**: Convert CSV files to JSON format.
- **YAML to JSON**: Convert YAML files to JSON format.
- **JSON to YAML**: Convert JSON files to YAML format.
- **Cron Expression Analyzer**: Analyze and validate cron expressions.
- **Lorem Ipsum Generator**: Generate random Lorem Ipsum text for testing and placeholder purposes.


## Installation
Currently Devkit is only installable with Go.

```bash
$ go install github.com/chaewonkong/devkit/cmd/devkit@v0.1.1
```

## Basic Usage

```bash
# Display help
$ devkit --help

# Generate a UUID
$ devkit uuid
966716aa-afac-4ac4-8805-5086d899509c

# Encode a URL
$ devkit url --encode "http://example.com?name=John Doe&age=30"
http%3A%2F%2Fexample.com%3Fname%3DJohn+Doe%26age%3D30

# Decode a URL
$ devkit url --decode "http%3A%2F%2Fexample.com%3Fname%3DJohn+Doe%26age%3D30"
http://example.com?name=John Doe&age=30
```

Other commands can be used in a similar manner. For more information, refer to the help documentation.

## License

- MIT