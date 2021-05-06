package main

import (
  
  "net/http"
  )
  
  
func Bot(token string,settings *Settings) (*bot,error) {

  telepher := &bot{}

  telepher.Token = token

  if settings != nil{
    if settings.BaseURL == ""{
      settings.BaseURL = DefaultBaseURL
    }
    if settings.BaseFileURL == ""{
      settings.BaseFileURL =  DefaultBaseFileURL
    }
    if settings.Client == nil{
      settings.Client = DefaultClient
    }
     if settings.ParseMode == ""{
      settings.ParseMode = DefaultParseMode
    }
    telepher.BaseURL = settings.BaseURL
    telepher.BaseFileURL = settings.BaseFileURL
    telepher.Client = settings.Client

  } else {

  telepher.BaseURL =DefaultBaseURL
  telepher.BaseFileURL =  DefaultBaseFileURL
  telepher.Client =  DefaultClient
  telepher.ParseMode = DefaultParseMode
  }
  
  me, err := telepher.GetMe()
	if err != nil {
		return nil, err
	}
  telepher.Self = me
  return telepher , nil
  
}

type bot struct{
  Token string
  BaseURL string
  BaseFileURL string
  Client      *http.Client
  ParseMode string
  Self          *User
}

type Settings struct{
  BaseURL string
  BaseFileURL string
  Client *http.Client
  ParseMode string
}
