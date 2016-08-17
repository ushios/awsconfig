package awsconfig

import (
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	// CredentialFilePath is path to credentials file
	// See also: http://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-getting-started.html
	CredentialFilePath string
	// CredentialName is name of credentials file setting
	CredentialName string
	// Timeout aws-sdk session timeout
	Timeout time.Duration
	// Region is region of AWS
	Region string
)

func init() {
	if CredentialFilePath == "" {
		CredentialFilePath = os.Getenv("HOME") + "/.aws/credentials"
	}

	if CredentialName == "" {
		CredentialName = "default"
	}

	if Timeout == 0 {
		Timeout = time.Duration(5 * time.Second)
	}

	if Region == "" {
		Region = "us-east-1"
	}
}

// Credentials return credential info
func Credentials() *credentials.Credentials {
	var c *credentials.Credentials

	if fileExists(CredentialFilePath) {
		c = credentials.NewSharedCredentials(CredentialFilePath, CredentialName)
	} else {
		cl := ec2metadata.New(session.New(), &aws.Config{
			HTTPClient: &http.Client{Timeout: Timeout},
		})
		c = credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{
			Client: cl,
		})
	}

	return c
}

// Config return aws config
func Config() *aws.Config {
	return &aws.Config{
		Region:      aws.String(Region),
		Credentials: Credentials(),
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	if pathError, ok := err.(*os.PathError); ok {
		if pathError.Err == syscall.ENOTDIR {
			return false
		}
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}
