package m5adminapi

import (
	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/services"
)

// Wrapper over client.Client as client exists within internal
type Client = client.Client

// Parse strategy control (strict vs lenient JSON decoding)
// Strict will throw errors on unknown fields, therefor will be more fragile on unsupported api versions.
// Default is "Lenient"
type ParseStrategy = client.ParseStrategy

const (
	// allow unknown fields in response
	ParseStrategyLenient = client.ParseStrategyLenient

	// disallow unknown fields in response
	ParseStrategyStrict = client.ParseStrategyStrict
)

// Sets used parse Strategy to v
func SetParseStrategy(v ParseStrategy) error {
	return client.SetParseStrategy(v)
}

// gets current parse Strategy
func GetParseStrategy() ParseStrategy {
	return client.GetParseStrategy()
}

// Creates a new client, then can be used in NewAdminApiConnectionFromClient()
func NewClient() *Client {
	return &client.Client{}
}

// root object for api services
type AdminApiConnection struct {
	Client *client.Client

	// services
	Auth             *services.AuthService
	User             *services.UserService
	Card             *services.CardService
	Version          *services.VersionService
	System           *services.SystemService
	Department       *services.DepartmentService
	Usergroup        *services.UsergroupService
	Floor            *services.FloorService
	Unit             *services.UnitService
	Domain           *services.DomainService
	Accessgroup      *services.AccessgroupService
	Preselection     *services.PreselectionService
	Machinegrouptype *services.MachinegrouptypeService
	Machinegroup     *services.MachinegroupService
	Setting          *services.SettingService
	Connection       *services.ConnectionService
	Administrator    *services.AdministratorService
	Offlineunit      *services.OfflineunitService
	Date             *services.DateService
	Function         *services.FunctionService
	Timezone         *services.TimezoneService
	Readeraccess     *services.Readeraccess
	Event            *services.EventService

	// Unsafe service to interact with http client directly.
	// Should be avoided if possible
	Unsafe *services.UnsafeService
}

/*
// NewAdminApiConnection creates and returns a new AdminApiConnection with all services initialized.
//
// Parameters:
//   - address: The base server address (e.g., "http://localhost" or "https://api.example.com")
//   - urlPath: The API path prefix (e.g., "/m5adminapi/api")
//   - cred: Authentication credentials for the API connection
//   - tlsConfig: TLS configuration for HTTPS connections (can be nil for HTTP)
//
// Returns a fully initialized AdminApiConnection ready for API calls.
func NewAdminApiConnection(address string, urlPath string, cred *models.Credentials, tlsConfig *tls.Config) *AdminApiConnection {
	c := &client.Client{
		Address:     address,
		Path:        urlPath,
		Credentials: *cred,
		TLSConfig:   tlsConfig,
	}

	return registerCommonServices(c)
}
*/

// Allows more control over http client.
func NewAdminApiConnection(client *client.Client) *AdminApiConnection {

	return registerCommonServices(client)
}

func registerCommonServices(client *client.Client) *AdminApiConnection {
	return &AdminApiConnection{
		Client: client,

		// initiate resource services
		Auth:             services.NewAuthService(client),
		User:             services.NewUserService(client),
		Version:          services.NewVersionService(client),
		System:           services.NewSystemService(client),
		Card:             services.NewCardService(client),
		Department:       services.NewDepartmentService(client),
		Usergroup:        services.NewUsergroupService(client),
		Floor:            services.NewFloorService(client),
		Unit:             services.NewUnitService(client),
		Domain:           services.NewDomainService(client),
		Accessgroup:      services.NewAccessgroupService(client),
		Preselection:     services.NewPreselectionService(client),
		Machinegrouptype: services.NewMachinegrouptypeService(client),
		Machinegroup:     services.NewMachinegroupService(client),
		Setting:          services.NewSettingService(client),
		Connection:       services.NewConnectionService(client),
		Administrator:    services.NewAdministratorService(client),
		Offlineunit:      services.NewOfflineunitService(client),
		Date:             services.NewDateService(client),
		Function:         services.NewFunctionService(client),
		Timezone:         services.NewTimezoneService(client),
		Readeraccess:     services.NewReaderaccess(client),
		Event:            services.NewEventService(client),

		Unsafe: services.New(client),
	}
}
