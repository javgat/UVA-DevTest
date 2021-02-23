package userdao

// User user
//
// swagger:model User
type User struct {

	// email
	// Example: carlos@mail.com
	// Required: true
	Email *string `json:"email"`

	// pwhash
	// Example: e$ia9s7ATDGba39pakscAKs
	// Required: true
	Pwhash *string `json:"pwhash"`

	// username
	// Example: carlosg72
	// Required: true
	Username *string `json:"username"`
}
