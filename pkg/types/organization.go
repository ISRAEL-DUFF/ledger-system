package types

type CreateOrganization struct {
	Name         string
	Address      string
	EmailAddress string
	PhoneNumber  string
	OwnerID      string
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
