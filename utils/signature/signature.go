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
func (s signature) Sign(data interface{}) (string, error) {
	str, err := json.Marshal(data)
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

//Verify 验证签名
func (s *signature) Verify(data interface{}, sign string) (bool, error) {
	makeSign, err := s.Sign(data)
	if err != nil {
		return false, err
	}

	return makeSign == sign, nil
}
