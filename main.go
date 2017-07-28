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
	"strconv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				   //kg, err := strconv.Atoi(message.Text) * 
   				   //if err != nil {
      				// handle error
   				   //}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(weight(message.Text))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func weight(x string) string{

    st :=x
    s := string(st[0])
    var aaa string
    if _, err := strconv.Atoi(st[1:len(st)]); err == nil {
        value, err := strconv.ParseFloat(st[1:len(st)], 64)
        if err != nil {
        // do something sensible
        }
        amount := float64(value)
        switch s {
        case string('p'):
        aaa=weightLbs(amount)+" Lbs"
        //return weightLbs(amount)+" Lbs"
        //fmt.Println(weightLbs(amount)+" Lbs")
        case string('k'):
        aaa=eightKg(amount)+" Kg"	
        //return weightKg(amount)+" Kg"	
        //fmt.Println(weightKg(amount)+" Kg")
        }
    }else{
        aaa="invalid input"
    }
    return aaa
}


func weightLbs(x float64) string{
    return strconv.FormatFloat(x * 2.2, 'f', 2, 64)
}
func weightKg(x float64) string{
    return strconv.FormatFloat(x * 0.45, 'f', 2, 64)
}