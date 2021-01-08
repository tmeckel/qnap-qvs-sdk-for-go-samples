package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/qnap/qvs-sdk-for-go-samples/internal/clients"
	"github.com/qnap/qvs-sdk-for-go-samples/internal/config"
)

func printDiskInfo(ctx context.Context) {
	dskcl, err := clients.NewDisksClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create disks client. Error: %v\n", err)
	}
	dsklist, err := dskcl.List(ctx, 6)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve list of disks attached to vm")
	} else {
		for _, dsk := range *dsklist.Data {
			fmt.Fprintf(os.Stdout, "Disk RootPath: %s\n", *dsk.RootPath)
		}
	}
}

func main() {

	autorest.SenderFactoryInstance = clients.SenderFactory
	config.ParseEnvironment()

	cl, err := clients.NewAuthClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create auth client. Error: %v\n", err)
		return
	}
	ctx := context.Background()

	lstat, err := cl.Login(ctx, config.ClientID(), config.ClientSecret())
	if nil != err {
		fmt.Fprintf(os.Stderr, "Failed to login. Error: %v\n", err)
		return
	}
	if *lstat.Status != 0 {
		fmt.Fprintf(os.Stderr, "Failed to login. Status: %d\n", *lstat.Status)
		return
	}

	defer func() {
		_, err = cl.Logout(ctx)
		if nil != err {
			fmt.Fprintf(os.Stderr, "Failed to logout %v\n", err)
		}
	}()

	vmcl, err := clients.NewVirtualMachinesClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create virtual machines client. Error: %v\n", err)
		return
	}

	//vmlist, err := vmcl.List(ctx, false, nil, nil, nil)
	vmlist, err := vmcl.ListStates(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to retrieve list of VMs")
	} else {
		for _, vm := range *vmlist.Data {
			fmt.Fprintf(os.Stdout, "VM Name: %s\n", *vm.Name)
		}
	}
	for _, vm := range *vmlist.Data {
		if strings.EqualFold("win10", *vm.Name) {
			if strings.EqualFold("running", *vm.PowerState) {
				fmt.Fprintf(os.Stdout, "Stopping VM Name: %s\n", *vm.Name)
				// bug in QVS => to gracefully shutdown a VM it must be triggered multiple times
				bIsStopped := false
				for j := 0; j < 10 && !bIsStopped; j++ {
					_, err := vmcl.Shutdown(ctx, *vm.ID)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Failed to shutdown VMs")
						break
					}
					fmt.Fprintf(os.Stdout, "Successfully initiated shutdown for VM ... Wating for VM changing state\n")
					time.Sleep(5 * time.Second)

					for i := 0; i < 10 && !bIsStopped; i++ {
						//stats, err := vmcl.StoppingProgress(ctx)
						stats, err := vmcl.ListStates(ctx)
						if err != nil {
							fmt.Fprintln(os.Stderr, "Failed to get shutdown progress")
							break
						}
						//for _, vmStat := range *stats.Data.VmsProperty {
						for _, vmStat := range *stats.Data {
							if strings.EqualFold("win10", *vmStat.Name) {
								if !strings.EqualFold("running", *vmStat.PowerState) {
									bIsStopped = true
									break
								}
							}
						}

						if !bIsStopped {
							fmt.Fprintf(os.Stdout, "VM has not reached stopped state ...\n")
							time.Sleep(5 * time.Second)
						}
					}
				}
				if !bIsStopped {
					fmt.Fprintf(os.Stderr, "Failed to stop VM\n")
				} else {
					fmt.Fprintf(os.Stdout, "Successfully stopped VM\n")
				}
			} else {
				fmt.Fprintf(os.Stdout, "Starting VM Name: %s\n", *vm.Name)
				sresp, err := vmcl.Start(ctx, *vm.ID)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Failed to start VMs")
				} else {
					fmt.Fprintf(os.Stdout, "Successfully started VM: status %d\n", *sresp.Status)
				}
			}
		}
	}
	printDiskInfo(ctx)
}
