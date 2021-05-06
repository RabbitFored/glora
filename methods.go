package main
import (
"strconv"
  "net/url"
 "encoding/json"
 _ "fmt" )


func (telepher bot)GetMe()(*User,error){
data ,err :=telepher.post("getMe",nil)

	if err != nil {
    return nil, err
	} 
	var user struct {
		Result User
	}
 	if err := json.Unmarshal(data,&user); err != nil {
   return nil, err
	}  
  return &user.Result , nil
}
func (telepher bot)GetUpdates(offset int,limit int)([]Update,error) {

v := url.Values{}
v.Add("offset",strconv.Itoa(offset))
v.Add("limit", strconv.Itoa(limit))

data ,_ :=telepher.post("getUpdates", v)

	var update struct {
		Result []Update
	}

  
	if err := json.Unmarshal(data,&update); err != nil {
   return nil, err
	}  
return update.Result , nil

}
func (telepher bot) SendMessage(chat_id string ,text string,option *Options)(*Message,error){
v := url.Values{}
v.Add("chat_id",chat_id)
v.Add("text", text)
v.Add("parse_mode",telepher.ParseMode)
v.Add("reply_to_message_id",strconv.Itoa(option.ReplyToMessageID))
vd ,err :=telepher.post("sendMessage", v)

var res struct{ 
Result *Message
}
 err = json.Unmarshal(vd, &res) 
    if err != nil {
return nil, err }
return res.Result,nil

}


func (telepher bot) SendPhoto(chat_id string,path string){
    
param := map[string]string{
		"chat_id":   chat_id,
	}


telepher.postWithFile("photo",path,param)


}

