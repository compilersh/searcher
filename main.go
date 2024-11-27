package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux"
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-cli <storeName> <searchQuery>")
		return
	}

	storeName := os.Args[1]
	searchQuery := os.Args[2]

	switch storeName {
	case "jula":
		baseURL := "https://jula.se/search?query="
		fullURL := baseURL + url.QueryEscape(searchQuery)
		if err := openBrowser(fullURL); err != nil {
			fmt.Printf("Failed to open browser: %v\n", err)
		}
	case "biltema":
		baseURL := "https://www.biltema.se/soksida/?query="
		fullURL := baseURL + url.QueryEscape(searchQuery)
		if err := openBrowser(fullURL); err != nil {
			fmt.Printf("Failed to open browser: %v\n", err)
		}

	default:
		fmt.Printf("Store '%s' is not supported.\n", storeName)
	}
}
