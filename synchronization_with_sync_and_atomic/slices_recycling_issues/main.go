package main

var (
	list = make([][]byte, 20)
	m sync.Mutex
	)
	for i := 0; i < 20; i++ {
	go func(v int) {
	time.Sleep(time.Second * time.Duration(1+v/4))
	b := Get()
	defer func() {
	Put(b)
	wg.Done()
	}()
	fmt.Fprintf(b, "Goroutine %2d using %p, after %.0fs\n", v, b, time.Since(start).Seconds())
	m.Lock()
	list[v] = b.Bytes()
	m.Unlock()
	}(i)
	}
	wg.Wait()

	for i := range list {
		fmt.Printf("%d - %s", i, list[i])
		}

//solution
m.Lock()
list[v] = make([]byte, b.Len())

copy(list[v], b.Bytes())
m.Unlock()