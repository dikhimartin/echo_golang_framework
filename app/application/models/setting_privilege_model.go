package models

// its use for definition database GORM
type SettingPrivilege struct {
  ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
  Code_privilege    string   `gorm:"type:varchar(255)"`
  Name_menu    		  string   `gorm:"type:varchar(255)"`
  Remarks    		    string   `gorm:"type:varchar(255)"`
  Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        string   `gorm:"type:varchar(191)"`
}
func (SettingPrivilege) TableName() string {
  return "tb_setting_privilege"
}


// its use for call model from controllers
type ModelPrivilege struct {
    ID                 string   `json:"id"` 
    Code_privilege     string   `json:"code_privilege"`
    Name_menu          string   `json:"name_menu"`
    Remarks            string   `json:"remarks"`
    Status             string   `json:"status"`
}


// == its use for migration view_schema
type SchemePrivilege struct {
  	ID              	 string   `json:"id"` 
  	Code_privilege     string   `json:"code_privilege"`
  	Name_menu       	 string   `json:"name_menu"`
  	Remarks       		 string   `json:"remarks"`
  	Status       		   string   `json:"status"`
}







