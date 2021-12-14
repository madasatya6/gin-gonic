package upload

import "fmt"
import "os"
import "io"
import "github.com/gin-gonic/gin"
import "path/filepath"
import "bytes"
import "errors"
import "strings"
import "time"
import "net/http"
import "mime/multipart"

var UploadPath = "app/static/uploads"

/********* 
* @author madasatya6
* Validate file upload
* @param min_size in Kb
* @param max_size in Kb
*/
func FileValidate(c *gin.Context, form_name string, extensions []string, min_size int64, max_size int64, required bool) (string, error) {

	var file_name string
	max_size = 1024 * max_size
	min_size = 1024 * min_size

	file, err := c.FormFile(form_name)
	if err != nil {
		
		var err_status = ""
		switch err {
			case http.ErrMissingFile: 
				err_status = "File required"
			default:
				err_status = "File error"
		}

		if required {
			if err_status == "File required" {
				return file_name, errors.New(fmt.Sprintf("File %s dibutuhkan", form_name))
			} else {
				return file_name, errors.New("File "+ form_name +" error")
			}
		} else {
			if err_status == "File required" {
				return file_name, nil
			} else {
				return file_name, err 
			}
		}

	}

	src, err := file.Open()
	if err != nil {
		return file_name, err 
	}
	defer src.Close()
	
	file_name = fmt.Sprintf("%v", file.Filename)
	
	var buff bytes.Buffer 
	fileSize, err := buff.ReadFrom(src)
	if err != nil {
		return file_name, err 
	}
	
	if fileSize < min_size {
		eror := errors.New("Ukuran file "+ form_name +" kurang dari " + fmt.Sprintf("%v kb", (min_size/1024)))
		return file_name, eror
	}
	if fileSize > max_size {
		eror := errors.New("Ukuran file "+ form_name +" lebih dari " + fmt.Sprintf("%v kb", max_size/1024))
		return file_name, eror
	}

	var fileExtension = filepath.Ext(file.Filename)
	onlyExt := strings.ReplaceAll(fileExtension, ".", "")
	status := true
	string_extensions := ""
	for i, name := range extensions {
		if i == 0 {
			string_extensions = name 
		} else {
			string_extensions += ", " + name
		}
		if name == onlyExt {
			status = false
		}
	}
	
	if status {
		eror := errors.New("Ekstension "+ form_name +" yang diperbolehkan " + string_extensions)
		return file_name, eror
	} 

	return file_name, nil
}

/**
* Jika rename dikosongkan maka akan dinamai secara otomatis
*/
func UploadFileAndRename(c *gin.Context, form_name string, upload_path string, rename string) (string ,error) {
	
	var file_name string
	file, err := c.FormFile(form_name)
	
	if err != nil {
		var err_status = ""
		switch err {
			case http.ErrMissingFile: 
				err_status = "File required"
			default:
				err_status = "File error"
		}

		if err_status == "File required" {
			return file_name, nil
		} else {
			return file_name, err 
		}
	}

	src, err := file.Open()
	if err != nil {
		return file_name, err
	}
	defer src.Close()

	dir, err := os.Getwd()
	if err != nil {
		return file_name, err 
	}

	tme := time.Now()
	removeWhiteSpace := strings.ReplaceAll(fmt.Sprintf("%v",file.Filename), " ", "_")
	file_name = fmt.Sprintf("%v%v%v%v%v%v%v-%v",tme.Year(), tme.Month(), tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), (tme.UnixNano()/1000000), removeWhiteSpace)
	
	if rename != "" {
		file_name = rename + filepath.Ext(file.Filename)
	}

	destination := filepath.Join(dir, UploadPath, upload_path, file_name)
	targetLocation, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return file_name, err 
	}

	defer targetLocation.Close()

	if _, err = io.Copy(targetLocation, src); err != nil {
		return file_name, err
	}

	return file_name, nil
}

func DeleteFile(location string) error {
	e := os.Remove(location)
	if e != nil {
		return e 
	}
	return nil
}

/**
* Validate Upload multiple file
*/
func ValidateMultipleFile(file *multipart.FileHeader, extensions []string, min_size int64, max_size int64, required bool) (string, error) {
	
	//file in kb
	min_size = 1024 * min_size
	max_size = 1024 * max_size

	var fileName string 
	src, err := file.Open()
	if err != nil {
		return fileName, err 
	}
	defer src.Close()

	fileName = fmt.Sprintf("%v", file.Filename)

	var buff bytes.Buffer 
	fileSize, err := buff.ReadFrom(src)

	if err != nil {
		return fileName, err 
	}

	if fileSize < min_size {
		eror := errors.New("Ukuran file kurang dari "+fmt.Sprintf("%v kb", (min_size/1024)))
		return fileName, eror 
	}

	if fileSize > max_size {
		eror := errors.New("Ukuran file lebih dari "+fmt.Sprintf("%v kb", (max_size/1024)))
		return fileName, eror 
	}

	fileExtension := filepath.Ext(file.Filename)
	onlyExt := strings.ReplaceAll(fileExtension, ".", "")
	status := true 
	stringExtensions := strings.Join(extensions, ",")

	for _, name := range extensions {
		if name == onlyExt {
			status = false 
		}
	}

	if status {
		err = errors.New("Ektensi yang dibolehkan "+stringExtensions)
		return fileName, err 
	}

	return fileName, nil
}

/**
* Upload multiple file and rename
*/
func UploadMultipleFileAndRename(file *multipart.FileHeader, upload_path string, rename string) (string ,error) {
	
	var file_name string

	src, err := file.Open()
	if err != nil {
		return file_name, err
	}
	defer src.Close()

	dir, err := os.Getwd()
	if err != nil {
		return file_name, err 
	}

	tme := time.Now()
	removeWhiteSpace := strings.ReplaceAll(fmt.Sprintf("%v",file.Filename), " ", "_")
	file_name = fmt.Sprintf("%v%v%v%v%v%v%v-%v",tme.Year(), tme.Month(), tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), (tme.UnixNano()/1000000), removeWhiteSpace)
	
	if rename != "" {
		file_name = rename + filepath.Ext(file.Filename)
	}

	destination := filepath.Join(dir, UploadPath, upload_path, file_name)
	targetLocation, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return file_name, err 
	}

	defer targetLocation.Close()

	if _, err = io.Copy(targetLocation, src); err != nil {
		return file_name, err
	}

	return file_name, nil
}
