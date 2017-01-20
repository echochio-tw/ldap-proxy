# ldap-proxy

1. Edit conf.json
2. go run local_ldap.go
3. run sample file ..... or other you access ldap file

[root@test ldap-local]# go run local_ldap.go
[server] 2017/01/20 11:33:48 Listening on 192.168.0.70:10389
[server] 2017/01/20 11:33:51 Connection client [1] from 192.168.0.48:56056 accepted
[server] 2017/01/20 11:33:51 <<< 1 - BindRequest - hex=&{301802010160130201030404656e6c6980085040737377307264}
2017/01/20 11:33:51 AD server 192.168.0.220
2017/01/20 11:33:51 user enli
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

sample file :
------------------------------------------------------------------------------------
import ldap

def authenticate(address, username, password):
    conn = ldap.initialize('ldap://' + address)
    conn = ldap.open(host=address, port=10389)
    conn.protocol_version = 3
    conn.set_option(ldap.OPT_REFERRALS, 0)
    try:
        res=conn.simple_bind_s(username, password)
        print type(res)
        print res
    except ldap.INVALID_CREDENTIALS:
        return "Invalid credentials"
    except ldap.SERVER_DOWN:
        return "Server down"
    except ldap.LDAPError, e:
        if type(e.message) == dict and e.message.has_key('desc'):
            return "Other LDAP error: " + e.message['desc']
        else: 
            return "Other LDAP error: " + e
    finally:
        conn.unbind_s()
    return "Succesfully authenticated"

#change your go server ip & ad server user/pass 
print authenticate("192.168.0.70","enli","P@ssw0rd")
