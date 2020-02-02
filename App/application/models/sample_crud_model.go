package models

type SampleCrud struct {
	Id        	  							string `json:"id"`
	Text_input        						string `json:"text_input"`
	Text_area        						string `json:"text_area"`
	Created_by        						string `json:"created_by"`
	Updated_by        						string `json:"Updated_by"`
	Created_at        						string `json:"created_at"`
	Updated_at        						string `json:"updated_at"`
	Status        							string `json:"status"`
	Additional        						string `json:"additional"`
}

type SampleCruds struct {
	 SampleCruds []SampleCrud `json:"sample_crud"`
}
