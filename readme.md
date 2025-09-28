## Golang 运维工具 Web 管理界面

3️⃣ 初始化项目

新建项目目录，例如：

mkdir myops
cd myops

初始化 Go Modules：

go mod init myops

安装依赖：

go get github.com/gin-gonic/gin
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/mem
go get github.com/shirou/gopsutil/disk

4️⃣ 项目结构示例
myops/
│── go.mod
│── go.sum
│── main.go
└── templates/
└── index.html

main.go：你的 Golang 运维工具代码

templates/index.html：Web 界面模板

5️⃣ 运行项目
go run main.go

2️⃣ 永久配置国内镜像

在命令行输入：

go env -w GOPROXY=https://goproxy.cn,direct

3️⃣ 其他常用国内 Go 镜像

阿里云：https://mirrors.aliyun.com/goproxy/

七牛：https://goproxy.cn（推荐）

设置方法类似：

go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
