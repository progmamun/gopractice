package main

func main() {
	var a chan int
	close(a)
}

/*func main() {
	var a chan int
	a <- 1
}
func main() {
	var a chan int
	<-a
}
*/
