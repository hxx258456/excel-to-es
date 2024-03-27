package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type CareerCompanyCount struct {
	ID          int    `json:"id"`           //
	CareerCode  string `json:"career_code"`  // 职业编码
	CareerName  string `json:"career_name"`  // 职业名称
	CompanyName string `json:"company_name"` // 企业名称
	Number      int    `json:"number"`       // 招聘人数
	IsUp        string `json:"is_up"`        // 是否上升
}

func (c CareerCompanyCount) GetId() string {
	return fmt.Sprintf("%d", c.ID)
}

func (CareerCompanyCount) Index() string {
	return "career_company_count_index"
}
func (CareerCompanyCount) Mapping() string {
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
      "company_name": {
        "type": "keyword"
      },
      "number": {
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

func (CareerCompanyCount) GenDoc(k int, v []string) (transfor.Indexer, error) {
	number, _ := strconv.Atoi(v[3])
	return &CareerCompanyCount{
		ID:          k,
		CareerCode:  v[0],
		CareerName:  v[1],
		CompanyName: v[2],
		Number:      number,
		IsUp:        v[4],
	}, nil
}
