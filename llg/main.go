package main

import (
    "fmt"
	"context"
    "os"
	"os/exec"
	"time"
)

func main() {
    timeout := 20*time.Second
    cmdName := "ls"
    cmdArgs := []string{"-l", "-R", "."}
    
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    cmd := exec.CommandContext(ctx, cmdName, cmdArgs...)

    out, err := cmd.Output()

    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println(err)
		os.Exit(1)
    }

    if _, ok := err.(*exec.ExitError); ok {
        fmt.Println(err)
		os.Exit(1)
    }

    fmt.Println(string(out))
}

