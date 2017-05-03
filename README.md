使用GoDep
---
* 安装godep
```sh
go get github.com/tools/godep
```

* Godep需要golang.org/x/sys/unix包
```sh
go get golang.org/x/sys/unix
```

* 进入程序目录，首先需要使用go get获取相关依赖
```sh
cd /Volumes/Share/Go_program/src/cookeem.com
go get gopkg.in/gin-gonic/gin.v1
godep save -v ./...
```

* 成功后会自动创建./Godeps目录，下边有Godeps.json配置文件
```sh
ls -l Godeps/
total 16
-rw-r--r--  1 cookeem  staff  1361  4 12 10:57 Godeps.json
-rw-r--r--  1 cookeem  staff   136  4 12 10:57 Readme

cat Godeps/Godeps.json
{
	"ImportPath": "cookeem.com",
	"GoVersion": "go1.8",
	"GodepVersion": "v79",
	"Packages": [
		"./..."
	],
	"Deps": [
		{
			"ImportPath": "github.com/gin-gonic/gin/binding",
			"Comment": "v1.0-2-g3900df0",
			"Rev": "3900df04d2a88e22beaf6a2970c63648b9e1b0e1"
		},
		{
			"ImportPath": "github.com/gin-gonic/gin/render",
			"Comment": "v1.0-2-g3900df0",
			"Rev": "3900df04d2a88e22beaf6a2970c63648b9e1b0e1"
		},
		{
			"ImportPath": "github.com/golang/protobuf/proto",
			"Rev": "98fa357170587e470c5f27d3c3ea0947b71eb455"
		},
		{
			"ImportPath": "github.com/manucorporat/sse",
			"Rev": "ee05b128a739a0fb76c7ebd3ae4810c1de808d6d"
		},
		{
			"ImportPath": "github.com/mattn/go-isatty",
			"Comment": "v0.0.2",
			"Rev": "fc9e8d8ef48496124e79ae0df75490096eccf6fe"
		},
		{
			"ImportPath": "golang.org/x/net/context",
			"Rev": "8b4af36cd21a1f85a7484b49feb7c79363106d8e"
		},
		{
			"ImportPath": "golang.org/x/sys/unix",
			"Rev": "f3918c30c5c2cb527c0b071a27c35120a6c0719a"
		},
		{
			"ImportPath": "gopkg.in/gin-gonic/gin.v1",
			"Comment": "v1.1.4",
			"Rev": "e2212d40c62a98b388a5eb48ecbdcf88534688ba"
		},
		{
			"ImportPath": "gopkg.in/go-playground/validator.v8",
			"Comment": "v8.18.1",
			"Rev": "5f57d2222ad794d0dffb07e664ea05e2ee07d60c"
		},
		{
			"ImportPath": "gopkg.in/yaml.v2",
			"Rev": "a5b47d31c556af34a302ce5d659e6fea44d90de0"
		}
	]
}
```

* 成功后会自动创建./vendor，并且把依赖包复制到本目录下
```sh
ls -l vendor/
total 0
drwxr-xr-x  6 cookeem  staff  204  4 12 10:51 github.com
drwxr-xr-x  3 cookeem  staff  102  4 12 10:51 golang.org
drwxr-xr-x  5 cookeem  staff  170  4 12 10:51 gopkg.in
```

* 运行go程序
```sh
godep go run gin/gin_demo.go
```

* 编译go程序
```sh
godep go build gin/gin_demo.go
```


在Linux的Jenkins下进行godep编译
---
* 在jenkins下，需要把vendor目录修改为src目录，在jenkins下执行以下构建命令：
```sh
export GOPATH=`pwd`
rm -rf src
mv vendor src
go build gin/gin_demo.go
```


GitLab与Jenkins集成
---

### Jenkins与GitLab互联：
- Jenkins GitLab Plugin、Jenkins Git Plugin、Gitlab Hook Plugin（用于gitlab的push触发自动构建）
- 在Jenkins的"系统管理" -》"系统设置" -》"Gitlab"中设置Connection name、Gitlab host URL、Credentials
- 其中Credentials使用Gitlab API Token，打开Gitlab的"User Settings" -》"Account" -》 "Private token"
- 把Gitlab的"Private token"粘贴到Jenkins的Gitlab设置的Credentials，然后验证测试

### Jenkins新建项目，实现gitlab push自动构建：
- 新建"构建一个自由风格的软件项目"
- "General" -》"	GitLab connection"，选择对应的gitlab（配置位于Jenkins的"系统管理" -》"系统设置" -》"Gitlab"）
- "构建触发器" -》"Build when a change is pushed to GitLab. GitLab CI Service URL: http://localhost:8080/project/godemo_gitlab" -》"高级" -》"Secret token" -》"Generate"
- （该操作在GitLab中进行）选择godemo项目 -》Settings -》Integrations
URL输入：http://jenkins:8080/project/godemo_gitlab（对应Jenkins"构建触发器"部分提示的URL）
Secret Token：（对应Jenkins"构建触发器"部分自动Generate的"Secret token"）
Enable SSL verification：必须取消
- "源码管理" -》 "Git" -》"Repository URL"：http://gitlab/cookeem/godemo
- "源码管理" -》 "Git" -》"Credentials"：选择对应的密钥（配置位于Jenkins的"系统管理" -》"系统设置" -》"Gitlab"）
- "构建环境" -》 "Set up Go programming language tools" -》 "Go version"：选择对应的版本（配置位于Jenkins的"系统管理" -》"Global Tool Configuration" -》"Go"）
- "构建" -》 "Execute shell"，内容为：
```sh
pwd
echo "###################"
printenv
echo "###################"
ls -al
echo "###################"
export GOPATH=`pwd`
rm -rf src
mv vendor src
echo "###################"
go build gin/gin_demo.go
```

