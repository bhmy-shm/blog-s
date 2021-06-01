package util

import "strconv"

/* -------------- 响应处理的类型转换 ------------ */
type StrTo string

// 按照string类型输出
func(s StrTo) String() string{
	return string(s)
}

// string转换成Int类型
func(s StrTo) MustInt()int{
	v,_:=s.int()
	return v
}
func(s StrTo) int() (int,error){
	v,err := strconv.Atoi(s.String())
	return v,err
}

// string 转换成 Uint32类型
func(s StrTo) MustUInt32() uint32{
	v,_:= s.uint32()
	return v
}
func(s StrTo) uint32() (uint32,error){
	v,err := strconv.Atoi(s.String())
	return uint32(v),err
}

func(s StrTo) MustUInt8() uint8{
	v,_:= s.uint8()
	return v
}
func(s StrTo) uint8() (uint8,error){
	v,err := strconv.Atoi(s.String())
	return uint8(v),err
}


