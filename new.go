package progressbar

import "time"

func Add(maxBytes int64, title string, opts ...Opt) MultiPB {
	defaultMPB.Add(maxBytes, title, opts...)
	return defaultMPB
}

// New creates a managed MultiPB progressbar object so you
// can setup the properties of the bar.
//
//	bar := progressbar.New()
//	bar.Add(
//		resp.ContentLength,
//		"downloading go1.14.2.src.tar.gz",
//		// progressbar.WithSpinner(14),
//		// progressbar.WithStepper(3),
//		progressbar.WithBarStepper(0),
//	)
//
// A MultiPB or PB progressbar object is a writable object
// which can receive the data writing via Writer interface:
//
//	f, _ := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY, 0o644)
//	_, _ = io.Copy(io.MultiWriter(f, bar), resp.Body)
//	f.Close()
//	bar.Close()
//
// The MultiPB object can be added into Tasks container. For
// more information to see NewTasks() and NewDownloadTasks().
func New(opts ...MOpt) MultiPB {
	bar := multiBar(opts...)
	return bar
}

type SchemaData struct {
	Indent  string
	Prepend string
	Bar     string
	Percent string
	Title   string
	Current string
	Total   string
	Elapsed string
	Speed   string
	Append  string

	PercentFloat float64
	ElapsedTime  time.Duration
}
