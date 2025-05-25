package Part6Select_timeAfter
import (

"fmt"
"time"
)

func Producer(ch chan int)  {
	for i := 1; i <5 ; i++ {
	 	fmt.Println("producer sending", i)
		ch<-i
		time.Sleep(time.Duration(i)*time.Second)
	}
	close(ch)
}

func Consumer(ch chan int)  {
	for  {
	 	select{
		 case val, ok := <-ch:
			 if !ok {
				 fmt.Println("channel closed")
				 return
			 }
			 fmt.Println("receive value", val)

		 case <-time.After(time.Second * 3):
			 fmt.Println("time out")
	 	}
	}
}