package models

// its use for definition database GORM
type SettingUser struct {
  ID                int              `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  Full_name         string           `gorm:"type:varchar(50)"`
  Username          string           `gorm:"type:varchar(50)"`
  Email           	string           `gorm:"type:varchar(50)"`
  Telephone         string           `gorm:"type:varchar(50)"`
  Address         	string           `gorm:"type:varchar(255)"`
  Gender       	    string           `gorm:"type:enum('L','P')"`
  Image             string           `gorm:"type:varchar(255)"`
  Password         	string           `gorm:"type:varchar(100)"`
  Auth_token        string           `gorm:"type:varchar(255)"`
  Status            string           `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  CreatedAt         string           `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  UpdatedAt         string           `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  Additional        string           `gorm:"type:varchar(191)"`
}
func (SettingUser) TableName() string {
  return "tb_setting_user"
}

// its use for call model from controllers
type ModelUser struct {
    ID                       string      `json:"id"`
    Id_setting_grup          string      `json:"id_setting_grup"`
    Name_Grup                string      `json:"name_grup"`
    Full_name                string      `json:"full_name"`
    Username                 string      `json:"username"`
    Email                    string      `json:"email"`
    Telephone                string      `json:"telephone"`
    Address                  string      `json:"address"`
    Gender                   string      `json:"gender"`
    Password                 string      `json:"password"`
    Status                   string      `json:"status"`
    Image                    string      `json:"image"`
    Auth_token               string      `json:"auth_token"`
    CreatedAt                string      `json:"created_at"`
    Additional               string      `json:"additional"`
}

// == its use for migration view_schema
type SchemeUser struct {
  ID          string      `json:"id"`
  Full_name   string      `json:"full_name"`
  Username    string      `json:"username"`
  Email       string      `json:"email"`
  Telephone   string      `json:"telephone"`
  Address     string      `json:"address"`
  Gender      string      `json:"gender"`
  Password    string      `json:"password"`
  Status      string      `json:"status"`
  Auth_token  string      `json:"auth_token"`
  CreatedAt   string      `json:"created_at"`
}
