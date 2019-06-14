package user

//User is the abstraction for registering new users
type User struct {
	Name     string
	Username string
	Email    string
	Password string
	Location [2]float32
}

//New creates a new user
func New(name string, username string, email string, password string, location [2]float32) *User {
	return &User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
		Location: location,
	}
}
