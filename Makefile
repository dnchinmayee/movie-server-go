.PHONY: all build clean

# Default target
all: build

# Build the binary
build:
	/home/chikki/apps/go/bin/go build -o movie-server main.go

# Clean up build artifacts
clean:
	rm -f movie-server
