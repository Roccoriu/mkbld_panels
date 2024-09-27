# Makefile for building the Go program for multiple platforms

# Replace 'your_program_name' with the desired name for your executable
BINARY_NAME = mkbhd_downloader

all: windows_amd64 linux_amd64 macos_arm64

windows_amd64:
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)_windows_amd64.exe
	GOOS=windows GOARCH=arm64 go build -o dist/$(BINARY_NAME)_windows_arm64.exe

linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o dist/$(BINARY_NAME)_linux_arm64

macos_arm64:
	GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME)_macos_arm64

