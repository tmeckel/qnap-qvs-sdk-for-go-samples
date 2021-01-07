package main

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/qnap/qvs-sdk-for-go-samples/clients"
	"github.com/qnap/qvs-sdk-for-go-samples/utils"
)

func main() {

	autorest.SenderFactoryInstance = utils.SenderFactory

	baseURI := "https://fileserver/qvs"

	cl, err := clients.NewAuthClient(baseURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create auth client. Error: %v\n", err)
		return
	}
	ctx := context.Background()

	lstat, err := cl.Login(ctx, "remote_manager", b64.StdEncoding.EncodeToString([]byte("")))
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

}
