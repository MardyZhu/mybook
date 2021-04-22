package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 这个文件里是创建了一个cobra的基础命令

// 配置文件
var cfgFile string
var projectBase string
var userLicense string

// rootCmd 代表没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   "root_cmd",
	Short: "root cmd",
	Long:  "This is a progress about root_cmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root")
	},
}

// Execute 将所有子命令添加到root命令，并适当设置标志
// 有main.main() 调用。它只需要对rootCmd调用一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// 在init()函数中定义标志并处理配置
func init() {
	cobra.OnInitialize(initConfig)
	// 设置flags的value
	// flags 标志提供修饰符以控制命令的操作方式
	// -- persistent 持久的

	// PersistentFlags() 的作用是：将该标志用于分配给它的命令以及该命令下的每个命令
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is empty.)")
	// projectBase 是全局命令，该操作意味着将标志分配为根上的持久标志
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectBase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "yori-MyName", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	// 用viper绑定arguments
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectBase", rootCmd.PersistentFlags().Lookup("projectBase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	viper.SetDefault("author", "Name-yori")
	viper.SetDefault("license", "test-lic")

}

// 初始化配置信息
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	}
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("Can't read config:", err)
	// 	os.Exit(1)
	// }
}
