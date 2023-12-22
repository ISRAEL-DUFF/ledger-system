package types

type API_KEY_TYPE string

const (
	TEST_KEY API_KEY_TYPE = "test"
	LIVE_KEY API_KEY_TYPE = "live"
)

type CreateOrganization struct {
	Name         string
	Address      string
	EmailAddress string
	PhoneNumber  string
	OwnerID      string
}

type CreateOrganizationDto struct {
	Name         string
	Address      string
	EmailAddress string
	PhoneNumber  string
}

type CreateUser struct {
	EmailAddress string
	FullName     string
	Password     string
	PhoneNumber  string
}

type CreateAPIKEY struct {
	OrganizationID string
	TestSecretKey  string
	TestPublicKey  string
	LivePublicKey  string
	LiveSecretKey  string
}

type APIKEY struct {
	PublicKey string
	Secretkey string
}

type CreateAPIKEYResponse struct {
	TestKeys APIKEY
	LiveKeys APIKEY
}
