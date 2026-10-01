package main

import (
	"fmt"
	"log"
	"os"

	"github.com/priyanshoon/feeder/internal/config"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: prog <command> [args...]")
		os.Exit(1)
	}

	argCommand := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(programState, argCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
