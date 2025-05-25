package Part5BufferedChannel
import (

"fmt"
"time"
)

func Producer(ch chan int)  {
 	for i := 0; i < 10; i++ {
 	 	fmt.Println("producer sending", i)
		  ch <- i
		  time.Sleep(time.Millisecond*100)
 	}
	 close(ch)
}

func Consumer(id int,ch chan int)  {
	for value := range ch {
		fmt.Printf("Consumer %d: received %d\n",id, value)
		time.Sleep(time.Second*1)	//	slow process
	}
	fmt.Printf("Consumer %d: done\n",id)
}