package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/qnap/qvs-sdk-for-go-samples/internal/clients"
	"github.com/qnap/qvs-sdk-for-go-samples/internal/config"
	"github.com/qnap/qvs-sdk-for-go/services/users"
)

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

	ucl, err := clients.NewUsersClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create users client. Error: %v\n", err)
		return
	}

	userList, err := ucl.List(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user list. Error: %v\n", err)
		return
	}

	if *userList.Status != 0 {
		fmt.Fprintf(os.Stderr, "Failed to get user list. Status: %d\n", *userList.Status)
		return
	}

	for _, user := range *userList.Data {
		fmt.Fprintf(os.Stdout, "User Name: %s\n", *user.Username)
		if strings.EqualFold("remote_manager", *user.Username) && !*user.IsSuperuser {
			fmt.Fprintf(os.Stdout, "Promoting user %s to superuser\n", *user.Username)
			bTrue := true
			updateState, err := ucl.Update(ctx, *user.ID, &users.WriteUser{
				IsSuperuser: &bTrue,
			})

			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to promote user to superuser. Error: %v\n", err)
			}

			if *updateState.Status != 0 {
				fmt.Fprintf(os.Stderr, "Failed to promote user to superuser. Status: %v\n", *&updateState.Status)
			} else {
				fmt.Fprintf(os.Stdout, "Successfully promoted user to superuser (%v).\n", *updateState.Data.IsSuperuser)
			}

		}

	}

	usr := (*userList.Data)[0]
	permResp, err := ucl.GetPermissions(ctx, *usr.ID, 6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get VM permissions. Error: %v\n", err)
	} else {
		if *permResp.Status != 0 {
			fmt.Fprintf(os.Stderr, "Failed to get VM permissions. Status: %v\n", *permResp.Status)
		} else {
			fmt.Fprintf(os.Stdout, "VM permissions: %v\n", *permResp.Data.Permissions)
		}
	}

}
