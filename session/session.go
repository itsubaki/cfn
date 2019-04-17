package session

import (
	"github.com/aws/aws-sdk-go/aws"
	ses "github.com/aws/aws-sdk-go/aws/session"
)

func New(region ...string) *ses.Session {
	opts := ses.Options{SharedConfigState: ses.SharedConfigEnable}
	if len(region) > 0 && region[0] != "" {
		opts.Config.Region = aws.String(region[0])
	}
	return ses.Must(ses.NewSessionWithOptions(opts))
}
