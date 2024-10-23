package proxy

import (
	"github.com/spf13/cobra"
	"github.com/syzhang42/go-fire/auth"
	"github.com/syzhang42/go-fire/log"
	"github.com/syzhang42/hermes/utils/ver"
)

var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "general hermes service",
}

func Execute() {
	auth.Must(rootCmd.Execute())
}

func info() {
	log.Info("==> GitTag   :", ver.GitTag)
	log.Info("==> GitCommit:", ver.CommitLog)
	log.Info("==> GoVersion:", ver.GoVersion)
	log.Info("==> Author   :", ver.Author)
	log.Infof("==> hermes(APP:%v) born in %v <==\n", ver.Version, ver.BuildTime)
	log.Info("")
}
