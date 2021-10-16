package main

ch, wg, once := make(chan Result), sync.WaitGroup{}, sync.Once{}
go func() {
wg.Wait()
fmt.Println("* Search done *")
once.Do(func() {
close(ch)
})
}()
go func() {
<-ctx.Done()
fmt.Println("* Context done *")
once.Do(func() {
close(ch)
})
}()