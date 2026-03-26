# Adminapi SDK for Go

This is a **Non-offical** sdk for M5 adminapi.

Using this SDK does not replace reading the official API documentation.

## Maintenance

This project is currently maintained on a best-effort basis.
Do not expect an *Alwyas up do date* SDK for the API. Since it is a new project, major changes can occur as the SDK matures.

### Tested api version
Currently:
- adminapi ver 1.5.4

*No api version enforcement, but changes in api response can cause issues.*

## Examples
Check examples folder!

## Quick start

```go

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
```


### Available Services

The SDK provides the following services accessible through the `AdminApiConnection`:

- Accessgroup
- Administrator
- Auth (login/logout)
- Card
- Connection
- Date
- Department
- Domain
- Event
- Floor
- Function
- Machinegroup
- Machinegrouptype
- Offlineunit
- Preselection
- Readeraccess
- Setting
- System
- Timezone
- Unit
- User
- Usergroup
- Version