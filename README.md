# ldap-proxy

Note need python 2.7 & golang

must install python library ldap

copy github.com/vjeantet/ldapserver as ldap server in ldapserver directory

1. Edit conf.json:
2. exec --> go run server.go (or build server.go & exec it)
3. run sample file ..... or other you access ldap file
4. stop server with ctrl-c

[root@test ldap-local]# go run local_ldap.go

[server] 2017/01/20 11:33:48 Listening on 192.168.0.101:10389

[server] 2017/01/20 11:33:51 Connection client [1] from 192.168.0.66:56056 accepted

[server] 2017/01/20 11:33:51 <<< 1 - BindRequest - hex=&{301802010160130201030404656e6c6980085040737377307264}

2017/01/20 11:33:51 AD server 192.168.0.100

2017/01/20 11:33:51 user admin

2017/01/20 11:33:51 pass P@ssw0rd

2017/01/20 11:33:51 Succesfully authenticated

2017/01/20 11:33:51 Succesfully User=enli, Pass=P@ssw0rd

[server] 2017/01/20 11:33:51 >>> 1 - BindResponse - hex=300c02010161070a010004000400

[server] 2017/01/20 11:33:51 <<< 1 - UnbindRequest - hex=&{30050201024200}

[server] 2017/01/20 11:33:51 client 1 close()

[server] 2017/01/20 11:33:51 client 1 close() - stop reading from client

[server] 2017/01/20 11:33:51 client 1 close() - Abandon signal sent to processors

[server] 2017/01/20 11:33:51 client [1] request processors ended

[server] 2017/01/20 11:33:51 client [1] connection closed

^C[server] 2017/01/20 11:34:52 gracefully closing client connections...

[server] 2017/01/20 11:34:52 all clients connection closed

用 golang 當 ldap 主機 ...但 ldap client 寫不出來 ....挖哩 golang 沒有找到可用的

用 golang 呼叫 python --> ad.py 去當 ldap client

想說當不同的 domain 時　go-ldap server 去找不同的 LDAP 或去找 mysql 認證 

先寫個簡單的用 AD 的 LDAP 當參考 .....

所有原始碼 : https://github.com/chio-nzgft/ldap-proxy

先編輯  conf.json  檔案 

改成服務的主機 & AD 主機的 IP

修改  client-ldap-proxy-test.py  內的 

print authenticate("192.168.0.101","admin","P@ssw0rd")

修改 你的 主機及 帳號密碼 資訊

看執行後成果 是否是認證成功  Succesfully authenticated

這樣你就可以修改  server.go  

看 username 就可找不同 AD 或 ldap

例如 :
admin@test1.com 找 ad1
admin@test1.com 找 ad2
