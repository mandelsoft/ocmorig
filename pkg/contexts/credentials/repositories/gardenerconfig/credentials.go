package gardenerconfig

import (
	"github.com/open-component-model/ocm/pkg/contexts/credentials/cpi"
)

type credentialGetter struct {
	getCredentials func() (cpi.Credentials, error)
}

var _ cpi.CredentialsSource = credentialGetter{}

func (c credentialGetter) Credentials(ctx cpi.Context, cs ...cpi.CredentialsSource) (cpi.Credentials, error) {
	return c.getCredentials()
}
