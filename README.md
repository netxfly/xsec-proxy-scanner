## xsec-proxy-scanner
### 概述
xsec-proxy-scanner是一款速度超快、小巧的代理服务器扫描器，使用场景为：

1. 定期扫描自已公司服务器，排查是否有外网服务器开启了代理服务；
1. 扫描公网中的代理服务器（搞成分布式的，也可以卖代理了）

支持的协议有：
1. http
1. https
1. socks4
1. socks4a
1. socks5

### 使用说明

```bash
$ ./proxy_scanner
NAME:
   xsec proxy scanner - A SOCKS4/SOCKS4a/SOCKS5/HTTP/HTTPS proxy scanner

USAGE:
   proxy_scanner [global options] command [command options] [arguments...]

VERSION:
   20171205

AUTHOR(S):
   netxfly <x@xsec.io>

COMMANDS:
     scan     start to scan proxy
     dump     dump proxies to a text file
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                 debug mode
   --scan_num value, -n value  scan num (default: 1000)
   --timeout value, -t value   timeout (default: 5)
   --filename value, -f value  filename (default: "iplist.txt")
   --dumpfile value, -o value  filename (default: "xsec_proxies.txt")
   --help, -h                  show help
   --version, -v               print the version
```
### 参数说明

1. scan参数表示开始扫描，扫描结束后会将结果保存到当前目录的`xsec_proxies.db`和`xsec_proxies.txt`文件中，db是数据库文件，txt是文本结果；
1. dump表示将指定DB文件中的结果导出为文本文件。
1. 以下为可选参数：
    - --debug，表示是否启用debug模式，看到具体的扫描过程； 
    - --scan_num，表示每次扫描的服务器数量，默认为1000；
    - --timeout，表示每个扫描请求的超时时间，默认为5秒；
    - --filename，表示要扫描的iplist的文件名，默认为当前目录下的iplist.txt文件
    - dumpfile，表示在使用dump参数时，导出的文本文件名

### 运行截图

![](https://docs.xsec.io/images/xsec_proxy_scanner.png)

从上图中看出，timeout参数调得过短会存在漏报现象，需要根据网络情况合理调整timeout参数。

编译好的二进制版本下载地址[https://github.com/netxfly/xsec-proxy-scanner/releases](https://github.com/netxfly/xsec-proxy-scanner/releases)
