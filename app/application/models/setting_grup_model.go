package models

// its use for definition database GORM
type SettingGrup struct {
  ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  Name_Grup         string   `gorm:"type:varchar(50)"`
  Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        string   `gorm:"type:varchar(191)"`
}
func (SettingGrup) TableName() string {
  return "tb_setting_grup"
}

// its use for call model from controllers
type ModelGrup struct {
    ID              string   `json:"id"` 
    Name_Grup       string   `json:"name_grup"`
    Status          string   `json:"status"`
}

// == its use for migration view_schema
type SchemeGroup struct {
  	ID              string   `json:"id"` 
  	Name_Grup       string   `json:"name_grup"`
  	Status        	string   `json:"status"`
}


