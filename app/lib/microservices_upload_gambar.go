package lib

import (
	"fmt"
	"bytes"
	"io"
	"net/http"
	"encoding/json"
	"mime/multipart"
	"github.com/labstack/echo"
)

type UploadFileMultipleStruct struct {
	Filename string `json:"Filename"`
	Header   struct {
		ContentDisposition []string `json:"Content-Disposition"`
		ContentType        []string `json:"Content-Type"`
	}   `json:"Header"`
	Size int `json:"Size"`
}

type ResponseUpload struct{
	Filename      string         `json:"file_name"` 
	RelativePath  string         `json:"relative_path"` 
	Status        ResponseStatus `json:"status"` 
}

type ResponseStatus struct{
	Code           int    `json:"code"`
	Message        string `json:"message"`
}

func UploadFile(c echo.Context, url_api, path, file_name, field_image string) (ResponseUpload) {
	form, err := c.MultipartForm()
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}
	files := form.File[field_image]
	if files == nil{
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 204,
				Message : "no_content_found",
			},
		}
		return response
	}

	file, _ 	   := c.FormFile(field_image)
	file_image, _  := file.Open()
	defer file_image.Close()

	if file_name == ""{
		file_name = file.Filename
	}else{
		file_name = file_name + "_" + file.Filename
	}

	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)
	fileWriter, err := multiPartWriter.CreateFormFile(field_image, file_name)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	_, err = io.Copy(fileWriter, file_image)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	fieldPath, err := multiPartWriter.CreateFormField("path")
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	_, err = fieldPath.Write([]byte(path))
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	fieldImage, err := multiPartWriter.CreateFormField("field_image")
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	_, err = fieldImage.Write([]byte(field_image))
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}
	multiPartWriter.Close()


	req, err := http.NewRequest("POST", url_api, &requestBody)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	file_name     = fmt.Sprintf("%v", result["data"])

	response_data := ResponseUpload{
		Filename       : file_name,
		RelativePath   : path +"/"+ file_name,
		Status : ResponseStatus{
			Code    : 200,
			Message : "OK",
		},
	}

	return response_data
}

func ReUploadFile(c echo.Context, url_api, path, old_file, file_name, field_image string) (ResponseUpload){
	form, err := c.MultipartForm()
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}
	files := form.File[field_image]
	if files == nil{
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 204,
				Message : "no_content_found",
			},
		}
		return response
	}
	file, _ 	   := c.FormFile(field_image)
	file_image, _  := file.Open()
	defer file_image.Close()


	if file_name == ""{
		file_name = file.Filename
	}else{
		file_name = file_name + "_" + file.Filename
	}


	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multiPartWriter.CreateFormFile(field_image, file_name)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	_, err = io.Copy(fileWriter, file_image)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	fieldPath, err := multiPartWriter.CreateFormField("path")
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}
	_, err = fieldPath.Write([]byte(path))
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	if old_file != ""{
		fieldOldFile, err := multiPartWriter.CreateFormField("old_file")
		if err != nil {
			logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
		}
		_, err = fieldOldFile.Write([]byte(old_file))
		if err != nil {
			logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
		}
	}

	fieldImage, err := multiPartWriter.CreateFormField("field_image")
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}
	_, err = fieldImage.Write([]byte(field_image))
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}
	multiPartWriter.Close()


	req, err := http.NewRequest("POST", url_api, &requestBody)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logs.Println(err)
		response := ResponseUpload{
			Status : ResponseStatus{
				Code    : 500,
				Message : "internal_server_error",
			},
		}
		return response
	}

	// result
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	file_name     = fmt.Sprintf("%v", result["data"])
	response_data := ResponseUpload{
		Filename       : file_name,
		RelativePath   : path +"/"+ file_name,
		Status : ResponseStatus{
			Code    : 200,
			Message : "OK",
		},
	}

	return response_data
}

