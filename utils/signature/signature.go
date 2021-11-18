package signature

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type signature struct {
	appid  string
	secret string
	salt   string
	expire string
}

//New 创建sign签名对象
func New(appid, secret string) *signature {
	return &signature{
		appid:  appid,
		secret: secret,
	}
}

//Salt 增加盐值
func (s *signature) Salt(salt string) *signature {
	s.salt = salt
	return s
}

//Expire 增加时间
func (s *signature) Expire(expire string) *signature {
	s.expire = expire
	return s
}

//Sign 签名
func (s *signature) Sign(data interface{}) (string, error) {
	query := url.Values{}
	query.Add("appid", s.appid)
	query.Add("data", strval(data))
	query.Add("salt", s.salt)
	query.Add("expire", s.expire)
	queryStr := query.Encode()
	stringSign := queryStr + fmt.Sprintf("&secret=%s", s.secret)
	h := md5.New()
	h.Write([]byte(stringSign))
	signStr := hex.EncodeToString(h.Sum(nil))
	return signStr, nil
}

//SignData 对数据签名并返回签名数据
func (s *signature) SignData(data interface{}) (map[string]interface{}, error) {
	sign, err := s.Sign(data)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m["appid"] = s.appid
	m["data"] = data
	m["sign"] = sign
	m["salt"] = s.salt
	m["expire"] = s.expire

	return m, nil
}

//Verify 验证签名
func (s *signature) Verify(data interface{}, sign string) (bool, error) {
	makeSign, err := s.Sign(data)
	if err != nil {
		return false, err
	}

	return makeSign == sign, nil
}

func strval(v interface{}) string {
	var value string
	if v == nil {
		return value
	}

	switch v.(type) {
	case float64:
		ft := v.(float64)
		value = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := v.(float32)
		value = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := v.(int)
		value = strconv.Itoa(it)
	case uint:
		it := v.(uint)
		value = strconv.Itoa(int(it))
	case int8:
		it := v.(int8)
		value = strconv.Itoa(int(it))
	case uint8:
		it := v.(uint8)
		value = strconv.Itoa(int(it))
	case int16:
		it := v.(int16)
		value = strconv.Itoa(int(it))
	case uint16:
		it := v.(uint16)
		value = strconv.Itoa(int(it))
	case int32:
		it := v.(int32)
		value = strconv.Itoa(int(it))
	case uint32:
		it := v.(uint32)
		value = strconv.Itoa(int(it))
	case int64:
		it := v.(int64)
		value = strconv.FormatInt(it, 10)
	case uint64:
		it := v.(uint64)
		value = strconv.FormatUint(it, 10)
	case string:
		value = v.(string)
	case []byte:
		value = string(v.([]byte))
	default:
		newValue, _ := json.Marshal(v)
		value = string(newValue)
	}
	return value
}
