package flash

import (
    ."fmt"
    "net"
    "os"
    "time"
    //"io/ioutil"
    //."reflect"
    //"strings"
    "net/http"
    "io"
    ."strconv"

)


const (

    TCP_XML  = `<?xml version="1.0" encoding="UTF-8"?>
<cross-domain-policy>
<allow-access-from domain="*" to-ports="*"/>
</cross-domain-policy>`
    HTTP_XML  = `<?xml version="1.0" encoding="UTF-8"?>
<cross-domain-policy>
<allow-access-from domain="*"/>
</cross-domain-policy>`

)




func RunTCP() (int, error){
    listener, err := net.Listen("tcp", "0.0.0.0:843")
    Println("建立843:",listener, err)
    checkErr(err)
    Println("843建立成功!")
    for {
        conn, err := listener.Accept()
        //checkErr(err)
        if err != nil {
            continue
        }
        go doServerStuff(conn)
    }
    return 843, err
}


func RunHTTP(port int) (int, error){
    http.HandleFunc(`/crossdomain.xml`,crossdomain)

    Println(`ListenAndServe:`,":"+Itoa(port))
    err := http.ListenAndServe(":"+Itoa(port), nil)
    if err != nil {
        Println(`ListenAndServe err:`,err)
    }
    return port,err
}

func crossdomain(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/xml")
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, HTTP_XML)
}


func doServerStuff(conn net.Conn) {
    Println("843 doServerStuff>")


    Println("!!!回复xml")
    conn.Write([]byte(TCP_XML))


//        buf := make([]byte, 512) //生成一个缓存数组
//        n, err := conn.Read(buf)
//        Println("843客户端消息：",n,err,string(buf[0:n]))
//
////            if strings.HasPrefix(string(buf),"<policy-file-request/>") {
////                //defer conn.Close()
////                Println("!!!回复xml")
////                conn.Write([]byte(XML))
////            }
//
////    Println("!!!回复xml")
////    conn.Write([]byte(XML))
//
//        for {
//            buf := make([]byte, 512)
//            n, err := conn.Read(buf) //读取客户机发的消息
//            if err != nil {continue}
//            Println(n,err,string(buf[0:n])) //打印出来
//
//        }
//
//    //    for {
////    buf, err := ioutil.ReadAll(conn)
////    Println("843客户端消息:",err,buf,string(buf),len(strings.TrimSpace(string(buf))),len("<policy-file-request/>"))
////    if strings.HasPrefix(string(buf),"<policy-file-request/>") {
////        //defer conn.Close()
////        Println("!!!回复xml")
////        conn.Write([]byte(XML))
////    }
//
    quit := func() {
        time.Sleep(time.Second*0) //给客户端1秒的响应的时间，否则客户端有可能读不到数据就提前Close了
        conn.Close()
        Println("!!!close")
    }
    defer quit()

}

//检查错误
func checkErr(err error) int {
    var code int = 1;

    if err != nil {
        defer os.Exit(code)
        Fprintf(os.Stderr,"Fatal error: %s",err.Error())
        if err.Error() == "EOF" {
            //fmt.Println("用户退出了")
            code = 0
            return code
        }
        code = -1
        return code
    }

    return code
}
