package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// 实现后台启动go服务
// 执行命令启动服务 bgrun start
// 执行命令关闭服务 bgrun stop

// 先写一个简单的go服务：start
func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello yori!")
}

func service() {
	http.HandleFunc("/", sayhello)           // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// 先写一个简单的go服务：end
// -----------------------------------------------------------------------------------------------------------------------------------------------

// 用cobra实现go服务后台运行
// 启动服务的命令：bgrun start
var daemon bool // daemon 守护线程
var bgRunStartCmd = &cobra.Command{
	Use:  "bgrun_start", // 使用命令
	Long: "bg run pg start",
	Run: func(cmd *cobra.Command, args []string) {
		if daemon {
			fmt.Println("bgrun strat  daemon= ", daemon)
			// 这里的第一个参数name是服务的名称，第二个参数是要执行的命令，
			// 第一个参数+第二个参数=本次要执行的命令，比如这次传入的参数实际执行的命令是：./cobra bgrun_start
			command := exec.Command("./cobra", "bgrun_start")
			command.Start()
			fmt.Printf("gonne start, [PID] %d running...\n", command.Process.Pid)
			ioutil.WriteFile("gonne.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
			daemon = false
			os.Exit(0) // 实现后台运行全靠它
		} else {
			fmt.Println("no bgrun start")
		}
		service()
	},
}

// 关闭服务的命令 ：bgrun stop
var bgRunStopCmd = &cobra.Command{
	Use:  "bgrun_stop", // 使用命令
	Long: "bg run pg stop",
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile("gonne.lock")
		command := exec.Command("kill", string(strb))
		command.Start()
		fmt.Println("server is close")
	},
}

func init() {
	// 第一个参数取值，第二个参数代码--deamon，第三个参数代表-d,第四个参数代码不加-d时候的默认值，第五参数是描述
	bgRunStartCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "is daemon?")
	rootCmd.AddCommand(bgRunStartCmd)
	rootCmd.AddCommand(bgRunStopCmd)

}
