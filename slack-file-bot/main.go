package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main (){

	os.Setenv("SLACK_BOT_TOKEN","xoxb-6604068403200-6574656633318-gp1xLPOotUucCascW2KID6FD")
	os.Setenv("CHANNEL_ID","C06H5MRP944")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"DevOps_Fresher_Resume.pdf"}

for i:= 0; i<len(fileArr); i++{
	parameter := slack.FileUploadParameters{
		Channels: channelArr, 
		File: fileArr[i],
	}
	 file, err := api.UploadFile(parameter)
		if  err !=nil{
		fmt.Printf("Error in uploading file: %s",err )
		return 
	}
	fmt.Printf("Name:%s , URL: %s\n",file.Name, file.URL)

	}
	
}
