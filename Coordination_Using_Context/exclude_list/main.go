package main

defer wg.Done()
_, name := filepath.Split(file)
if o != nil {
for _, e := range o.Exclude {
if e == name {
return
}
}
}

info, err := os.Stat(file)
if err != nil {
select {
case <-ctx.Done():
return
default:
ch <- Result{File: file, Err: err}
}
return
}
