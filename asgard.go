package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Odin struct {
	KeepShellConnection bool
	ExecuteCommand      string
	Jitter              int64
}

// askOdin fetches an Odin struct from the remote C2 HTTP endpoint.
func askOdin(url string) Odin {
	resp, err := http.Get(url)
	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Failed to get Odin config from %s: %v\n", url, err)
		return Odin{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read body: %v\n", err)
		return Odin{}
	}

	var odin Odin
	if err := json.Unmarshal(body, &odin); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode Odin config: %v\n", err)
		return Odin{}
	}
	return odin
}

func main() {
	// CLI flags
	useRemote := flag.Bool("remote", false, "Fetch configuration from remote Odin server (overrides other options)")
	serverURL := flag.String("url", "http://localhost:8080/serve", "Remote Odin server URL")
	keepShell := flag.Bool("keep-shell", false, "Keep the shell connection alive")
	cmd := flag.String("cmd", "", "Command to execute")
	jitter := flag.Int64("jitter", 0, "Jitter value in milliseconds")

	flag.Parse()

	var config Odin

	if *useRemote {
		config = askOdin(*serverURL)
		fmt.Println("Fetched Odin config from remote server:")
	} else {
		config = Odin{
			KeepShellConnection: *keepShell,
			ExecuteCommand:      *cmd,
			Jitter:              *jitter,
		}
		fmt.Println("Configuration from CLI flags:")
	}

	// Pretty print the Odin config (for demonstration)
	out, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(out))

	// Implement your logic here using config
	// For example: if config.ExecuteCommand != "" { ... }
}
