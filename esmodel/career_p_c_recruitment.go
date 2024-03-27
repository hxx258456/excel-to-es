package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type CareerPCRecruitment struct {
	ID          int     `json:"id"`           //
	CareerCode  string  `json:"career_code"`  // 职业编码
	CareerName  string  `json:"career_name"`  // 职业名称
	Province    string  `json:"province"`     // 省份
	Number      int     `json:"number"`       // 招聘人数
	Percent     float64 `json:"percent"`      // 招聘占比
	PercentStr  string  `json:"percent_str"`  // 招聘占比信息字符串
	MinSalary   int     `json:"min_salary"`   // 最低薪资
	MaxSalary   int     `json:"max_salary"`   // 最高薪资
	SalaryRange string  `json:"salary_range"` // 薪资范围
}

func (r CareerPCRecruitment) GetId() string {
	return fmt.Sprintf("%d", r.ID)
}

func (CareerPCRecruitment) Index() string {
	return "career_p_c_recruitment_index"
}

func (CareerPCRecruitment) Mapping() string {
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
      "province": {
        "type": "keyword"
      },
      "number": {
        "type": "integer"
      },
      "min_salary": {
        "type": "integer"
      },
      "max_salary": {
        "type": "integer"
      },
      "percent": {
        "type": "double"
      },
      "percent_str": {
        "type": "keyword"
      },
      "salary_range": {
        "type": "keyword"
      }
    }
  }
}
`
}

func (CareerPCRecruitment) GenDoc(k int, v []string) (transfor.Indexer, error) {
	number, _ := strconv.Atoi(v[3])
	percent, _ := strconv.ParseFloat(v[4], 64)
	minSalary, _ := strconv.Atoi(v[7])
	maxSalary, _ := strconv.Atoi(v[8])
	return &CareerPCRecruitment{
		ID:          k,
		CareerCode:  v[0],
		CareerName:  v[1],
		Province:    v[2],
		Number:      number,
		Percent:     percent,
		PercentStr:  v[5],
		MinSalary:   minSalary,
		MaxSalary:   maxSalary,
		SalaryRange: v[8],
	}, nil
}
