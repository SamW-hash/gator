package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	cmdlist map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmdlist[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	if err := handler(s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdlist[name] = f
}
