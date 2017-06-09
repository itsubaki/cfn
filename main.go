package main

import (
	"os"

	"github.com/itsubaki/cfn/changeset"
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
			Name:  "config, c",
			Value: "cfn.yaml",
		},
	}

	validate := cli.Command{
		Name:    "validate",
		Aliases: []string{"v"},
		Usage:   "Validates a specified template",
		Flags:   flags,
		Action:  stack.Validate,
	}

	stack := cli.Command{
		Name:    "stack",
		Aliases: []string{"s"},
		Usage:   "Create, Update, Delete, Describe Stack",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Action:  stack.Create,
				Aliases: []string{"c"},
				Usage:   "Creates a stack as specified in the template",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:    "update",
				Action:  stack.Update,
				Aliases: []string{"u"},
				Usage:   "Updates a stack as specified in the template",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:    "delete",
				Action:  stack.Delete,
				Aliases: []string{"d"},
				Usage:   "Deletes a specified stack",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:    "describe",
				Action:  stack.Describe,
				Aliases: []string{"desc"},
				Usage:   "Returns the description for the specified stack",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
		},
	}

	changeset := cli.Command{
		Name:    "changeset",
		Aliases: []string{"cs"},
		Usage:   "Create, Execute, Delete, Describe Changeset",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Action:  changeset.Create,
				Aliases: []string{"c"},
				Usage:   "Creates a list of changes that will be applied to a stack so that you can review the changes before executing them",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:    "execute",
				Action:  changeset.Execute,
				Aliases: []string{"e"},
				Usage:   "Updates a stack using the input information that was provided when the specified change set was created",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:    "delete",
				Action:  changeset.Delete,
				Aliases: []string{"d"},
				Usage:   "Deletes the specified change set",
				Flags: append(flags,
					cli.StringFlag{
						Name:  "name, n",
						Usage: "StackName",
					}),
			},
			{
				Name:    "describe",
				Action:  changeset.Describe,
				Aliases: []string{"desc"},
				Usage:   "Returns the inputs for the change set and a list of changes that AWS CloudFormation will make if you execute the change set",
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
		changeset,
	}

	app.Run(os.Args)
}
