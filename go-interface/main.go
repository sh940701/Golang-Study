package main

import (
	fedex "go-interface/modules/fedex"
	koreaPost "go-interface/modules/koreaPost"
)

// fedex와 koreaPost 모두 내부적으로 Send 함수를 가지고 있기 때문에
// Sender 인터페이스로 사용 가능하다.
type Sender interface {
	Send(parcel string)
}

// 각 koreaPost와 fedex 객체는 SendBook 함수로 넘어오면서
// Sender 타입으로 wrapping된다. -> 각각의 객체 타입의 Send함수를 호출한다.
func SendBook(name string, sender Sender)  {
	sender.Send(name)
}

func main()  {
	sender0 := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender0)
	SendBook("그리스인 조르바", sender0)


	sender1 := &fedex.FedexSender{}
	SendBook("궁금한 이야기 why", sender1)
	SendBook("궁금한 이야기 who", sender1)
}