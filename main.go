package main

import (
	"encoding/json" // For reading/parsing JSON configuration
	"fmt"           // For formatted input/output
	"net/http"      // For HTTP server functionality
	"os"            // For file system operations and command-line arguments

	"github.com/akamensky/argparse" // For parsing command-line arguments
)

// Options encapsulates configuration options for the reverse proxy.
type Options struct {
	Location       map[string]string // Key-value mapping of routes to target locations
	Threads        int               // Number of threads for handling requests
	CustomHeaders  bool              // Allow custom headers in requests
	AllowAnyOrigin bool              // Allow requests from any origin
	AllowMethods   []string          // List of allowed HTTP methods
}

// CreateParsers creates and returns a command-line argument parser.
func CreateParsers() string {
	parser := argparse.NewParser("amazing-reverse-proxy", "Amazing reverse proxy, crafted by Edgar")
	path := parser.String("p", "path", &argparse.Options{Required: false, Help: "Path to the configuration file", Default: "./config.json"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	return *path
}

// readConfigFile reads and parses configuration options from a JSON file.
func readConfigFile(pathFile string) Options {
	var config Options

	data, err := os.ReadFile(pathFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing configuration file:", err)
		os.Exit(1)
	}

	return config
}

func main() {
	pathFile := CreateParsers()
	options := readConfigFile(pathFile)
	CreateServer(options)
}

// CreateServer initializes and starts an HTTP server based on provided options.
func CreateServer(options Options) {
	for route, target := range options.Location {
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			// Implement actual request handling logic here, instead of "Hello World"
			fmt.Println("Forwarding request to", target) // Example placeholder
			// Forward request to target location
		})
	}

	port := ":8080"
	fmt.Printf("Server started at address http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
