package models

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
