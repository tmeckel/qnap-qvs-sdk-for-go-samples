package clients

import (
	"github.com/qnap/qvs-sdk-for-go-samples/internal/config"
	"github.com/qnap/qvs-sdk-for-go/services/vms"
)

func NewVirtualMachinesClient() (*vms.VirtualMachinesClient, error) {
	baseURI := config.TenantID() + vms.DefaultBaseURI

	cl := vms.NewVirtualMachinesClientWithBaseURI(baseURI)
	err := configureClient(&cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}

func NewDisksClient() (*vms.DisksClient, error) {
	baseURI := config.TenantID() + vms.DefaultBaseURI

	cl := vms.NewDisksClientWithBaseURI(baseURI)
	err := configureClient(&cl.Client, baseURI)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}
