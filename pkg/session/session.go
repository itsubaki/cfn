package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func New(region ...string) *session.Session {
	opts := session.Options{SharedConfigState: session.SharedConfigEnable}
	if len(region) > 0 && region[0] != "" {
		opts.Config.Region = aws.String(region[0])
	}

	return session.Must(session.NewSessionWithOptions(opts))
}
