package payjs

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type PayJS struct {
	mchid   string
	privKey string
	apiUrl  string
}

var (
	APIURL = `https://payjs.cn/api/native`
)

type TradeParam struct {
	Total_fee    int    `json:"total_fee"`
	Out_trade_no string `json:"out_trade_no"`
	Body         string `json:"body"`
	Notify_url   string `json:"notify_url"`
}

func New(mchid string, privKey string) *PayJS {
	pj := &PayJS{}
	pj.apiUrl = APIURL
	pj.mchid = mchid
	pj.privKey = privKey
	return pj
}

func sign(params url.Values, privKey string) string {
	params.Del(`sign`)
	var keys = make([]string, 0, 0)
	for key := range params {
		if params.Get(key) != `` {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	
	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(params.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var src = strings.Join(pList, "&")
	src += "&key=" + privKey
	
	md5bs := md5.Sum([]byte(src))
	md5res := hex.EncodeToString(md5bs[:])
	return strings.ToUpper(md5res)
}

func (pj *PayJS) CreateTrade(param TradeParam) (res string, err error) {
	var p = url.Values{}
	jsonbs, _ := json.Marshal(param)
	jsonmap := make(map[string]interface{})
	json.Unmarshal(jsonbs, &jsonmap)
	for k, v := range jsonmap {
		p.Add(k, fmt.Sprintf("%v", v))
	}
	p.Add("mchid", pj.mchid)
	
	p.Add("sign", sign(p, pj.privKey))
	
	cli := http.Client{}
	r, err := cli.PostForm(pj.apiUrl, p)
	if err != nil {
		return ``, err
	}
	defer r.Body.Close()
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ``, err
	}
	return string(bs), nil
}
