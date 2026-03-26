package main

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi"
	"github.com/Oskar-jansson/m5adminapi/models"
)

//nolint:all
func main() {

	// Errors returned from sdk ar in this format.
	// Its to seperate Api reported errors from internal errors.
	var sdkErr *models.SdkError

	credentials := &models.Credentials{
		Id:       "1",
		System:   "RASYSTEM",
		Lang:     "sv",
		User:     "apiOperator",
		Password: "1234",
		Apitype:  "main",
		Apikey:   "1A2B3C4D",
	}

	httpClient := m5adminapi.NewClient().
		SetAddress("http://localhost").
		SetPath("/m5adminapi/api").
		SetCredentials(*credentials)

	api := m5adminapi.NewAdminApiConnection(httpClient)

	// Get api version
	version, sdkErr := api.Version.Get(context.Background())
	if sdkErr != nil {
		// Error
	}
	fmt.Println(version.AsString)

	// Login and recive accesstoken.
	// Can also use Auth.SetAccesstoken() to use an existing accesstoken.
	sdkErr = api.Auth.Login(context.Background())
	if sdkErr != nil {
		// Error
	}

	// Get user with id "2"
	user, sdkErr := api.User.Get(context.Background(), 2)
	if sdkErr != nil {
		// Error
	}
	fmt.Println(user.Firstname)

	// List users where firstname is "john" and lastname "doe"
	userList, sdkErr := api.User.List(context.Background(), "?filter=firstname:'john'¤lastname:'doe'")
	if sdkErr != nil {
		// Error
	}

	for _, user := range userList.Users {
		fmt.Printf("%s %s lives in %s", *user.Firstname, *user.Lastname, *user.City)
	}

}
