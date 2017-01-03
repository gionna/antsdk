package zhima

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 芝麻信用评分普惠版
// 用于商户做准入判断 商户输入准入分 判断用户是否准入
type ZhimaCreditScoreBriefGetRequest struct {
  api.IAlipayRequest
  TerminalType      string                                      `json:"terminal_type"`
  TerminalInfo      string                                      `json:"terminal_info"`
  ProdCode          string                                      `json:"prod_code"`
  NotifyUrl         string                                      `json:"notify_url"`
  ReturnUrl         string                                      `json:"return_url"`
  BizContent        ZhimaCreditScoreBriefGetRequestBizContent   `json:"biz_content"`
}

type ZhimaCreditScoreBriefGetRequestBizContent struct {
  TransactionId   string  `json:"transaction_id"`     // 商户请求的唯一标志，64位长度的字母数字下划线组合。该标识作为对账的关键信息，商户要保证其唯一性，对于用户使用相同transaction_id的查询，芝麻在一天（86400秒）内返回首次查询数据，超过有效期的查询即为无效并返回异常，有效期内的重复查询不重新计费
  ProductCode     string  `json:"product_code"`       // 产品码，直接使用［示例］给出的值
  CertType        string  `json:"cert_type"`          // 证件类型 目前支持三种IDENTITY_CARD(身份证),PASSPORT(护照),ALIPAY_USER_ID(支付宝uid)
  CertNo          string  `json:"cert_no"`            // 对应的证件号(未脱敏)或支付宝uid
  Name            string  `json:"name"`               // 用户姓名 当证件类型为ALIPAY_USER_ID时不需要传入
  AdmittanceScore int     `json:"admittance_score"`   // 350～950之间 业务判断的准入标准 建议业务确定一个稳定的判断标准 频繁的变更该标准可能导致接口被停用
}

func (this *ZhimaCreditScoreBriefGetRequest) GetApiMethodName() string {
  return "zhima.credit.score.brief.get"
}

func (this *ZhimaCreditScoreBriefGetRequest) GetApiVersion() string {
  return "1.0"
}

func (this *ZhimaCreditScoreBriefGetRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *ZhimaCreditScoreBriefGetRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *ZhimaCreditScoreBriefGetRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *ZhimaCreditScoreBriefGetRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *ZhimaCreditScoreBriefGetRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *ZhimaCreditScoreBriefGetRequest) IsNeedEncrypt() bool {
  return false
}

func (this *ZhimaCreditScoreBriefGetRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
