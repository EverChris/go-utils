package nova

import (
    . "fmt"
    "strconv"

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

//二进制转十六进制
func btox(b string) string {
    base, _ := strconv.ParseInt(b, 2, 10)
    return strconv.FormatInt(base, 16)
}

//十六进制转二进制
func xtob(x string) string {
    base, _ := strconv.ParseInt(x, 16, 10)
    return strconv.FormatInt(base, 2)
}






