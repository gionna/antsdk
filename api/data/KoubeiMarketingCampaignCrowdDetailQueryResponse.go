package data

import (
  "github.com/gionna/antsdk/api"
)

type KoubeiMarketingCampaignCrowdDetailQueryResponse struct {
  api.AlipayResponse
  CrowdGroupInfo  string `json:"crowd_group_info"`  // op:表示操作符，目前支持的有EQ相等,GT大于,GTEQ大于等于,LT小于,LTEQ小于等于,NEQ不等,LIKE模糊匹配,IN在枚举范围内,NOTIN不在枚举范围内,BETWEEN范围比较,LEFTDAYS几天以内,RIGHTDAYS几天以外,LOCATE地理位置比较,LBS地图位置数据 tagCode:就是标签code，详细标签信息参见附件标签信息 value：标签对应的值
  Name            string `json:"name"`              // 人群名称
}
