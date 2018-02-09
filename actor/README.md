# Actor发送消息
A -- send msg --> B

```
func (self *ActorContext) Send(source *ActorRef, target *ActorRef, msg interface{}) {
  self
}
```