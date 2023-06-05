package main

import (
	"bufio"
	"fmt"
	"github.com/pion/webrtc/v3"
	"log"
	"os"
	"strconv"
	"time"
	"vediomeeting/internal/helper"
)

func main() {
	// 1. 创建一个peer connection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := peerConnection.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	// 2. 创建data channel
	dataChannel, err := peerConnection.CreateDataChannel("foo", nil)
	dataChannel.OnOpen(func() {
		log.Println("data channel has opened")
		i := -1000
		for range time.NewTicker(time.Second * 5).C {
			if err := dataChannel.SendText("hello world" + strconv.Itoa(i)); err != nil {
				log.Println(err)
			}
		}
	})
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Println(string(msg.Data))
	})
	// 3. 创建offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Println(err)
	}
	// 4. 设置local description
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		log.Println(err)
		return
	}
	// 5. print offer
	println("OFFER:")
	print(helper.Encode(offer))
	println()
	// 6. input answer
	var answer webrtc.SessionDescription
	println("Please input ANSWER: ")
	answerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	helper.Decode(answerStr, &answer)
	// 7. 设置remote description
	if err := peerConnection.SetRemoteDescription(answer); err != nil {
		log.Println(err)
	}
	select {}
}
