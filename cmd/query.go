package cmd

import (
	"script-exporter/global"

	"github.com/spf13/cobra"
)

var cfg string

var queryCmd = &cobra.Command{
	//二级子命令
	Use: "query",
	//别名
	Aliases: []string{"check"},
	//简短提示
	Short: "short query",
	//详细提示
	Long: "long query",
	Run: func(cmd *cobra.Command, args []string) {
		global.Respath = cfg
	},
}

// 参数检查
// var curArgsCheckCmd = &cobra.Command{
// 	Use: "cuscheck",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		global.Respath = args[0]
// 	},
// 	Args: func(cmd *cobra.Command, args []string) error {
// 		if len(args) != 1 {
// 			return errors.New("请输入脚本执行结果的路径!!!")
// 		}
// 		return nil
// 	},
// }

func init() {
	rootCmd.AddCommand(queryCmd)
	// rootCmd.AddCommand(curArgsCheckCmd)
	// 绑定命令行输入，绑定一个参数
	// 参数分别表示，绑定的变量，参数长名(--str)，参数短名(-s)，默认内容，帮助信息
	queryCmd.Flags().StringVarP(&cfg, "config", "c", "./res.yaml", "请选择配置文件!!!")
}
