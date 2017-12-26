# payjs
payjs.cn
Go版 SDK by trysh

```go

func main() {
    prikey := ``//通信密钥
    mchid := ``商户号
    pj := payjs.New(mchid, prikey)//新建一个引擎
    res, err := pj.CreateTrade(payjs.TradeParam{    //创建一个扫码支付
        Total_fee:    1,    //1分钱
        Out_trade_no: time.Now().Format("test20060102_150405.999999999"),
        Body:         `测试的标题`,
        Notify_url:   `https://xxx.com/callback`,
    })
    if err != nil {
        log.Println(`payjs CreateTrade err`, err)
        return
    }
    /*返回示例:
    {"code_url":"weixin://wxpay/bizpayurl?pr=xxxxx",
    "out_trade_no":"xxx",
    "payjs_order_id":"2017xxx",
    "return_code":1,
    "return_msg":"SUCCESS",
    "total_fee":"1",
    "sign":"xxx"}
    */
}

```

