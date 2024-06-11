package domain

type (
	Category struct {
		ID      uint64    `json:"id" gorm:"primary_key"`
		Name    string    `json:"name" gorm:"unique"`
		Product []Product `gorm:"foreignKey:CategoryID"`
	}

	CreateCategory struct {
		Name string `json:"name"`
	}

	CategoryResponse struct {
		ID   uint64 `json:"id" gorm:"primary_key"`
		Name string `json:"name" gorm:"unique"`
	}
)
