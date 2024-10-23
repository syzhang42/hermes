package proxy

import (
	"github.com/spf13/cobra"
	"github.com/syzhang42/hermes/server"
	"github.com/syzhang42/hermes/utils/ormx"
	"github.com/syzhang42/hermes/utils/ver"
)

func init() {
	restCmd := &cobra.Command{
		Use:   "proxy",
		Short: "hermes proxy server",
		Run: func(restCmd *cobra.Command, args []string) {
			info()
			ver.Version, _ = restCmd.Flags().GetString("version")
			ver.CfgPath, _ = restCmd.Flags().GetString("cfgPath")
			//1、初始化db
			ormx.Init(ver.CfgPath)
			//2、启动服务
			server.Run(ver.CfgPath)
		},
	}

	rootCmd.AddCommand(restCmd)
	restCmd.Flags().StringP("version", "v", ver.Version, "proxy config version")
	restCmd.Flags().StringP("cfgPath", "c", ver.CfgPath, "proxy config path")
}
