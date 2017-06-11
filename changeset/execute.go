package changeset

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cli "gopkg.in/urfave/cli.v1"
)

func Execute(c *cli.Context) {
	if len(c.Args()) < 2 {
		fmt.Println("error: first argument(stack name) is required")
		fmt.Println("error: second argument(change-set name) is required")
		os.Exit(1)
	}

	stackName := c.Args().Get(0)
	changeSetName := c.Args().Get(1)
	fmt.Print(stackName)

	req := &cf.ExecuteChangeSetInput{
		StackName:     &stackName,
		ChangeSetName: &changeSetName,
	}

	client := cf.New(session.Must(session.NewSession()))
	_, err := client.ExecuteChangeSet(req)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		return
	}

	fmt.Println(" updated. " + changeSetName)
}
