#Shadowsocks Pinger 

this is simple cli go app whit two phase server and client (for test)

## Use case?
for now just used for check if url connect and worked correctly or not (when looking for valid vpn) <br>
in the future it can use to auto search for available links and somehow send them  

## Requirements
* already installed go

### how to use?
1. clone this project
2. download [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2) ``` go get -u -v github.com/shadowsocks/go-shadowsocks2``` 
3. build unix app with below command
```bash
$ make all
```
4. replace your own shadowsocks link in `.env` file (for now just support outline links)
5. running server side to start shadowsocks server (start socks5 on local port 1080 )
```bash
$ ./ss_pinger serve
```
6. running client side to test connection server every 5 seconds
```bash
$ ./ss_pinger client
```
* every log saved in `service.log` file after running client 