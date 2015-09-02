package nova

import . "encoding/json"

var JSON json = json{encode,decode}

type json struct{
    Encode func(x interface{})string `JSON.Encode`
    Decode func(s string)interface{} `JSON.Decode`
}

func encode(v interface{}) string {
    result,_ := Marshal(v);
    return string(result)
}

func decode(s string) interface{} {
    var result interface{}
    Unmarshal([]byte(s), &result)
    return result
}

