package main

// ReadingService adapts ReadingList for RPC
type ReadingService struct {
	ReadingList
}

// sets the success pointer value from error
func setSuccess(err error, b *bool) error {
	*b = err == nil
	return err
}
func (r *ReadingService) AddBook(b Book, success *bool) error {
	return setSuccess(r.ReadingList.AddBook(b), success)
}
func (r *ReadingService) RemoveBook(isbn string, success *bool) error {
	return setSuccess(r.ReadingList.RemoveBook(isbn), success)
}
func (r *ReadingService) GetProgress(isbn string, pages *int) (err error) {
	*pages, err = r.ReadingList.GetProgress(isbn)
	return err
}

type Progress struct {
	ISBN  string
	Pages int
}

func (r *ReadingService) SetProgress(p Progress, success *bool) error {
	return setSuccess(r.ReadingList.SetProgress(p.ISBN, p.Pages), success)
}
func (r *ReadingService) AdvanceProgress(p Progress, success *bool) error {
	return setSuccess(r.ReadingList.AdvanceProgress(p.ISBN, p.Pages), success)
}
if len(os.Args) != 2 {
	log.Fatalln("Please specify an address.")
	}
	if err := rpc.Register(&common.ReadingService{}); err != nil {
	log.Fatalln(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", os.Args[1])
	if err != nil {
	log.Fatalln(err)
	}
	log.Println("Server Started")
	if err := http.Serve(l, nil); err != nil {
	log.Fatal(err)
	}