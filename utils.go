package nova

import (
    . "fmt"
    "strconv"
    "time"
)

func main() {
    x:=3.1415926
    Println("你好"+ToString(x))
    Println("nihao"+ToString(0.11))

    ip := IPAddr{1,2,3,4}
    Println(ip)

}
func ToString(a interface{}) string{
    //fmt.Println("type:", reflect.TypeOf(a))
    if  v,p:=a.(int);p{
        return strconv.Itoa(v)
    }

    if v,p:=a.(float64);  p{
        return strconv.FormatFloat(v,'f', -1, 64)
    }

    if v,p:=a.(float32); p {
        return strconv.FormatFloat(float64(v),'f', -1, 32)
    }

    if v,p:=a.(int16); p {
        return strconv.Itoa(int(v))
    }
    if v,p:=a.(uint); p {
        return strconv.Itoa(int(v))
    }
    if v,p:=a.(int32); p {
        return strconv.Itoa(int(v))
    }
    return "wrong"
}


type IPAddr [4]byte
func (ip IPAddr) String() string {
    return Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}


func Time() string{
    return time.Now().Format("2006-1-2 15:04:05")
}




