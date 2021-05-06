package main 

import (
    "log"
)

func (telepher bot) Start()(Update){
  ch := make(chan Update)
 go telepher.Poll(ch)
  	for {
		 res, _ := <-ch
        return res
     }
  
}
func(telepher bot) Poll(ch chan Update)  {
    LastUpdateID :=  10
  for{

k , err := telepher.GetUpdates(LastUpdateID, 1)
	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())

	}

  var key = len(k) - 1
  if key != -1 {
 LastUpdateID = (k[key].UpdateID)+1 
	
   ch <- k[key]
}
}
}
