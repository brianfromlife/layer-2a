package user

// User ...
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	CreatedAt   string `json:"created_at"`
}

// FactoryConfig parameters for fields of user entity
type FactoryConfig struct {
	MinIDLength          int
	MinNameLength        int
	MinEmailLength       int
	MinPhoneLength       int
	MinCountryCodeLength int
}

// Factory class to manage functionality of all entities
type Factory struct {
	fc *FactoryConfig
}

// NewFactory create new instance
func NewFactory(config *FactoryConfig) *Factory {
	return &Factory{fc: config}
}

// validateUser validate all values inside of user entity
func (f *Factory) validateUser(u *User) error {
	return nil
}

// UnmarshalUserFromDatabase transform database values into domain user
func (f *Factory) UnmarshalUserFromDatabase(ID int64, Name, Email, CountryCode, Phone, CreatedAt string) (*User, error) {

	u := &User{
		ID:          ID,
		Name:        Name,
		Email:       Email,
		Phone:       Phone,
		CountryCode: CountryCode,
		CreatedAt:   CreatedAt,
	}

	if err := f.validateUser(u); err != nil {
		return nil, err
	}

	return u, nil
}
