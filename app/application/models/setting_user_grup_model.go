package models


// its use for definition database GORM
type SettingUserGrup struct {
  ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  Id_setting_user   int      `gorm:"type:int(10); index; NOT NULL"` 
  Id_setting_grup   int      `gorm:"type:int(10); index; NOT NULL"` 
  Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        string   `gorm:"type:varchar(191)"`
}
func (SettingUserGrup) TableName() string {
  return "tb_setting_user_grup"
}

// == its use for migration view_schema
type SchemeUserGroup struct {
	ID            		string      `json:"id"`
	Id_setting_user 	string      `json:"id_setting_user"`
	Id_setting_grup 	string      `json:"id_setting_grup"`
	Status        		string      `json:"status"`
}