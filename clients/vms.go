package clients

import (
	"github.com/qnap/qvs-sdk-for-go-samples/utils"
	"github.com/qnap/qvs-sdk-for-go/services/vms"
)

func NewVirtualMachinesClient(baseURI string) (*vms.VirtualMachinesClient, error) {
	cl := vms.NewVirtualMachinesClientWithBaseURI(baseURI)
	err := utils.ConfigureClient(cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}

func NewDisksClient(baseURI string) (*vms.DisksClient, error) {
	cl := vms.NewDisksClientWithBaseURI(baseURI)
	err := utils.ConfigureClient(cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}
