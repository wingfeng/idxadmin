package entity

import (
	"strconv"
)

// 解决前端大数精度丢失问题
type ID int64

// MarshalJSON returns a json byte array string of the snowflake ID.
func (f ID) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')
	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (f *ID) UnmarshalJSON(b []byte) error {

	if b[0] == '"' && b[len(b)-1] != '"' {
		i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
		if err != nil {
			return err
		}

		*f = ID(i)
		return nil
	} else {
		//byte to int64
		i, err := strconv.ParseInt(string(b), 10, 64)
		if err != nil {
			return err
		}
		*f = ID(i)
		return nil
	}

}

func (f *ID) String() string {
	str := strconv.FormatInt(int64(*f), 10)
	return str
}
