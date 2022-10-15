#!/bin/bash


go-shadowsocks2 -cipher aes-256-gcm -c "$1" -verbose -socks 127.0.0.1:1080  >> service.log 2>&1
