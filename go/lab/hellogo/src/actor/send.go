package main

import (
	"fmt"
)

type Message interface{}

type Mail struct {
	Msg  Message
	From Actor
	To   Actor
}

type MailHandler func(msg Message, from Actor, self Actor)

type Path string

type Actor struct {
	path    Path
	mailbox chan Mail
}

func (actor *Actor) OnReceived(handle MailHandler) {
	go func() {
		for mail := range actor.mailbox {
			handle(mail.Msg, mail.From, *actor)
		}
	}()
}

func (actor *Actor) Close() {
	close(actor.mailbox)
}

func (actor *Actor) GetPath() Path {
	return actor.path
}

func (from *Actor) Tell(to Actor, msg Message) {
	to.mailbox <- Mail{msg, *from, to}
}

func (from *Actor) Ask(to Actor, msg Message) Message {
	waitForReply := make(chan Mail)
	(&Actor{from.path, waitForReply}).Tell(to, msg)
	reply := <-waitForReply
	return reply.Msg
}

func NewActor(path Path, mailboxSize uint) Actor {
	return Actor{path, make(chan Mail, mailboxSize)}
}

func main() {
	fmt.Println("hi")
	exit := make(chan bool)
	wenzhe := NewActor("wenzhe", 10)
	qiqi := NewActor("qiqi", 10)

	wenzhe.OnReceived(func(msg Message, from Actor, self Actor) {
		fmt.Println(from.GetPath(), "say", msg, "to", self.GetPath())
		self.Tell(qiqi, "ABCDEFG")
		self.Tell(qiqi, "Love you baby")
		reply := self.Ask(qiqi, "baby baby")
		fmt.Println("Get reply:", reply)
		self.Tell(qiqi, "Good night baby")
	})
	qiqi.OnReceived(func(msg Message, from Actor, self Actor) {
		fmt.Println(from.GetPath(), "sing a song", msg, "to", self.GetPath())
		switch msg {
		case "baby baby":
			self.Tell(from, "dadi dadi")
		case "Good night baby":
			exit <- true
		}
	})
	qiqi.Tell(wenzhe, "hello")
	wenzhe.Tell(qiqi, "Do Rei Mi")
	// qiqi.Close()
	<-exit
}
