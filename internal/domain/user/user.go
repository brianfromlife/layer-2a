package user

// User
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	CreatedAt   string `json:"created_at"`
}

// FactoryConfig
type FactoryConfig struct{}

// Factory
type Factory struct {
	fc *FactoryConfig
}

// NewFactory
func NewFactory(config *FactoryConfig) *Factory {
	return &Factory{fc: config}
}

// validateUser
func (f *Factory) validateUser(u *User) error {
	return nil
}

// UnmarshalUserFromDatabase
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
