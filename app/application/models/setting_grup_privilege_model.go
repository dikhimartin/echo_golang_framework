package models

// its use for definition database GORM
type SettingGrupPrivilege struct {
  ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
  Id_setting_grup   int      `gorm:"type:int(10); index; NOT NULL"` 
  Remarks         	string   `gorm:"type:varchar(50)"`
  Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        string   `gorm:"type:varchar(191)"`
}
func (SettingGrupPrivilege) TableName() string {
  return "tb_setting_grup_privilege"
}

// == its use for migration view_schema
type SchemeGrupPrivilege struct {
  	ID              	  string   `json:"id"` 
  	Id_setting_grup       string   `json:"id_setting_grup"`
  	Remarks       		  string   `json:"remarks"`
  	Status        		  string   `json:"status"`
}


