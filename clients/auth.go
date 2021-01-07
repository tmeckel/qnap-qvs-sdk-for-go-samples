package clients

import (
	"github.com/qnap/qvs-sdk-for-go-samples/utils"
	"github.com/qnap/qvs-sdk-for-go/services/auth"
)

func NewAuthClient(baseURI string) (*auth.Client, error) {
	cl := auth.NewClientWithBaseURI(baseURI)
	err := utils.ConfigureClient(cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}
