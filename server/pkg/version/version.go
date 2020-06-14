package version

import (
	"fmt"
	"runtime"
)

// Info 表示版本相关信息
type Info struct {
	GitTag string `json:"git_tag"` // git tag
	GitCommit string `json:"git_commit"` // git commit
	GitTreeState string `json:"git_tree_state"`
	BuildDate string `json:"build_date"` // 构建日期
	GoVersion string `json:"go_version"`
	Compiler string `json:"compiler"`
	Platform string `json:"platform"`
}

// String 返回版本字符
func (info *Info) String() string {
	return info.GitTag
}

// Get 返回详细的版本信息
func Get() Info {
	return Info{
		GitTag: gitTag,
		GitCommit: gitCommit,
		GitTreeState: gitTreeState,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler: runtime.Compiler,
		Platform: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
