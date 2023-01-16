package response

import "time"

type UserResponse struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Address    string         `json:"address"`
	Phone      string         `json:"phone"`
	Created_At time.Time      `json:"created_at"`
	Updated_At time.Time      `json:"updated_at"`
}