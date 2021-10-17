package main


/*wg := sync.WaitGroup{}
merge := make(chan map[string]int)
wg.Add(len(src))
go func() {
wg.Wait()
close(merge)
}()*/

for _, ch := range src {
	go func(ch <-chan map[string]int) {
	defer wg.Done()
	for v := range ch {
	select {
	case <-ctx.Done():
	return
	case merge <- v:
	}
	}
	}(ch)
	}
	count := make(map[string]int)
	for {
	select {
	case <-ctx.Done():
	return count
	case c, ok := <-merge:
	if !ok {
	return count
	}
	for k, v := range c {
	count[k] += v
	}
	}
	}
	
	func main() {
		ctx, canc := context.WithCancel(context.Background())
		defer canc()
		src := SourceLineWords(ctx, ioutil.NopCloser(strings.NewReader(cantoUno)))
		count1, count2 := WordOccurrence(ctx, src), WordOccurrence(ctx, src)
		final := MergeCounts(ctx, count1, count2)
		fmt.Println(final)
		}