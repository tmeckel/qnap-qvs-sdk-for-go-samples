package clients

import (
	"github.com/qnap/qvs-sdk-for-go-samples/internal/config"
	"github.com/qnap/qvs-sdk-for-go/services/auth"
)

func NewAuthClient() (*auth.Client, error) {
	baseURI := config.TenantID() + auth.DefaultBaseURI

	cl := auth.NewClientWithBaseURI(baseURI)
	err := configureClient(&cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}
