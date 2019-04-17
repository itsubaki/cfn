package changeset

import (
	"fmt"
	"os"
	"strings"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	ses "github.com/itsubaki/cfn/session"
	cli "gopkg.in/urfave/cli.v1"
)

func Describe(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(change-set) is required")
		os.Exit(1)
	}

	changeSetName := c.Args().Get(0)
	tmp := strings.Replace(changeSetName, "changeset-", "", -1)
	stackName := tmp[:strings.LastIndex(tmp, "-")]

	req := &cf.DescribeChangeSetInput{
		StackName:     &stackName,
		ChangeSetName: &changeSetName,
	}

	client := cf.New(ses.New())
	out, err := client.DescribeChangeSet(req)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		return
	}

	fmt.Println(out)
}
