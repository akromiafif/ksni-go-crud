package response

type UserResponse struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Address    string         `json:"address"`
	Phone      string         `json:"phone"`
}