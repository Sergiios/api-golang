package domain

type Central struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	MAC  string `json:"MAC" gorm:"unique"`
	IP   string `json:"IP" gorm:"unique"`
}
