package main

for condition {
	mu.Lock()
	defer mu.Unlock()
	action()
	}
	for condition {
		func() {
		mu.Lock()
		defer mu.Unlock()
		action()
		}()
		}
		
		for condition {
			mu.Lock()
			action()
			mu.Unlock()
			}