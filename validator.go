package nova

import(
    "regexp"
    "strings"
    "strconv"
)

func IsEmail(s string) bool{
    reg := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
    return reg.MatchString(s)
}

func IsTelno(i int) bool{
    reg := regexp.MustCompile(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`)
    return reg.MatchString(strconv.Itoa(i))
}

func IsIPStr(s string) bool{
    reg := regexp.MustCompile(`^(?:(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])\.){3}(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])$`)
    return reg.MatchString(s)
}

//四个字节，32位整数，int32
func IsIPInt(i int) bool{
    if i > 4294967295 || i < 0 {
        return false
    }
    return true
}

func IsURL(s string) bool{
    reg := regexp.MustCompile(`[a-zA-z]+://[^\s]*`)
    return reg.MatchString(s)
}

//scheme:http,rtmp,rtmfp
func IsSpecifiedURL(s, scheme string) bool{
    reg := regexp.MustCompile(scheme+`://[^\s]*`)
    return reg.MatchString(s)
}

//Unix directory starts by `/`
func IsUnixDir(s string) bool{
   return strings.HasPrefix(s,`/`)
}

//只能由字母、数字和下划线组合，且首字符不能是数字
func IsIllegalName(s string) bool{
    b := []byte(s)
    if b[0] >= 48 && b[0] <= 57 {
       return false
    }
    reg := regexp.MustCompile(`^[A-Za-z0-9_]+$`)
    return reg.MatchString(s)
}

func HasPrefix(s, prefix string) bool{
    return strings.HasPrefix(s,prefix)
}
