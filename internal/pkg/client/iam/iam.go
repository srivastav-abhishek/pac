package iam

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"os"
)

// GetIAMAuth returns the IAM authenticator
func GetIAMAuth() (*core.IamAuthenticator, error) {
	key := os.Getenv("IBMCLOUD_APIKEY")
	if key == "" {
		return nil, fmt.Errorf("empty IBMCLOUD_APIKEY, set the IBMCLOUD_APIKEY environment variable")
	}
	return &core.IamAuthenticator{
		ApiKey: key,
	}, nil
}
