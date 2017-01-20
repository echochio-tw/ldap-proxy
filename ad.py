import ldap
__author__ = 'nixCraft'
import sys
def authenticate(address, username, password):
    conn = ldap.initialize('ldap://' + address)
    conn = ldap.open(host=address, port=389)
    conn.protocol_version = 3
    conn.set_option(ldap.OPT_REFERRALS, 0)
    try:
        res=conn.simple_bind_s(username, password)
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

address = str(sys.argv[1])
username = str(sys.argv[2])
password = str(sys.argv[3])
print authenticate(address,username,password)
