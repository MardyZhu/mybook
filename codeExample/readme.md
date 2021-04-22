- cobra：眼镜蛇
- cobra official website：https://cobra.dev/
- install: import "github.com/spf13/cobra"

# 学习Cobra第一步：
### 一、Cobra基础组成结构
Cobra is built on a structure of commands, arguments & flags.
Commands represent actions, Args are things and Flags are modifiers for those actions.
- modifiers 修饰符
- flags 标志提供修饰符以控制命令的操作方式

### 二、开始实践
当我们需要手动接入Cobra的时候，需要创建main.go 和 rootCmd文件：
    1. main.go 文件的目的只有一个：初始化Cobra
    2. rootCmd 的目的是：代表没有调用子命令的基础命令

### 三、和Cobra配套的其他包
#### 1、viper：https://github.com/spf13/viper  
    - viper：毒蛇  
    - 中文教程：https://www.liwenzhou.com/posts/Go/viper_tutorial/
    
### 四、常见问题
#### 1、如何设置必须标志
    标志默认是可选的。如果你想在缺少标志时命令报错，请设置该标志为必需：
```
var region string

rootCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")
```