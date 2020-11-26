package models

// its use for definition database GORM
type SampleCrud struct {
 	ID                 int      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"` 
	Text_input         string   `gorm:"type:varchar(50)"`
	Text_area          string   `gorm:"type:varchar(255)"`
	Created_by       int      `gorm:"type:int(10); index;"` 
	Updated_by       int      `gorm:"type:int(10); index;"` 
  	Status            string   `gorm:"type:enum('Y','N'); comment:'Y:Active, N:Inactive'; default:'Y'"`
  	CreatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP"`
  	UpdatedAt         string   `gorm:"type:timestamp(0); default:CURRENT_TIMESTAMP""`
  	Additional        string   `gorm:"type:varchar(191)"`
}

func (SampleCrud) TableName() string {
  return "tb_sample_crud"
}


// its use for call model from controllers
type ModelSampleCrud struct {
    ID              string   `json:"id"` 
    Text_input      string   `json:"text_input"`
    Text_area       string   `json:"text_area"`
    Status          string   `json:"status"`
    CreatedAt       string   `json:"created_at"`
    UpdatedAt       string   `json:"updated_at"`
    Additional      string   `json:"additional"`
}

