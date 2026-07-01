package build

// Populated at link time via -ldflags -X.
var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)
