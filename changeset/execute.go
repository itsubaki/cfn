package changeset

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cli "gopkg.in/urfave/cli.v1"
)

func Execute(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(change-set name) is required")
		os.Exit(1)
	}

	changeSetName := c.Args().Get(0)
	tmp := strings.Replace(changeSetName, "changeset-", "", -1)
	stackName := tmp[:strings.LastIndex(tmp, "-")]
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
