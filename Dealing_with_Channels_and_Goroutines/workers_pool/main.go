package main


type result struct {
	Loc [2]float64
	Result []maps.Result
	Error error
	}
	keys := strings.Split(os.Getenv("MAPS_APIKEYS"), ",")
requests := make(chan [2]float64, len(keys))
results := make(chan result, len(keys))
done := make(chan struct{})
for i := range keys {
	go func(id int) {
	log.Printf("Starting worker %d with API key %q", id, keys[id])
	client := maps.NewClient(maps.DailyCap, keys[id])
	for j := range requests {
	var r = result{Loc: j}
	log.Printf("w[%d] working on %v", id, j)
	r.Result, r.Error = client.ReverseGeocode(j[0], j[1])
	results <- r
	}
	done <- struct{}{}
	}(i)
	}

	go func() {
		for _, l := range [][2]float64{
		{40.4216448, -3.6904040},
		{40.4163111, -3.7047328},
		{40.4123388, -3.7096724},
		{40.4145150, -3.7064412},
		} {
		requests <- l
		}
		close(requests)
		}()
		go func() {
			count := 0
			for range done {
			if count++; count == len(keys) {
			break
			}
			}
			close(results)
			}()
			
			for r := range results {
				log.Printf("received %v", r)
				}