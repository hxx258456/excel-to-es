package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type EarlyBatch struct {
	Id                int    `json:"id"`
	Year              int    `json:"year"`                // 年份
	ProvinceCode      int    `json:"province_code"`       // 省份编码
	Province          string `json:"province"`            // 省份
	Title             string `json:"title"`               // title
	UTitle            string `json:"u_title"`             //
	CollegeEnrollCode string `json:"college_enroll_code"` //
	CollegeCode       string `json:"college_code"`        // 院校编码
	CollegeName       string `json:"college_name"`        // 院校名称
	MajorCode         string `json:"major_code"`          //
	MajorName         string `json:"major_name"`          // 专业名称
	EduLevel          string `json:"edu_level"`           // 学历层次
	PlanNum           string `json:"plan_num"`            // 招生人数
	Cost              string `json:"cost"`                //
	LeanYear          string `json:"lean_year"`           //
	Remark            string `json:"remark"`              // 说明
}

func (EarlyBatch) Index() string {
	return "early_batch_index"
}

func (EarlyBatch) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "year": {
        "type": "integer"
      },
      "province_code": {
        "type": "integer"
      },
      "province": {
        "type": "keyword"
      },
      "title": {
        "type": "keyword"
      },
      "u_title": {
        "type": "keyword"
      },
      "college_enroll_code": {
        "type": "keyword"
      },
      "college_code": {
        "type": "keyword"
      },
      "college_name": {
        "type": "keyword"
      },
      "major_code": {
        "type": "keyword"
      },
      "major_name": {
        "type": "keyword"
      },
      "edu_level": {
        "type": "keyword"
      },
      "plan_num": {
        "type": "keyword"
      },
      "cost": {
        "type": "keyword"
      },
      "lean_year": {
        "type": "keyword"
      },
      "remark": {
        "type": "text"
      }
    }
  }
}
`
}

func (EarlyBatch) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if k == 0 {
		return nil, fmt.Errorf("invalid k")
	}
	year, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	provinceCode, err := strconv.Atoi(v[1])
	if err != nil {
		return nil, err
	}
	return EarlyBatch{
		Id:                k,
		Year:              year,
		ProvinceCode:      provinceCode,
		Province:          v[2],
		Title:             v[3],
		UTitle:            v[4],
		CollegeEnrollCode: v[5],
		CollegeCode:       v[6],
		CollegeName:       v[7],
		MajorCode:         v[8],
		MajorName:         v[9],
		EduLevel:          v[10],
		PlanNum:           v[11],
		Cost:              v[12],
		LeanYear:          v[13],
		Remark:            v[14],
	}, nil
}

func (e EarlyBatch) GetId() string {
	return fmt.Sprintf("%d", e.Id)
}
