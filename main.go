// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	//var err2 error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	//events, err2 := strings.Contains(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	
	bot, err := linebot.New(<channel secret>, <channel token>)
if err != nil {
    ...
}
res, err := bot.GetUserProfile(<userId>).Do();
if err != nil {
    ...
}
println(res.Displayname)
println(res.PicutureURL)
println(res.StatusMessage)
	
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				//----------------回聲範例---------------------
				/*if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text/*+" TextOK!")).Do(); err != nil {
					//發送訊息的格式
					log.Print(err)
				}*/
				//----------------------------------------------------------------------
				
				//----------------關鍵字回復--------------------
				if strings.Contains(message.Text, "Hi") || strings.Contains(message.Text, "HI")|| strings.Contains(message.Text, "hi")|| strings.Contains(message.Text, "hI") {
					out := fmt.Sprintf("HI~~~~")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} 
				 else if strings.Contains(message.Text, "初音") {
					originalContentURL := "https://s65.youmaker.com/flv/2014/2-14/mp4909826563a3583eddc78f4a5da9cf3bbdeffecf79065.mp4"
					previewImageURL := "https://upload.wikimedia.org/wikipedia/zh/0/00/Miku_Hatsune.png"
					bot.ReplyMessage(event.ReplyToken, linebot.NewVideoMessage(originalContentURL, previewImageURL)).Do()
				} 

}
