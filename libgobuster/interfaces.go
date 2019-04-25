package libgobuster

// GobusterPlugin is an interface which plugins must implement
type GobusterPlugin interface {
	PreRun() error
	Run(string) ([]Result, error)
	ResultToString(Result) (string, error)
	GetConfigString() (string, error)
}

// GobusterPluginOption is an interface which plugin options must implement
type GobusterPluginOption interface {
	Sanitize()
}
