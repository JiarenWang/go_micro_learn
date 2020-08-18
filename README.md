# go_micro_learn

TODO:
1.数据库
2.数据库模型xorm


curl -XPOST -d'{"S":"hello, world"}' localhost:8080/hello


nginx(https)       go  怎么弄?? go的服务器架构  https最好能在前端处理掉 短时存储用redis//写成中间件的形式
    |
    网关  监听服务  rpc 使用grpc  服务注册与发现(discovery)
    |
    具体服务(写成不同服务,不同启动文件)

    grpc 四种



5.go消息队列??? 不同服务通过消息队列一步请求
6.缓存 redis 最好也写中间件 写成单独服务
7.日志是什么 日志中间件 日志收集
8. test
9.认证 https://studyiris.com/example/auth/oauth2.html   令牌token 一般用此https://studyiris.com/example/exper/jwt.html 方式验证登陆 认证token存在redis

10.权限系统?
11.限制请求次数 https://studyiris.com/example/exper/limitHandler.html

//中间件
jwt json web token  Authorization头进行认证
logger
casbin 权限出处理中间件
ctx.readjson(&user)  从请求里读取数据到结构体里
//把结构体类型  转成json
ctx.JSON(doe)
ctx.Params().Get("id")
       
        ctx.JSON(map[string]interface{}{
            //当然，您可以传递任何自定义返回值。
            "user_id": userID,
        })
ctx.Application().Logger().Infof("response sent to " + ctx.Path())
设置一个或多个输出：app.Logger()。SetOutput(io.Writer ...)添加一个或多个输出：app.Logger()。AddOutput(io.Writer ...)



先在go.mod文件里添加依赖,export GO111MODULE=on 然后执行go mod vendor下载到vendor文件夹里
请设置一下第三库下载代理：
export GOPROXY=https://athens.azurefd.net/
如果是Windows操作系统,请设置一下环境变量GOPROXY=https://athens.azurefd.net/
export GOPROXY=https://goproxy.io

最小化 docker image制作
FROM scratch
ADD main /
EXPOSE 8086
CMD ["/main"]

alter user 'root'@'localhost' IDENTIFIED BY 'Applehxx9!';

docker pull editernamef/app:v1

ssh root@49.235.142.16 //oscar那台

https://www.runoob.com/mysql/mysql-install.html

CMD uname -a //显示虚拟化容器的系统信息

使用当前目录的 Dockerfile 创建镜像，标签为 runoob/ubuntu:v1。

docker build -t runoob/ubuntu:v1 .

https://github.com/upx/upx可以对可执行文件压缩

go tool dist list 可以查看go支持的平台和架构

docker container prune 移除所有stopped容器
docker container ls -a
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install
CGO_ENABLED=0 GOOS=linux go install
可以生成可执行文件
CGO_ENABLED=0表示静态链接   https://johng.cn/cgo-enabled-affect-go-static-compile/



docker build -t editernam/app:16.0 .
docker run -p 8086:80 editernam/app:16.0
docker run --mount type=bind,source=/Users/valentine.wang/log,target=/log -p 8086:8086 editernamef/app:v1.0  挂载卷

docker run --mount type=bind,source=/home/log,target=/log -p 8086:8086 editernamef/app:v1.0
docker node ls

# docker login -u xxx -pass yyy https://myrepo.com
# service create --with-registry-auth foo myrepo.com/foo-img:0.1
docker swarm join --token SWMTKN-1-4saspn5f9pjls1tymqj7vdxpf0gsfk99ceqimbkze6uj1aam0b-1y0ssrl0jsb0ujkq60evjcm4l 172.17.0.7:2377

docker service create --with-registry-auth --replicas 2 -p 8086:8086  --name app editernamef/app:v1.0

docker service create --replicas 3 -p 80:80 --name nginx nginx:1.13.7-alpine
docker service ls
docker service ps nginx fuwumoingcx
docker service logs
docker service scale fuwumoingcx=5
docker service rm fuwumoingcx


docker tag ubuntu:18.04 username/ubuntu:18.04
docker push username/ubuntu:18.04
docker search username

docker image ls //只显示顶层镜像
docker image ls -a //显示所有用到的镜像
docker image rm [选项] <镜像1> [<镜像2> ...]



如果使用了 -d 参数运行容器。

