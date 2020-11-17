package models

// its use for definition database GORM
type Permission struct {
  	ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  	Name        	  string   `gorm:"type:varchar(50)"`
  	Additional        string   `gorm:"type:varchar(191)"`
}
func (Permission) TableName() string {
  return "tb_permission"
}

// its use for call model from controllers
type ModelPermission struct {
  	ID                string   `json:"id"` 
  	Name        	    string     `json:"name"`
  	Additional        string   `json:"additional"`
    CheckOrUncheck    string   `json:"check_or_uncheck"`
}

// == its use for migration view_schema
type SchemePermission struct {
  	ID                string   `json:"id"` 
  	Name        	  string   `json:"name"`
  	Additional        string   `json:"additional"`
}



