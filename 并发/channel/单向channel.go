package main

func main() {
	//创建一个channel，默认双向
	ch := make(chan int)

	//双向channel能隐式转化为单向channel
	var writeCh chan<- int = ch //只能写，不能读
	writeCh <- 666

	var readCh <-chan int = ch //只能读不能写
	<-readCh

	//单向不能转化为双向！

}
