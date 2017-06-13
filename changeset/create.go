package changeset

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	cli "gopkg.in/urfave/cli.v1"
)

func Create(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group name) is required")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("config"))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cf.New(session.Must(session.NewSession()))
	for _, tmplpath := range config.Template() {
		fmt.Print(tmplpath)

		name := cfg.StackName(c.Args().Get(0), tmplpath)
		body, err := cfg.TemplateBody(tmplpath)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		changeSetName := "changeset-" + name + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		iam := "CAPABILITY_IAM"
		niam := "CAPABILITY_NAMED_IAM"
		req := &cf.CreateChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
			TemplateBody:  &body,
			Capabilities:  []*string{&iam, &niam},
			Tags:          config.Tag(),
		}

		res, err := client.CreateChangeSet(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		desc := &cf.DescribeChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
		}
		err = client.WaitUntilChangeSetCreateComplete(desc)
		if err == nil {
			fmt.Println(" created. " + *res.Id)
			continue
		}

		// Delete ChangeSet of Failed Status
		input := &cf.DeleteChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
		}
		_, err = client.DeleteChangeSet(input)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
		}

		fmt.Println(" no update.")

	}
}
