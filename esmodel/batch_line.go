package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type BatchLine struct {
	ID           int    `json:"id"`            // 唯一标识
	ProvinceCode int    `json:"province_code"` // 省份编码
	Province     string `json:"province"`      // 省份
	Year         int    `json:"year"`          // 年份
	Batch        string `json:"batch"`         // 批次
	BatchType    string `json:"batch_type"`    // 批次类型
	Subjects     string `json:"subjects"`      // 科目
	Score        int    `json:"score"`         // 分数
	LineScore    int    `json:"line_score"`    // 压线分
	PassingScore string `json:"passing_score"` // 专业及格分
	MajorScore   string `json:"major_score"`   // 专业分
}

func (BatchLine) Index() string {
	return "batch_line_index"
}

func (BatchLine) Mapping() string {
	return `
{
  "mappings": {
	"properties": {
	  "id": {
		"type": "integer"
	  },
	  "province_code": {
		"type": "integer"
	  },
	  "province": {
		"type": "keyword"
	  },
	  "year": {
		"type": "integer"
	  },
	  "batch": {
		"type": "keyword"
	  },
	  "batch_type": {
		"type": "keyword"
	  },
	  "subjects": {
		"type": "keyword"
	  },
	  "score": {
		"type": "integer"
	  },
	  "line_score": {
		"type": "integer"
	  },
	  "passing_score": {
		"type": "text"
	  },
	  "major_score": {
		"type": "text"
	  }
	}
  }
}
`
}

func (BatchLine) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if k == 0 {
		return nil, fmt.Errorf("invalid k")
	}
	if len(v) < 11 {
		v = append(v, "", "", "")
	}
	provinceCode, _ := strconv.Atoi(v[0])
	year, _ := strconv.Atoi(v[2])
	score, _ := strconv.Atoi(v[6])
	lingScore, _ := strconv.Atoi(v[7])
	return BatchLine{
		ID:           k,
		ProvinceCode: provinceCode,
		Province:     v[1],
		Year:         year,
		Batch:        v[3],
		BatchType:    v[4],
		Subjects:     v[5],
		Score:        score,
		LineScore:    lingScore,
		PassingScore: v[9],
		MajorScore:   v[10],
	}, nil
}

func (b BatchLine) GetId() string {
	return fmt.Sprintf("%d", b.ID)
}
