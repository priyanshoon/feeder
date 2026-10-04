package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/priyanshoon/feeder/internal/config"
	"github.com/priyanshoon/feeder/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("database error", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	programState := &state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerResetUser)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerFeeds)

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
