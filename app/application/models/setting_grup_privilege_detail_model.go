package models


// its use for definition database GORM
type SettingGrupPrivilegeDetail struct {
  ID                		  int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
  Id_setting_grup_privilege   int      `gorm:"type:int(10); index; NOT NULL"` 
  Code_permissions         	  string   `gorm:"type:char(50)"`
  CreatedAt         		  string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         		  string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        		  string   `gorm:"type:varchar(191)"`
}
func (SettingGrupPrivilegeDetail) TableName() string {
  return "tb_setting_grup_privilege_detail"
}

// == its use for migration view_schema
type SchemeGrupPrivilegeDetail struct {
  	ID              	 			string   `json:"id"` 
  	Id_setting_grup_privilege       string   `json:"id_setting_grup_privilege"`
  	Code_permissions       		  	string   `json:"code_permissions"`
}


