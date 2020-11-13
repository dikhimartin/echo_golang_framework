package models

type Permission struct {
  	ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  	Name        	  string   `gorm:"type:varchar(50)"`
  	Additional        string   `gorm:"type:varchar(191)"`
}
func (Permission) TableName() string {
  return "tb_permission"
}

type SchemePermission struct {
  	ID                string   `json:"id"` 
  	Name        	  string   `json:"name"`
  	Additional        string   `json:"additional"`
}


