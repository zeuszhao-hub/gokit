package signature

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
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
	fmt.Println(data)
	str, err := json.Marshal(data)
	fmt.Println(str)
	if err != nil {
		return "", err
	}
	query := url.Values{}
	query.Add("data", string(str))
	query.Add("salt", s.salt)
	query.Add("appid", s.appid)
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

	if s.salt != "" {
		m["salt"] = s.salt
	}

	if s.expire != "" {
		m["expire"] = s.expire
	}

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
