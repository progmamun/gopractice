package main

if info.IsDir() {
	files, err := ioutil.ReadDir(file)
	if err != nil {
	select {
	case <-ctx.Done():
	return
	default:
	ch <- Result{File: file, Err: err}
	}
	return
	}
	select {
	case <-ctx.Done():
	default:
	wg.Add(len(files))
	for _, f := range files {
	go fileSearch(ctx, ch, wg, filepath.Join(file,
	f.Name()), term, o)
	}
	}
	return
	}
	