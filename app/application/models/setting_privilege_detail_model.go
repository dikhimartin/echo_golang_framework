package models

type SettingPrivilegeDetail struct {
  ID                	 int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  Id_setting_privilege   int      `gorm:"type:int(10); index; NOT NULL"` 
  Permissions            string   `gorm:"type:enum('1','2', '3', '4'); comment:'1: Create, 2: Read/View, 3: Edit, 4: Delete '"`
  Additional        	 string   `gorm:"type:varchar(191)"`
}
func (SettingPrivilegeDetail) TableName() string {
  return "tb_setting_privilege_detail"
}


