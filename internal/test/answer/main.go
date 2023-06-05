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
	// 2. on data channel
	peerConnection.OnDataChannel(func(dataChannel *webrtc.DataChannel) {
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
	})
	// 3. input offer
	var offer webrtc.SessionDescription
	println("please input OFFER: ")
	offerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	helper.Decode(offerStr, &offer)
	// 4. 设置remote description
	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}
	// 5. create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}
	// 6. 设置local description
	if err := peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}
	// 7. gather complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)
	<-gatherComplete
	// 8. print answer
	fmt.Println("ANSWER: ")
	fmt.Print(helper.Encode(peerConnection.LocalDescription()))
	fmt.Println()
	select {}
}
