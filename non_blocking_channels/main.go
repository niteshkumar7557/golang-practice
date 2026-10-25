package main 

import (
	"fmt"
)


func producer(ch chan<- int){
	for i := range 5{
		ch <- i
	}
	close(ch)
}
func filter(in <-chan int, out chan<- int){
	for i := range in{
		if i%2 == 0{
			out <- i
		}
	}
	close(out)
}

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2{
		fmt.Println(val)
	}
}



// func main(){
// 	data := make(chan int)
// 	quit := make(chan bool)

// 	go func(){
// 		for {
// 			select{
// 			case d := <-data:
// 				fmt.Println("Recieved Data", d)
// 			case <- quit:
// 				fmt.Println("Stopping...")
// 			default:
// 				fmt.Println("Waiting for the data...")
// 				time.Sleep(500 * time.Millisecond)
// 			}
// 		}
// 	}()


// 	for i := range 5 {
// 		data <- i
// 		time.Sleep(time.Second)
// 	}
// 	quit <- true
// }
