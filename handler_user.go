package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>\n", cmd.Name)
	}

	name := cmd.Args[0]

	err := s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w\n", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
