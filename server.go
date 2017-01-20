package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "os/exec"
    "strings"
    "encoding/json"
    ldap "./ldapserver"
)

type Config struct {
       ListenAddr string
       AdserverIP string
}

var c Config

func main() {
    //read config
    r, err := os.Open("conf.json")
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)

    err = decoder.Decode(&c)
    if err != nil {
        log.Fatalln(err)
    }

    //ldap logger
    ldap.Logger = log.New(os.Stdout, "[server] ", log.LstdFlags)

    //Create a new LDAP Server
    server := ldap.NewServer()

    routes := ldap.NewRouteMux()
    routes.Bind(handleBind)
    server.Handle(routes)

    // listen on 10389
    go server.ListenAndServe(c.ListenAddr)

    // When CTRL+C, SIGINT and SIGTERM signal occurs
    // Then stop server gracefully
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    close(ch)

    server.Stop()
}

func handleBind(w ldap.ResponseWriter, m *ldap.Message) {
    r := m.GetBindRequest()
    res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
    app := "python"
    arg0 := "ad.py"
    arg2 := string(r.Name())
    arg3 := string(r.AuthenticationSimple())

    log.Printf("AD server %s", c.AdserverIP)
    log.Printf("user %s", arg2)
    log.Printf("pass %s", arg3)
    cmd := exec.Command(app, arg0, c.AdserverIP, arg2, arg3)
    stdout, err :=cmd.Output()
    outinfo := string(stdout)
    outinfo = strings.Replace(outinfo, "\n", "", -1)
    log.Printf(outinfo)
    if err != nil {
        log.Printf("system error !!")
        res.SetResultCode(ldap.LDAPResultInvalidCredentials)
        res.SetDiagnosticMessage("system error !!")
        w.Write(res)
        return
    }

    if outinfo == "Succesfully authenticated" {
        w.Write(res)
        log.Printf("Succesfully User=%s, Pass=%s", string(r.Name()), string(r.AuthenticationSimple()))
        return
    }

    log.Printf("Bind failed User=%s, Pass=%s", string(r.Name()), string(r.AuthenticationSimple()))
    res.SetResultCode(ldap.LDAPResultInvalidCredentials)
    res.SetDiagnosticMessage("invalid credentials")
    w.Write(res)
}
