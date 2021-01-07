package clients

import (
	"github.com/qnap/qvs-sdk-for-go-samples/utils"
	"github.com/qnap/qvs-sdk-for-go/services/users"
)

func NewUsersClient(baseURI string) (*users.Client, error) {
	cl := users.NewClientWithBaseURI(baseURI)
	err := utils.ConfigureClient(cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}
