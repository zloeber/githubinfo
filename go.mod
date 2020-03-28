module github.com/zloeber/githubinfo

go 1.13

require (
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.2
	github.com/zloeber/githubinfo/cmd/githubinfo v0.1.1
)

replace github.com/zloeber/githubinfo/cmd => ./cmd
