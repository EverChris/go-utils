/**
 * Package nova
 * Handle golang's bytes in AS3/JS style.
 */
package nova

import (
    "bytes"
    //"fmt"
)

type byteArray struct{
    Buffer *bytes.Buffer
}

type _ByteArray struct{}

var ByteArray _ByteArray = _ByteArray{}

func (ba _ByteArray) New() byteArray{
    buf := bytes.NewBuffer([]byte{})
    return byteArray{buf}
}

/***/
func (ba byteArray) Write(b []byte) byteArray{
    ba.Buffer.Write(b);
    return ba
}


//写入一个字节的8位整数
func (ba byteArray) WriteByte(i int8) byteArray{
    ba.Buffer.WriteByte(byte(i))
    return ba
}

//写入两个字节的16位整数,超过两个字节的部分截断,BigEndian
func (ba byteArray) WriteShort(i int16) byteArray{
    high := int8((i & 0xf0) >> 4)
    low := int8(i & 0x0f)
    ba.WriteByte(high)
    ba.WriteByte(low)
    return ba
}

//写入四个字节的32位整数,超过四个字节的部分截断,BigEndian
func (ba byteArray) WriteInt(i int32) byteArray{
    high := int16((i & 0xff00) >> 8)
    low := int16(i & 0x00ff)
    ba.WriteShort(high)
    ba.WriteShort(low)
    return ba
}

//写入四个字节的32位浮点数
func (ba byteArray) WriteFloat32(f float32) byteArray{
    //TODO
    return ba
}

//写入八个字节的64位浮点数
func (ba byteArray) WriteFloat64(f float64) byteArray{
    //TODO
    return ba
}

//写入八个字节的64位整数,超过八个字节的部分截断,BigEndian
func (ba byteArray) WriteDouble(i int64) byteArray{
    high := int32((i & 0xffff0000) >> 16)
    low := int32(i & 0x0000ffff)
    ba.WriteInt(high)
    ba.WriteInt(low)
    return ba
}

//写入utf8字符串
func (ba byteArray) WriteUTF8String(s string) byteArray{
    ba.Buffer.WriteString(s)
    return ba
}

//Writes a sequence of length bytes from the specified byte array, bytes, starting offset(zero-based index) bytes into the byte stream.
//If the length parameter is omitted, the default length of 0 is used;
//the method writes the entire buffer starting at offset. If the offset parameter is also omitted, the entire buffer is written.
//If offset or length is out of range, they are clamped to the beginning and end of the bytes array.
//Parameters:
//bytes - The ByteArray object.
//offset - A zero-based index indicating the position into the array to begin writing.
//length - An unsigned integer indicating how far into the buffer to write.    //TODO

//把指定byteArray的全部数据复制到当前byteArray，源数据不受影响，offset为在当前容器中写数据的起始位置
func (ba byteArray) WriteBytes(_ba byteArray, offset uint) byteArray{
    b := _ba.Buffer.Bytes()
    for offset > 0{
        offset --
        ba.Buffer.WriteByte(0)
    }
    ba.Buffer.Write(b)
    return ba
}

/***/

//读出一个字节的8位整数
func (ba byteArray) ReadByte() int8{
    b,_ := ba.Buffer.ReadByte()
    return int8(b)
}

//读出两个字节的16位整数
func (ba byteArray) ReadShort() int16{
    high := ba.ReadByte()
    low := ba.ReadByte()

    return int16(high) << 4 + int16(low)
}

//读出四个字节的32位整数
func (ba byteArray) ReadInt() int32{
    high := ba.ReadShort()
    low := ba.ReadShort()
    return int32(high) << 8 + int32(low)
}

//读出八个字节的64位整数
func (ba byteArray) ReadDouble() int64{
    high := ba.ReadInt()
    low := ba.ReadInt()
    return int64(high) << 16 + int64(low)
}

//读出指定长度的utf8字符串
func (ba byteArray) ReadUTF8String(length int) string{
    s := ""
    for len(s) < length {
        b,_ := ba.Buffer.ReadByte();
        s += string(b)
    }
    return s
}

//Reads the number of data bytes, specified by the length parameter, from the byte stream.
//The bytes are read into the ByteArray object specified by the bytes parameter,
//and the bytes are written into the destination ByteArray starting at the position specified by offset.
//Parameters:
//bytes - The ByteArray object to read data into.
//offset - The offset (position) in bytes at which the read data should be written.
//length - The number of bytes to read. The default value of 0 causes all available data to be read.  //TODO

//把当前byteArray的全部数据复制到指定byteArray，源数据不受影响，offset为在目标容器中写数据的起始位置
func (ba byteArray) ReadBytes(_ba byteArray, offset uint) byteArray{
    b := ba.Buffer.Bytes()
    for offset > 0{
        offset --
        _ba.Buffer.WriteByte(0)
    }
    _ba.Buffer.Write(b)
    return _ba
}

/***/
func (ba byteArray) Copy() (byteArray){
    _ba := ByteArray.New()
    return ba.ReadBytes(_ba,0)
}

func (ba byteArray) Len() int {
    return ba.Buffer.Len()
}

//Non-utf8string will not be recognized
func (ba byteArray) String() string {
    return ba.Buffer.String()
}

func (ba byteArray) Bytes() []byte {
    return ba.Buffer.Bytes()
}

func (ba byteArray) Clear() {
    ba.Buffer.Reset();
}


