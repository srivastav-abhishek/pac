package platform

import (
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/PDeXchange/pac/internal/pkg/client/iam"
)

// NewIAMIdentityClient creates iam identity client.
func NewIAMIdentityClient() (*iamidentityv1.IamIdentityV1, error) {
	auth, err := iam.GetIAMAuth()
	if err != nil {
		return nil, err
	}
	return iamidentityv1.NewIamIdentityV1(&iamidentityv1.IamIdentityV1Options{
		Authenticator: auth,
		URL:           iamidentityv1.DefaultServiceURL,
	})
}
