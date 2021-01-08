package clients

import (
	"github.com/qnap/qvs-sdk-for-go-samples/internal/config"
	"github.com/qnap/qvs-sdk-for-go/services/users"
)

func NewUsersClient() (*users.Client, error) {
	baseURI := config.TenantID() + users.DefaultBaseURI

	cl := users.NewClientWithBaseURI(baseURI)
	err := configureClient(&cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}
