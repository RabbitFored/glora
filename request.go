package main

import (
  "net/http"
  "net/url"
  "bytes"
  "io"
  "io/ioutil"
 	"path/filepath"
  "mime/multipart"
  "os"
 _ "encoding/json"
  "fmt"
  
  )
 
  
func (telepher bot) post(endpoint string, params url.Values )([]byte,error){
  var response *http.Response
  var err error
  
  
  URL := telepher.BaseURL + telepher.Token + "/" + endpoint

if params == nil {
		response, err = http.Post(URL, "application/x-www-form-urlencoded", nil)
	} else {
	response, err = http.PostForm(URL,params)
	}
	

    if err != nil {
return nil,err
}
    
  data, err := ioutil.ReadAll(response.Body)

         if err != nil {
   return nil , err
 }

return data,nil
}

func (telepher bot) postWithFile(filetype string, file string, params map[string]string){ 
  endpoint := "send" + filetype 
URL := telepher.BaseURL + telepher.Token + "/" + endpoint
File, err := os.Open(file)
    if err != nil {
fmt.Println(err)
	}
    	defer File.Close()
      body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(filetype, filepath.Base(file))
	if err != nil {
fmt.Println(err)
	}
  	_, err = io.Copy(part, File)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}	
  
  err = writer.Close()
	if err != nil {
fmt.Println(err)
	}
  req, err := http.NewRequest("POST", URL, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
  	if err != nil {
fmt.Println(err)
	}
  resp, _ := telepher.Client.Do(req)
fmt.Println(resp)

}



