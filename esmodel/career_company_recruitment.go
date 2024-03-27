package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
	"time"
)

type CareerCompanyRecruitment struct {
	ID          int        `json:"id"`           //
	CareerCode  string     `json:"career_code"`  // 职业编码
	CareerName  string     `json:"career_name"`  // 职业名称
	Office      string     `json:"office"`       // 职位
	Province    string     `json:"province"`     // 省份
	EduLevel    string     `json:"edu_level"`    // 学历
	MinSalary   int        `json:"min_salary"`   // 最低薪资
	MaxSalary   int        `json:"max_salary"`   // 最高薪资
	Salary      string     `json:"salary"`       // 薪资信息
	SalaryUnit  string     `json:"salary_unit"`  // 薪资单位
	Number      int        `json:"number"`       // 招聘人数
	Welfare     string     `json:"welfare"`      // 福利待遇
	MajorNeeds  string     `json:"major_needs"`  // 专业要求
	CompanyName string     `json:"company_name"` // 企业名称
	CompanyType string     `json:"company_type"` // 企业类型
	CompanySize string     `json:"company_size"` // 企业规模
	Origin      string     `json:"origin"`       // 招聘来源
	PublishTime *time.Time `json:"publish_time"` // 发布时间
}

func (r CareerCompanyRecruitment) GetId() string {
	return fmt.Sprintf("%d", r.ID)
}

func (CareerCompanyRecruitment) Index() string {
	return "career_company_recruitment_index"
}

func (CareerCompanyRecruitment) Mapping() string {
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
      "office": {
        "type": "keyword"
      },
      "province": {
        "type": "keyword"
      },
      "edu_level": {
        "type": "keyword"
      },
      "min_salary": {
        "type": "integer"
      },
      "max_salary": {
        "type": "integer"
      },
      "number": {
        "type": "integer"
      },
      "welfare": {
        "type": "keyword"
      },
      "major_needs": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_max_word",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "company_name": {
        "type": "keyword"
      },
      "company_type": {
        "type": "keyword"
      },
      "company_size": {
        "type": "keyword"
      },
      "origin": {
        "type": "keyword"
      },
      "publish_time": {
        "type":   "date"
      }
    }
  }
}
`
}

func (CareerCompanyRecruitment) GenDoc(k int, v []string) (transfor.Indexer, error) {
	minSalary, _ := strconv.Atoi(v[5])
	maxSalary, _ := strconv.Atoi(v[6])
	number, _ := strconv.Atoi(v[9])
	publishTime, _ := time.Parse(time.DateTime, v[16])
	return &CareerCompanyRecruitment{
		ID:          k,
		CareerCode:  v[0],
		CareerName:  v[1],
		Office:      v[2],
		Province:    v[3],
		EduLevel:    v[4],
		MinSalary:   minSalary,
		MaxSalary:   maxSalary,
		Salary:      v[7],
		SalaryUnit:  v[8],
		Number:      number,
		Welfare:     v[10],
		MajorNeeds:  v[11],
		CompanyName: v[12],
		CompanyType: v[13],
		CompanySize: v[14],
		Origin:      v[15],
		PublishTime: &publishTime,
	}, nil
}
