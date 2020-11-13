package models

type SettingPrivilege struct {
  ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
  Code_privilege    string   `gorm:"type:varchar(255)"`
  Name_menu    		string   `gorm:"type:varchar(255)"`
  Remarks    		string   `gorm:"type:varchar(255)"`
  Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        string   `gorm:"type:varchar(191)"`
}
func (SettingPrivilege) TableName() string {
  return "tb_setting_privilege"
}