$ docker run -d ubuntu:18.04 /bin/sh -c "while true; do echo hello world; sleep 1; done"
77b2dc01fe0f3f1265df143181e7b9af5e05279a884f4776ee75350ea9d8017a

docker container logs [container ID or NAMES]


VOLUME /data


你可以通过以下命令来便捷的查看镜像、容器、数据卷所占用的空间。
$ docker system df


在使用 -d 参数时，容器启动后会进入后台。

某些时候需要进入容器进行操作，包括使用 docker attach 命令或 docker exec 命令，推荐大家使用 docker exec 命令，原因会在下面说明。

docker exec -i 69d1 bash




在用 docker run 命令的时候，使用 --mount 标记来将 数据卷 挂载到容器里。在一次 docker run 中可以挂载多个 数据卷。

下面创建一个名为 web 的容器，并加载一个 数据卷 到容器的 /webapp 目录。

$ docker run -d -P \
    --name web \
    # -v my-vol:/wepapp \
    --mount source=my-vol,target=/webapp \
    training/webapp \
    python app.py


无主的数据卷可能会占据很多空间，要清理请使用以下命令

$ docker volume prune


创建具有5个副本任务的服务（--replicas）
使用该--replicas标志设置复制服务的副本任务数。以下命令redis使用5副本任务创建服务：

$ docker service create --name redis --replicas=5 redis:3.0.6


        let key = "hU8g6Tgv7hTf5Yf4"
        let data :String = "dBhhl0TVyoB99hyWZy0RGtighdjMTDC3mmGr8XAfX/DPf8TfWYA1tasiyeEX1V7qxXT91pQcY0qElRJvryzdYoXJnHHxKt1ZlyEz3TCeIdhxkbgP/kKIM5y/2seEu6kc15503AikH4hCPwvdz39/F34qNcXWHSbnWuFav7CSILcML6ly5vrtWxNUzGHPHePbQWX3otgFQvECLBBbPfw3HMlzohMdP/DyMSSfg4B4eiu2VFp11ZjvZgGtv0nItFu4Iw70SaxAzSw76+UP0ttDmhZXN+bC/rW7rewGJmu3nfdEvo+Lw1oxInci0N1CAvl/3SjumPFaSoJjWmFZRN27cG6wfP6biHSr5j8J1KVGsTo="


