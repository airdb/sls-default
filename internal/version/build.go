package version

import (
	"encoding/json"
	"os"
	"runtime"
)

// Build version info.
type BuildInfo struct {
	GoVersion string
	Env       string
	Version   string
	Build     string
	BuildTime string
}

var (
	Version   string
	Build     string
	BuildTime string
)

func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		GoVersion: runtime.Version(),
		Env:       os.Getenv("ENV"),
		Version:   Version,
		Build:     Build,
		BuildTime: BuildTime,
	}
}

func (info *BuildInfo) ToString() string {
	out, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	return string(out)
}
