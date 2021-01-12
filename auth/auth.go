package main

import (
	"context"
	"fmt"
	"os"

	"github.com/qnap/qvs-sdk-for-go-samples/internal/clients"
	"github.com/qnap/qvs-sdk-for-go-samples/internal/config"
)

func main() {

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

}