//        var result :Array<UInt8>"

        do {
            let aess :AES = try AES(key:key, iv:key, padding: Padding.pkcs7)
            let dsa :String? =  try data.decryptBase64ToString(cipher: aess)
            print(dsa ?? "no data")
        } catch {
            print(error)
        }






        git -c color.branch=false -c color.diff=false -c color.status=false -c diff.mnemonicprefix=false -c core.quotepath=false -c credential.helper=sourcetree add -f -- vendor/github.com/BurntSushi/toml/.gitignore vendor/github.com/CloudyKit/jet/.gitignore vendor/github.com/Shopify/goreferrer/.gitignore vendor/github.com/fatih/structs/.gitignore vendor/github.com/flosch/pongo2/.gitignore vendor/github.com/go-xorm/xorm/.drone.yml vendor/github.com/go-xorm/xorm/.gitignore vendor/github.com/go-xorm/xorm/CONTRIBUTING.md vendor/github.com/go-xorm/xorm/LICENSE vendor/github.com/go-xorm/xorm/README.md vendor/github.com/go-xorm/xorm/README_CN.md vendor/github.com/go-xorm/xorm/cache_lru.go vendor/github.com/go-xorm/xorm/cache_memory_store.go vendor/github.com/go-xorm/xorm/context_cache.go vendor/github.com/go-xorm/xorm/convert.go vendor/github.com/go-xorm/xorm/dialect_mssql.go vendor/github.com/go-xorm/xorm/dialect_mysql.go vendor/github.com/go-xorm/xorm/dialect_oracle.go vendor/github.com/go-xorm/xorm/dialect_postgres.go vendor/github.com/go-xorm/xorm/dialect_sqlite3.go vendor/github.com/go-xorm/xorm/doc.go vendor/github.com/go-xorm/xorm/engine.go vendor/github.com/go-xorm/xorm/engine_cond.go vendor/github.com/go-xorm/xorm/engine_context.go vendor/github.com/go-xorm/xorm/engine_group.go vendor/github.com/go-xorm/xorm/engine_group_policy.go vendor/github.com/go-xorm/xorm/engine_table.go vendor/github.com/go-xorm/xorm/error.go vendor/github.com/go-xorm/xorm/gen_reserved.sh vendor/github.com/go-xorm/xorm/go.mod vendor/github.com/go-xorm/xorm/go.sum vendor/github.com/go-xorm/xorm/helpers.go vendor/github.com/go-xorm/xorm/helpler_time.go vendor/github.com/go-xorm/xorm/interface.go vendor/github.com/go-xorm/xorm/json.go vendor/github.com/go-xorm/xorm/logger.go vendor/github.com/go-xorm/xorm/pg_reserved.txt vendor/github.com/go-xorm/xorm/processors.go vendor/github.com/go-xorm/xorm/rows.go vendor/github.com/go-xorm/xorm/session.go vendor/github.com/go-xorm/xorm/session_cols.go vendor/github.com/go-xorm/xorm/session_cond.go vendor/github.com/go-xorm/xorm/session_context.go vendor/github.com/go-xorm/xorm/session_convert.go vendor/github.com/go-xorm/xorm/session_delete.go vendor/github.com/go-xorm/xorm/session_exist.go vendor/github.com/go-xorm/xorm/session_find.go vendor/github.com/go-xorm/xorm/session_get.go vendor/github.com/go-xorm/xorm/session_insert.go vendor/github.com/go-xorm/xorm/session_iterate.go vendor/github.com/go-xorm/xorm/session_query.go vendor/github.com/go-xorm/xorm/session_raw.go vendor/github.com/go-xorm/xorm/session_schema.go vendor/github.com/go-xorm/xorm/session_stats.go vendor/github.com/go-xorm/xorm/session_tx.go vendor/github.com/go-xorm/xorm/session_update.go vendor/github.com/go-xorm/xorm/statement.go vendor/github.com/go-xorm/xorm/statement_args.go vendor/github.com/go-xorm/xorm/statement_columnmap.go vendor/github.com/go-xorm/xorm/statement_exprparam.go vendor/github.com/go-xorm/xorm/statement_quote.go vendor/github.com/go-xorm/xorm/syslogger.go vendor/github.com/go-xorm/xorm/tag.go vendor/github.com/go-xorm/xorm/test_mssql.sh vendor/github.com/go-xorm/xorm/test_mssql_cache.sh vendor/github.com/go-xorm/xorm/test_mymysql.sh vendor/github.com/go-xorm/xorm/test_mymysql_cache.sh vendor/github.com/go-xorm/xorm/test_mysql.sh vendor/github.com/go-xorm/xorm/test_mysql_cache.sh vendor/github.com/go-xorm/xorm/test_postgres.sh vendor/github.com/go-xorm/xorm/test_postgres_cache.sh vendor/github.com/go-xorm/xorm/test_sqlite.sh vendor/github.com/go-xorm/xorm/test_sqlite_cache.sh vendor/github.com/go-xorm/xorm/test_tidb.sh vendor/github.com/go-xorm/xorm/transaction.go vendor/github.com/go-xorm/xorm/types.go vendor/github.com/go-xorm/xorm/xorm.go vendor/github.com/iris-contrib/blackfriday/.gitignore vendor/github.com/json-iterator/go/.gitignore vendor/github.com/juju/errors/.gitignore vendor/github.com/kataras/golog/.gitignore vendor/github.com/kataras/iris/.gitignore vendor/github.com/kataras/pio/.gitignore vendor/github.com/klauspost/cpuid/.gitignore vendor/github.com/modern-go/concurrent/.gitignore vendor/github.com/modern-go/reflect2/.gitignore vendor/golang.org/x/sys/unix/.gitignore vendor/xorm.io/builder/.drone.yml vendor/xorm.io/builder/LICENSE vendor/xorm.io/builder/README.md vendor/xorm.io/builder/builder.go vendor/xorm.io/builder/builder_delete.go vendor/xorm.io/builder/builder_insert.go vendor/xorm.io/builder/builder_join.go vendor/xorm.io/builder/builder_limit.go vendor/xorm.io/builder/builder_select.go vendor/xorm.io/builder/builder_union.go vendor/xorm.io/builder/builder_update.go vendor/xorm.io/builder/cond.go vendor/xorm.io/builder/cond_and.go vendor/xorm.io/builder/cond_between.go vendor/xorm.io/builder/cond_compare.go vendor/xorm.io/builder/cond_eq.go vendor/xorm.io/builder/cond_expr.go vendor/xorm.io/builder/cond_if.go vendor/xorm.io/builder/cond_in.go vendor/xorm.io/builder/cond_like.go vendor/xorm.io/builder/cond_neq.go vendor/xorm.io/builder/cond_not.go vendor/xorm.io/builder/cond_notin.go vendor/xorm.io/builder/cond_null.go vendor/xorm.io/builder/cond_or.go vendor/xorm.io/builder/doc.go vendor/xorm.io/builder/error.go vendor/xorm.io/builder/go.mod vendor/xorm.io/builder/go.sum vendor/xorm.io/builder/sql.go vendor/xorm.io/builder/writer.go vendor/xorm.io/core/.drone.yml vendor/xorm.io/core/.gitignore vendor/xorm.io/core/LICENSE vendor/xorm.io/core/README.md vendor/xorm.io/core/benchmark.sh vendor/xorm.io/core/cache.go vendor/xorm.io/core/column.go vendor/xorm.io/core/converstion.go vendor/xorm.io/core/db.go vendor/xorm.io/core/dialect.go vendor/xorm.io/core/driver.go vendor/xorm.io/core/error.go vendor/xorm.io/core/filter.go vendor/xorm.io/core/go.mod vendor/xorm.io/core/go.sum vendor/xorm.io/core/ilogger.go vendor/xorm.io/core/index.go vendor/xorm.io/core/mapper.go vendor/xorm.io/core/pk.go vendor/xorm.io/core/rows.go vendor/xorm.io/core/scan.go vendor/xorm.io/core/stmt.go vendor/xorm.io/core/table.go vendor/xorm.io/core/tx.go vendor/xorm.io/core/type.go


        git -c color.branch=false -c color.diff=false -c color.status=false -c diff.mnemonicprefix=false -c core.quotepath=false -c credential.helper=sourcetree commit -q -F /var/folders/v1/3ycwjjj507v1kz074j5_80wh0000gn/T/SourceTreeTemp.dKxz6X -a
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/.travis.yml.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/AUTHORS.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/CODE_OF_CONDUCT.md.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/CONTRIBUTING.md.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/HISTORY.md.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/LICENSE.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/cache/AUTHORS.
        The file will have its original line endings in your working directory.
        warning: CRLF will be replaced by LF in vendor/github.com/kataras/iris/cache/LICENSE.
        The file will have its original line endings in your working directory.


        git -c color.branch=false -c color.diff=false -c color.status=false -c diff.mnemonicprefix=false -c core.quotepath=false -c credential.helper=sourcetree push -v --tags --set-upstream origin refs/heads/develop:refs/heads/develop
        Pushing to https://gitlab.com/0x001A/gobone
        warning: redirecting to https://gitlab.com/0x001A/gobone.git/
        POST git-receive-pack (196273 bytes)
        remote:
        remote: To create a merge request for develop, visit:
        remote:   https://gitlab.com/0x001A/gobone/merge_requests/new?merge_request%5Bsource_branch%5D=develop
        remote:
        Branch develop set up to track remote branch develop from origin.
        To https://gitlab.com/0x001A/gobone
           30a9504..cbf81fc  develop -> develop
        updating local tracking ref 'refs/remotes/origin/develop'
        Completed successfully




        VSCODE安装插件
        创建目录$GOPATH/src/golang.org/x，并切换到该目录
        mkdir -p $GOPATH/src/golang.org/x/
        cd $GOPATH/src/golang.org/x/
        克隆golang.org工具源码
        如果不克隆的话，go get -u -v golang.org/xxx肯定是timeout的，所以只能先把它们下载到本地src/golang.org/x/tools
        git clone https://github.com/golang/tools.git
        git clone https://github.com/golang/lint.git
        go env -w GOPROXY=https://goproxy.io,direct
        export GO111MODULE=on

        go env -w GOPROXY=https://goproxy.io,direct
go get golang.org/x/tools/gocode


MacBook-Pro:gobone wangjairen$ go env -w GOPROXY=https://goproxy.io,direct
AMacBook-Pro:gobone wangjairen$ export GO111MODULE=on

