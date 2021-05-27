package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)


type deadlinedConn struct {
	Timeout time.Duration
	net.Conn
}

func NewDeadlineTransport(timeout time.Duration) *http.Transport {
	transport := &http.Transport{
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, timeout)
			if err != nil {
				return nil, err
			}
			return &deadlinedConn{timeout, c}, nil
		},
	}
	return transport
}


func HttpPostWithCookie(url string, cookies []*http.Cookie, headers map[string]string) ([]byte, int, error) {
	var (
		resp       *http.Response
		statusCode int
	)
	request, err := http.NewRequest("POST", url, nil)
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}
	httpclient := &http.Client{Transport: NewDeadlineTransport(5 * time.Second)}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	resp, err = httpclient.Do(request)

	if resp != nil {
		statusCode = resp.StatusCode
	}
	if err != nil {
		return nil, statusCode, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, statusCode, err
	}
	return b, statusCode, nil
}


func HttpPost2(url string, headers map[string]string) ([]byte, int, []*http.Cookie, error) {
	var (
		resp       *http.Response
		statusCode int
		cookies    []*http.Cookie
	)
	request, err := http.NewRequest("POST", url, nil)
	httpclient := &http.Client{Transport: NewDeadlineTransport(5 * time.Second)}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	resp, err = httpclient.Do(request)
	if resp != nil {
		statusCode = resp.StatusCode
		cookies = resp.Cookies()
	}
	if err != nil {
		return nil, statusCode, nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, statusCode, cookies, err
	}
	return b, statusCode, cookies, nil
}


func SortMapToURL(mReq map[string]interface{}, join1, join2 string) string {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用join1,join2连接起来，略过空值  如keyjoin1valuejoin2 a=1&b=2
	var signStrings string
	for i, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + join1 + value + join2
			} else {
				signStrings = signStrings + k + join1 + value //最后一个不加此符号
			}
		} else {
			signStrings = signStrings + k + join1 + join2
		}
	}
	return strings.Trim(signStrings, join2)
}


type MMIDLoginRequest struct {
	Account     string `json:"account"`     //账号
	RememberAcc int    `json:"rememberAcc"` //是否记住账号
	Passwd      string `json:"passwd"`      //密码
	RememberPwd int    `json:"rememberPwd"` //是否记住账号
	Vericode    int    `json:"vericode"`    //验证码：当需要防刷时需要传入
	Game        int    `json:"game"`        //游戏id: (新游戏加入需要申请)
	Tad         string `json:"tad"`         //推广广告位
}

type MMIDLoginResponse struct {
	Result  int                `json:"result"`   //0：发送成功；其他：发送失败
	Data    *MMIDLoginDataInfo `json:"data"`     //json数据
	ErrDesc string             `json:"err_desc"` //发送失败时返回，错误描述
}


type MMIDLoginDataInfo struct {
	Uid  uint64 `json:"uid"` //米米号
	Ip   string `json:"ip"`  //ip
	Time string `json:"t"`   //时间戳
	Sid  string `json:"sid"` //session
}


func SDKMMIDLoginRequest(url string, reqData MMIDLoginRequest, cookies []*http.Cookie,
	ip string) (*MMIDLoginResponse, []*http.Cookie, error) {
	var (
		result []byte
		err    error
		cs     []*http.Cookie
		header map[string]string
	)
	header = make(map[string]string)
	header["FROM"] = ip

	info := &MMIDLoginResponse{}
	mReq := make(map[string]interface{})
	v := reflect.ValueOf(reqData)
	t := reflect.TypeOf(reqData)
	for i := 0; i < t.NumField(); i++ {
		value := fmt.Sprintf("%v", v.Field(i).Interface())
		if len(value) == 0 {
			continue
		}
		mReq[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	data := SortMapToURL(mReq, "=", "&")
	data = url + "?" + data

	//logger.Debugf("sdk mmid login request %v", string(data))
	if len(cookies) == 0 {
		result, _, cs, err = HttpPost2(data, header)
	} else {
		result, _, err = HttpPostWithCookie(data, cookies, header)
	}
	if err != nil {
		return info, cs, err
	}
	result = bytes.TrimLeft(result, "(")
	result = bytes.TrimRight(result, ");")
	err = json.Unmarshal(result, info)
	if err != nil {
		return info, cs, fmt.Errorf("request:%v result:%v err:%v", data, string(result), err)
	}
	return info, cs, nil
}

func GetMd5(str string) string {
	token := md5.Sum([]byte(str))
	return hex.EncodeToString(token[:])
}

func main() {
	//url1 := "http://test.account-co.61.com/lvericode/generate"
	url2 := "http://test.account-co.61.com/UserIdentity/authenticate"
	mmidLoginReqData := MMIDLoginRequest{
		Account:  "50026",
		Passwd:   GetMd5("123456"),
		Vericode: 4511,
		Game:     int(694),
	}
	cookies := make([]*http.Cookie, 0, 1)
	cookie := &http.Cookie{
		Name:       "TMACMAIN",
		Value:      "86pc10b28uuq29l4qm9mev8ks1",
		Path:       "/",
		Raw:        "TMACMAIN=86pc10b28uuq29l4qm9mev8ks1; path=/",
	}
	cookies = append(cookies, cookie)
	SDKMMIDLoginRequest(url2, mmidLoginReqData, cookies, "10.1.240.199")
	//str := base64.StdEncoding.EncodeToString(data)
	//fmt.Sprintf("数据： %v", str)
}

