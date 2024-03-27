package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

// CareerCTypeCount 职业招聘公司类型统计
type CareerCTypeCount struct {
	ID            int     `json:"id"`              //
	CareerCode    string  `json:"career_code"`     // 职业编码
	CareerName    string  `json:"career_name"`     // 职业名称
	CompanyType   string  `json:"company_type"`    // 企业类型
	Number        int     `json:"number"`          // 企业数量
	RecPercent    float64 `json:"rec_percent"`     // 招聘占比
	RecPercentStr string  `json:"rec_percent_str"` // 招聘占比字符串
	RecNumber     int     `json:"rec_number"`      // 招聘人数
	RecSum        int     `json:"rec_sum"`         // 招聘总人数
	IsUp          string  `json:"is_up"`           // 是否上升
}

func (c CareerCTypeCount) GetId() string {
	return fmt.Sprintf("%d", c.ID)
}

func (CareerCTypeCount) Index() string {
	return "career_company_type_index"
}

func (CareerCTypeCount) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "career_code": {
        "type": "keyword"
      },
      "career_name": {
        "type": "keyword"
      },
      "company_type": {
        "type": "keyword"
      },
      "number": {
        "type": "integer"
      },
      "rec_percent": {
        "type": "double"
      },
      "rec_percent_str": {
        "type": "keyword"
      },
      "rec_number": {
        "type": "integer"
      },
      "rec_sum": {
        "type": "integer"
      },
      "is_up": {
        "type": "keyword"
      }
    }
  }
}
`
}

func (CareerCTypeCount) GenDoc(k int, v []string) (transfor.Indexer, error) {
	number, _ := strconv.Atoi(v[3])
	prePercent, _ := strconv.ParseFloat(v[4], 64)
	recNumber, _ := strconv.Atoi(v[6])
	recSum, _ := strconv.Atoi(v[7])
	return &CareerCTypeCount{
		ID:            k,
		CareerCode:    v[0],
		CareerName:    v[1],
		CompanyType:   v[2],
		Number:        number,
		RecPercent:    prePercent,
		RecPercentStr: v[5],
		RecNumber:     recNumber,
		RecSum:        recSum,
		IsUp:          v[8],
	}, nil
}
