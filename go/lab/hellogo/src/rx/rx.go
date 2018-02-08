// rx.go
package main

import (
	"fmt"

	// "github.com/reactivex/rxgo"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func HelloRxGo() {
	fmt.Println("Hello Rx Go!")

	sub := observable.Just(1, 2.0, "hello").Subscribe(observer.Observer{
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing %v\n", item)
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	})
	<-sub
	fmt.Println("After <-sub")
}

func main() {
	HelloRxGo()
}
