package models

type SampleCrud struct {
  	ID                int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
  	Text_input        string   `gorm:"type:varchar(50)"`
  	Text_area         string   `gorm:"type:varchar(50)"`
  	Created_by   	  int      `gorm:"type:int(10); index; NOT NULL"` 
  	Updated_by   	  int      `gorm:"type:int(10); index; NOT NULL"` 
  	CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  	UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
 	Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  	Additional        string   `gorm:"type:varchar(191)"`
}
func (SampleCrud) TableName() string {
  return "tb_sample_crud"
}

