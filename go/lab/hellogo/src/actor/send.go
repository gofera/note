package main

import (
	"fmt"
)

type Message interface{}

type Mail struct {
	Msg  Message
	From ActorRef
	To   ActorRef
}

type MailHandler func(Mail)

type actor struct {
	mailbox chan Mail
}

func (actor *actor) OnReceived(handle MailHandler) {
	go func() {
		for mail := range actor.mailbox {
			handle(mail)
		}
	}()
}

func (actor *actor) Close() {
	close(actor.mailbox)
}

type ActorRef struct {
	Id string
	*actor
}

func NewActor(id string, mailboxSize uint) ActorRef {
	act := actor{make(chan Mail, mailboxSize)}
	return ActorRef{id, &act}
}

func Send(mail Mail) {
	mail.To.actor.mailbox <- mail
}

func main() {
	fmt.Println("hi")
	exit := make(chan bool)
	wenzhe := NewActor("wenzhe", 10)
	qiqi := NewActor("qiqi", 10)

	wenzhe.OnReceived(func(mail Mail) {
		self := mail.To
		fmt.Println(mail.From.Id, "say", mail.Msg, "to", mail.To.Id)
		Send(Mail{"baby baby", self, qiqi})
		Send(Mail{"Good night baby", self, qiqi})
	})
	qiqi.OnReceived(func(mail Mail) {
		msg := mail.Msg
		fmt.Println(mail.From.Id, "sing a song", msg, "to", mail.To.Id)
		if msg == "Good night baby" {
			exit <- true
		}
	})
	Send(Mail{"hello", qiqi, wenzhe})
	Send(Mail{"Do Rei Mi", wenzhe, qiqi})
	// qiqi.Close()
	<-exit
}
