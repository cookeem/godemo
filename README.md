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

* 进入程序目录，使用godep save保存相关依赖到Godeps
```sh
cd /Volumes/Share/Go_program/src/cookeem.com
godep save -v ./...
```

* 成功后会自动创建./Godeps目录，下边有Godeps.json配置文件，记录相关依赖的以及依赖版本；同时会自动创建./vendor，并且把依赖包复制到本目录下

* 进入程序目录，使用godep restore获取相关依赖
```sh
cd /Volumes/Share/Go_program/src/cookeem.com
godep restore -v
```

* 运行go程序
```sh
godep go run gin/gin_demo.go
```

* 编译go程序
```sh
godep go build gin/gin_demo.go
```


Jenkins与GitLab、Docker的集成
---

### Jenkins与GitLab、docker互联：
- CloudBees Docker Pipeline Plugin、	
CloudBees Docker Build and Publish plugin、Go Plugin、Jenkins GitLab Plugin、Gitlab Authentication plugin、Jenkins Git Plugin、Gitlab Hook Plugin（用于gitlab的push触发自动构建）
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

