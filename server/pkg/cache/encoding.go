package cache

import (
	"encoding"
	"encoding/json"
)

// Encoding 表示编码接口
type Encoding interface {
	Marshal(v interface{}) ([]byte, error)
	UnMarshal(data []byte, v interface{}) error
}

/*对外提供的Marshal和Unmarshal*/

// Marshal 返回marshal后的数据
func Marshal(e Encoding, v interface{}) (data []byte, err error) {

	bm, ok := v.(encoding.BinaryMarshaler)
	if ok && e == nil {
		data, err = bm.MarshalBinary()
		return
	}

	data, err = e.Marshal(v)
	if err == nil {
		return
	}
	if ok {
		data, err = bm.MarshalBinary()
	}
	return
}

func Unmarshal(e Encoding, data []byte, v interface{}) (err error) {
	bm, ok := v.(encoding.BinaryUnmarshaler)
	if ok && e == nil {
		err = bm.UnmarshalBinary(data)
		return err
	}
	err = e.UnMarshal(data, v)
	if err == nil {
		return
	}
	if ok {
		return bm.UnmarshalBinary(data)
	}
	return
}

/*JSON encoding ：JSON编码的encoding接口实现*/

// JSONEncoding 代表JSON编码
type JSONEncoding struct{}

func (j JSONEncoding) Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	return buf, err
}

func (j JSONEncoding) UnMarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
