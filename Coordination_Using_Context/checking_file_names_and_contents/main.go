package main

if o == nil || !o.Contents {
	if name == term {
	select {
	case <-ctx.Done():
	default:
	ch <- Result{File: file}
	}
	}
	return
	}

	f, err := os.Open(file)
if err != nil {
select {
case <-ctx.Done():
default:
ch <- Result{File: file, Err: err}
}
return
}
defer f.Close()

scanner, matches, line := bufio.NewScanner(f), []Match{}, 1
for scanner.Scan() {
select {
case <-ctx.Done():
break
default:
if text := scanner.Text(); strings.Contains(text, term) {
matches = append(matches, Match{Line: line, Text: text})
}
line++
}
}

select {
case <-ctx.Done():
break
default:
if err := scanner.Err(); err != nil {
ch <- Result{File: file, Err: err}
return
}
if len(matches) != 0 {
ch <- Result{File: file, Matches: matches}
}
}