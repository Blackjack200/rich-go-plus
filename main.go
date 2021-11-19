package main

import (
	"github.com/blackjack200/rich-go-plus/client"
	"github.com/blackjack200/rich-go-plus/codec"
	"time"
)

func main() {
	c, err := client.Dial("909326302757126186")
	if err != nil {
		panic(err)
	}
	println("success")
	do(err, c)
	do(err, c)
}

func do(err error, c *client.Client) {
	err = c.SetActivity(&codec.Activity{
		Details:    "Details",
		State:      "State",
		LargeImage: "apple",
		LargeText:  "fkk",
		SmallImage: "ff",
		SmallText:  "st",
		Party:      nil,
		Timestamps: nil,
		Secrets:    nil,
		Buttons: []*codec.Button{
			{
				Label: "YES",
				Url:   "https://baidu.com",
			}, {
				Label: "NO",
				Url:   "https://baidu.com",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	msg, suc := c.Read()
	println(msg.Payload)
	println(suc)
	time.Sleep(time.Second * 15)
	c.Close()
	println("success")
}
