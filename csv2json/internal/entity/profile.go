package entity

// Profile declare an user details
type Profile struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Role        string `json:"role"`
	Phone       string `json:"phone"`
}
