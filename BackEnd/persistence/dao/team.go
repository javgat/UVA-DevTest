package dao

// Team team
//
// teamdao team
type Team struct {
	// id
	// Identificador SQL
	ID int

	// description
	// Example: DevTest Team
	Description string `json:"description,omitempty"`

	// teamname
	// Example: devtestTeam
	// Required: true
	// Pattern: ^[^@ \t\r\n]+$
	Teamname *string `json:"teamname"`
}
