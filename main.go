package main

import (
	"os"

	"github.com/itsubaki/cfn/stack"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "cfn"
	app.Usage = ""
	app.Version = "0.0.1"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "credential",
			Value: "credential.yaml",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.yaml",
		},
	}

	validate := cli.Command{
		Name:    "validate",
		Aliases: []string{"v"},
		Usage:   "Validates a specified template.",
		Flags:   flags,
		Action:  stack.Validate,
	}

	stack := cli.Command{
		Name:    "stack",
		Aliases: []string{"s"},
		Usage:   "Create, Update, Delete, Describe Stack.",
		Subcommands: []cli.Command{
			{
				Name:   "create",
				Action: stack.Create,
				Usage:  "Creates a stack as specified in the template.",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:   "update",
				Action: stack.Update,
				Usage:  "Updates a stack as specified in the template.",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:   "delete",
				Action: stack.Delete,
				Usage:  "Deletes a specified stack.",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:   "describe",
				Action: stack.Update,
				Usage:  "Returns the description for the specified stack.",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
		},
	}

	app.Commands = []cli.Command{
		validate,
		stack,
	}

	app.Run(os.Args)
}
