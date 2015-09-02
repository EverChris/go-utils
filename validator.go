package nova

import(
    //"regexp"
    "strings"
)

func IsEmail(s string) bool{
    return true
}

func IsTelno(s string) bool{
    return true
}

func IsIPStr(s string) bool{
    return true
}

func IsIPInt(i int) bool{
    return true
}

//scheme:http,rtmp,rtmfp
func IsURL(s string, scheme string) bool{
    return true
}

//Unix directory starts by `/`
func IsUnixDir(s string) bool{
   return strings.HasPrefix(s,`/`)
}

//只能由字母、数字和下划线组合，且首字符不能是数字
func IsIllegalName(s string) bool{
    return true
}

func HasPrefix(s, prefix string) bool{
    return strings.HasPrefix(s,prefix)
}
