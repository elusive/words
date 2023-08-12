# Words 
Simple command line utility to count STDIN by line, word, or byte.

## Overview
This is a small application built while reading the book "Powerful Command Line Application" by Ricardo Gerardi. 
Which is available from the [Pragmatic Bookshelf](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-god).

## Usage
To use this utility you can build it locally if you have Go installed.  If you do not have Go installed you can follow the 
instructions for your operating system at (https://go.dev/doc/install) and then return here.

1. Clone the repository onto your local drive (or download the zip):
   ```bash
   git clone git@github.com:elusive/words.git
   ```
2. Go to the cloned repository folder and build the application:
   ```bash
   cd words && go build
   ```

Once you build the application you should see the binary in the repository folder.  Either "words.exe" or "words". The 
`go build` command will automatically build the application for your specific OS.  If you wish to build cross-platform
for an OS and Architecture other than your own you can do so using these commands:
```bash
env GOOS=windows GOARCH=amd64 go build
env GOOS=darwin GOARCH=arm64 go build
env GOOS=linux GOARCH=386 go build
```

### CLI Tests
The built application can be used to count files or command output, etc.
```bash
echo -e "Test string\n With multiple words\n and lines\n" | ./words.exe      # 7 words
echo -e "Test string\n With multiple words\n and lines\n" | ./words.exe -l   # 4 lines
echo -e "Test string\n With multiple words\n and lines\n" | ./words.exe -l   # 45 bytes
```

### Unit Tests
You can execute the unit tests once you have cloned the repository also, using this command:
```bash
go test -v
```

The current unit tests should result in the following output to the terminal:
```bash
=== RUN   TestCountWords
--- PASS: TestCountWords (0.00s)
=== RUN   TestCountLines
--- PASS: TestCountLines (0.00s)
=== RUN   TestCountBytes
--- PASS: TestCountBytes (0.00s)
=== RUN   TestLines
--- PASS: TestLines (0.00s)
PASS
ok      github.com/elusive/words        0.056s
```

## Further Questions
For any further questions please email:  `me at johng dot info`
