<p>Scanner is part of the bufio package. It is useful for stepping through files at specific delimiters. Commonly, the newline character is used as the delimiter to break up a file by lines. In a CSV file, commas would be the delimiter. The os.File can be wrapped in a bufio.Scanner just like a buffered reader. We call Scan() to read up to the next delimiter, and then use Text() or Bytes() to get the data that was read.

The delimiter is not just a simple byte or character. There is actually a special function you have to implement that will determine where the next delimiter is, how far forward to advance the pointer, and what data to return. If no custom SplitFunc is provided, it defaults to ScanLines which will split at every newline character. Other split functions included in bufio are ScanRunes, and ScanWords.</p>

<p>// To define your own split function, match this fingerprint
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// Returning (0, nil, nil) will tell the scanner
// to scan again, but with a bigger buffer because
// it wasn't enough data to reach the delimiter

In the next example, a bufio.Scanner is created from the file, and then we scan read the file word by word.</p>