func DeleteFile(c echo.Context, url_api, path, old_file string) bool{
	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)

	fieldPath, err := multiPartWriter.CreateFormField("path")
	if err != nil {
		logs.Println(err)
		return false
	}
	_, err = fieldPath.Write([]byte(path))
	if err != nil {
		logs.Println(err)
		return false
	}

	fieldOldFile, err := multiPartWriter.CreateFormField("old_file")
	if err != nil {
		logs.Println(err)
		return false
	}
	_, err = fieldOldFile.Write([]byte(old_file))
	if err != nil {
		logs.Println(err)
		return false
	}
	multiPartWriter.Close()

	req, err := http.NewRequest("POST", url_api, &requestBody)
	if err != nil {
		logs.Println(err)
		return false
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		logs.Println(err)
		return false
	}

	return true
}

func UploadFileMultiple(c echo.Context, url_api, path, field_image string) []UploadFileMultipleStruct{
	form, err := c.MultipartForm()
	if err != nil {
		logs.Println(err)
		return nil
	}
	files := form.File[field_image]
	if files == nil{
		logs.Println(err)
		return nil
	}

	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)

	for key , value := range files {
		file, _ := files[key].Open()
		defer file.Close()

		// Initialize the file field
		fileWriter, err := multiPartWriter.CreateFormFile(field_image, value.Filename)
		if err != nil {
			logs.Println(err)
		}

		// Copy the actual file content to the field field's writer
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			logs.Println(err)
		}
	}

	fieldPath, err := multiPartWriter.CreateFormField("path")
	if err != nil {
		logs.Println(err)
	}
	_, err = fieldPath.Write([]byte(path))
	if err != nil {
		logs.Println(err)
	}

	fieldImage, err := multiPartWriter.CreateFormField("field_image")
	if err != nil {
		logs.Println(err)
	}
	_, err = fieldImage.Write([]byte(field_image))
	if err != nil {
		logs.Println(err)
	}
	multiPartWriter.Close()

	req, err := http.NewRequest("POST", url_api, &requestBody)
	if err != nil {
		logs.Println(err)
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logs.Println(err)
	}


	var json_data map[string]interface{}
	json.NewDecoder(response.Body).Decode(&json_data)
	jsonString, _ := json.Marshal(json_data["data"])


	jsonByteArray := []byte(string(jsonString))

    var result []UploadFileMultipleStruct
    err = json.Unmarshal(jsonByteArray, &result)
    if err != nil {
    	logs.Println(err)
    }

	return result
}

func ReUploadFileMultiple(c echo.Context, url_api, path, field_image string) []UploadFileMultipleStruct{

	form, err := c.MultipartForm()
	if err != nil {
		logs.Println(err)
		return nil
	}
	files := form.File[field_image]
	if files == nil{
		logs.Println(err)
		return nil
	}

	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)

	// multiple_upload
	for key , value := range files {
		file, _ := files[key].Open()
		defer file.Close()

		// Initialize the file field
		fileWriter, err := multiPartWriter.CreateFormFile(field_image, value.Filename)
		if err != nil {
			logs.Println(err)
		}

		// Copy the actual file content to the field field's writer
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			logs.Println(err)
		}
	}

	fieldPath, err := multiPartWriter.CreateFormField("path")
	if err != nil {
		logs.Println(err)
	}
	_, err = fieldPath.Write([]byte(path))
	if err != nil {
		logs.Println(err)
	}

	fieldImage, err := multiPartWriter.CreateFormField("field_image")
	if err != nil {
		logs.Println(err)
	}
	_, err = fieldImage.Write([]byte(field_image))
	if err != nil {
		logs.Println(err)
	}
	multiPartWriter.Close()

	req, err := http.NewRequest("POST", url_api, &requestBody)
	if err != nil {
		logs.Println(err)
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logs.Println(err)
	}

	var json_data map[string]interface{}
	json.NewDecoder(response.Body).Decode(&json_data)
	jsonString, _ := json.Marshal(json_data["data"])

	jsonByteArray := []byte(string(jsonString))

    var result []UploadFileMultipleStruct
    err = json.Unmarshal(jsonByteArray, &result)
    if err != nil {
        logs.Println(err)
    }

	return result
}
