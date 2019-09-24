// Package cli is a command line interface
package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/micro/cli"
)

var (
	prompt = "micro> "

	commands = map[string]*command{
		"quit":       &command{"quit", "Exit the CLI", quit},
		"exit":       &command{"exit", "Exit the CLI", quit},
		"call":       &command{"call", "Call a service", callService},
		"list":       &command{"list", "List services, peers or routes", list},
		"get":        &command{"get", "Get service info", getService},
		"stream":     &command{"stream", "Stream a call to a service", streamService},
		"publish":    &command{"publish", "Publish a message to a topic", publish},
		"health":     &command{"health", "Get service health", queryHealth},
		"stats":      &command{"stats", "Get service stats", queryStats},
		"register":   &command{"register", "Register a service", registerService},
		"deregister": &command{"deregister", "Deregister a service", deregisterService},
	}
)

type command struct {
	name  string
	usage string
	exec  exec
}

func runc(c *cli.Context) {
	commands["help"] = &command{"help", "CLI usage", help}
	alias := map[string]string{
		"?":  "help",
		"ls": "list",
	}

	r, err := readline.New(prompt)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}
	defer r.Close()

	for {
		args, err := r.Readline()
		if err != nil {
			fmt.Fprint(os.Stdout, err)
			return
		}

		args = strings.TrimSpace(args)

		// skip no args
		if len(args) == 0 {
			continue
		}

		parts := strings.Split(args, " ")
		if len(parts) == 0 {
			continue
		}

		name := parts[0]

		// get alias
		if n, ok := alias[name]; ok {
			name = n
		}

		// break out on quit
		switch name {
		case "quit", "exit":
			return
		}

		if cmd, ok := commands[name]; ok {
			rsp, err := cmd.exec(c, parts[1:])
			if err != nil {
				println(err.Error())
				continue
			}
			println(string(rsp))
		} else {
			println("unknown command")
		}
	}
}

func Init(ctx *cli.Context) {
	runc(ctx)
}

func HealthCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "check",
			Usage:  "Query the health of a service",
			Action: printer(queryHealth),
		},
	}
}

func NetworkCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "graph",
			Usage:  "Get the network graph",
			Action: printer(networkGraph),
		},
		{
			Name:   "nodes",
			Usage:  "List nodes in the network",
			Action: printer(listNodes),
		},
		{
			Name:   "routes",
			Usage:  "List network routes",
			Action: printer(listRoutes),
		},
	}
}

func RegistryCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list",
			Usage: "List items in registry or network",
			Subcommands: []cli.Command{
				{
					Name:   "nodes",
					Usage:  "List nodes in the network",
					Action: printer(listNodes),
				},
				{
					Name:   "routes",
					Usage:  "List network routes",
					Action: printer(listRoutes),
				},
				{
					Name:   "services",
					Usage:  "List services in registry",
					Action: printer(listServices),
				},
			},
		},
		{
			Name:  "register",
			Usage: "Register an item in the registry",
			Subcommands: []cli.Command{
				{
					Name:   "service",
					Usage:  "Register a service with JSON definition",
					Action: printer(registerService),
				},
			},
		},
		{
			Name:  "deregister",
			Usage: "Deregister an item in the registry",
			Subcommands: []cli.Command{
				{
					Name:   "service",
					Usage:  "Deregister a service with JSON definition",
					Action: printer(deregisterService),
				},
			},
		},
		{
			Name:  "get",
			Usage: "Get item from registry",
			Subcommands: []cli.Command{
				{
					Name:   "service",
					Usage:  "Get service from registry",
					Action: printer(getService),
				},
			},
		},
	}
}

func Commands() []cli.Command {
	commands := []cli.Command{
		{
			Name:   "cli",
			Usage:  "Run the interactive CLI",
			Action: runc,
		},
		{
			Name:   "call",
			Usage:  "Call a service",
			Action: printer(callService),
		},
		{
			Name:   "stream",
			Usage:  "Create a service stream",
			Action: printer(streamService),
		},
		{
			Name:   "publish",
			Usage:  "Publish a message to a topic",
			Action: printer(publish),
		},
		{
			Name:   "stats",
			Usage:  "Query the stats of a service",
			Action: printer(queryStats),
		},
	}

	return append(commands, RegistryCommands()...)
}
