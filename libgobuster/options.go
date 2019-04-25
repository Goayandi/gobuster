package libgobuster

// Options helds all options that can be passed to libgobuster
type Options struct {
	Threads        int
	Wordlist       string
	OutputFilename string
	NoStatus       bool
	NoProgress     bool
	Quiet          bool
	WildcardForced bool
	Verbose        bool
}
