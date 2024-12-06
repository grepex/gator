package main

import "errors"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	_, exists := c.commands[cmd.name]
	if !exists {
		return errors.New("command not found")
	}

	f := c.commands[cmd.name]
	err := f(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

