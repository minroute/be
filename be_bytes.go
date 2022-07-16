package be

import (
	"encoding/json"
	"math"
	"reflect"
	"regexp"
)

var uint8Arr [8]uint8

//  将任意专程byte数组，
//  如果是map类型，会对其进行json序列化后返回，失败将返回一个空成员{}的的字节数组:[]byte{}
//	如果是二进制类型， 会尝试进行转换，失败将返回一个空成员{}的的字节数组:[]byte{}
func beBytes(v any) []byte {
	if v == nil {
		return nil
	}
	switch value := v.(type) {
	case string:
		return []byte(value)

	case []byte:
		return value

	default:
		if f, ok := value.(iBytes); ok {
			return f.Bytes()
		}

		var rv = reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Map:
			bytes, err := json.Marshal(v)
			if err != nil {
				//@todo ：此处需要进一步做跟进，目前解析失败，直接返回空bytes
				//intlog.Errorf(context.TODO(), `%+v`, err)
				return []byte{}
			}
			return bytes

		case reflect.Array, reflect.Slice:
			var (
				ok    = true
				bytes = make([]byte, rv.Len())
			)
			for i := range bytes {
				int32Value := beInt32(rv.Index(i).Interface())
				if int32Value < 0 || int32Value > math.MaxUint8 {
					ok = false
					break
				}
				bytes[i] = byte(int32Value)
			}
			if ok {
				return bytes
			}
		}

		// 二进制处理
		//@todo：此处的二进制处理，需要进一步学习。 原学习对象有自己的一套方式，需要进一步学习掌握
		rbDel := regexp.MustCompile(`[^01]`)
		s := rbDel.ReplaceAllString(beString(v), "")
		l := len(s)
		if l == 0 {
			//@todo ：此处需要进一步做跟进，目前解析失败，直接返回空bytes
			//intlog.Errorf(context.TODO(), `%+v`, err)
			return []byte{}
		}
		mo := l % 8
		l /= 8
		if mo != 0 {
			l++
		}
		bs := make([]byte, 0, l)
		mo = 8 - mo
		var n uint8
		for i, b := range []byte(s) {
			m := (i + mo) % 8
			switch b {
			case byte('1'):
				n += uint8Arr[m]
			}
			if m == 7 {
				bs = append(bs, n)
				n = 0
			}
		}
		return bs

		//return gbinary.Encode(v)
	}
}
