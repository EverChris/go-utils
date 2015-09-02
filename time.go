package nova

import "time"

const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"

// RFC822 with numeric zone RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"

// RFC1123 with numeric zone RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"

// Handy time stamps. Stamp      = "Jan _2 15:04:05"
    Kitchen     = "3:04PM"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)


//返回格式：1389058332
func Now() int64{
    return time.Now().Unix()
}

//返回格式：2014-01-08 09:04:41
func NowFormat() string{
    return time.Now().Format("2006-01-02 15:04:05")
}

//返回格式：2014-01-08 09:04:41
func TimeFormat(timestamp int)string{
    return time.Unix(timestamp,0).Format("2006-01-02 15:04:05")
}

//参数格式：2014-01-08 09:04:41
func StrToTimestamp(s string) int64{
    t,_ := time.Parse("2006-01-02 15:04:05",s)   //   "2014-01-08 09:04:41"
    return t.Unix()

}


