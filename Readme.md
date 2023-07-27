# Shadowsocks Pinger

Shadowsocks Pinger is a simple command-line Go application with two phases: server and client (for testing purposes).

## Use Case

The main use case of this application is to check if a URL can connect and work correctly. It is particularly useful for verifying the validity of VPN connections. In the future, the application may be expanded to automatically search for available links and somehow send them.

## Requirements

- Go (already installed)
- [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2) (Install with `go get -u -v github.com/shadowsocks/go-shadowsocks2`)

## How to Use

1. Clone this project to your local machine.
2. Download and install `go-shadowsocks2` by running the following command in your terminal:
   ```
   go get -u -v github.com/shadowsocks/go-shadowsocks2
   ```
3. Build the Unix app by executing the following command:
   ```
   make all
   ```
4. Replace your own Shadowsocks link in the `.env` file. Currently, only Outline links are supported.
5. To start the Shadowsocks server and enable SOCKS5 on local port 1080, run the following command:
   ```
   ./ss_pinger serve
   ```
6. To test the connection to the server every 5 seconds, run the following command:
   ```
   ./ss_pinger client
   ```
7. After running the client, every log will be saved in the `service.log` file.

## Contribution

Contributions to the Shadowsocks Pinger project are welcome. If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## Disclaimer

This project is for educational and testing purposes only. Please use it responsibly and in accordance with the laws and regulations of your country. The creators of this project are not responsible for any misuse or damage caused by this application.
