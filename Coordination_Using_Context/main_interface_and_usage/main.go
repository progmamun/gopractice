package main

type Options struct {
	Contents bool
	Exclude []string
	}
	func FileSearch(ctx context.Context, root, term string, o *Options)

	type Result struct {
		Err error
		File string
		Matches []Match
		}
		type Match struct {
		Line int
		Text string
		}

		for r := range FileSearch(ctx, directory, searchTerm, options) {
			if r.Err != nil {
				fmt.Printf("%s - error: %s\n", r.File, r.Err)
				continue
			}
			if !options.Contents {
				fmt.Printf("%s - match\n", r.File)
				continue
			}
			fmt.Printf("%s - matches:\n", r.File)
			for _, m := range r.Matches {
				fmt.Printf("\t%d:%s\n", m.Line, m.Text)
			}
		}