package models

type GetDataLogin struct {
  Id_user               int   		 `json:"id_user"`		     	
  Id_group   			     int      	     `json:"id_group"`
  Name_grup         	 string      `json:"name_grup"`     
  Full_name            string      `json:"full_name"`     
  Username          	 string      `json:"username"`     
  Email           		 string      `json:"email"`     
  Telephone         	 string      `json:"telephone"`     
  Address         		 string      `json:"address"`     
  Gender       	    	 string      `json:"gender"`    
  Image             	 string      `json:"image"`     
  Status            	 string      `json:"status"`    
}