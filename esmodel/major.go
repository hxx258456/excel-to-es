package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type Major struct {
	ID                int    `json:"id"`
	Course            string `json:"course"`              // 学科
	CourseCode        string `json:"course_code"`         // 学科编码
	Category          string `json:"category"`            // 学类
	CategoryCode      string `json:"category_code"`       // 学类编码
	Major             string `json:"major"`               // 专业名
	MajorCode         string `json:"major_code"`          // 专业编码
	Batch             string `json:"batch"`               // 学历层次
	Hot               int    `json:"hot"`                 // 热度
	CourseDuration    int    `json:"course_duration"`     // 修业年限
	CourseDurationStr string `json:"course_duration_str"` // 修业年限中文
	Degree            string `json:"degree"`              // 学位
	WomanRatio        int    `json:"woman_ratio"`         // 女生比例
	ManRatio          int    `json:"man_ratio"`           // 男生比例
	Description       string `json:"description"`         // 专业介绍
	TrainingTarget    string `json:"training_target"`     // 培养目标
	TrainingRequire   string `json:"training_require"`    // 培养要求
	CourseRequire     string `json:"course_require"`      // 学科要求
	Capacity          string `json:"capacity"`            // 知识能力
	ExamDirection     string `json:"exam_direction"`      // 考研方向
	MajorPrograms     string `json:"major_programs"`      // 主要课程
	Celebrity         string `json:"celebrity"`           // 社会名人
}

func (m Major) GetId() string {
	return fmt.Sprintf("%d", m.ID)
}

func (Major) Index() string {
	return "major_index"
}

func (Major) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "course": {
        "type": "keyword"
      },
      "course_code": {
        "type": "keyword"
      },
      "category": {
        "type": "keyword"
      },
      "category_code": {
        "type": "integer"
      },
      "major": {
        "type": "keyword"
      },
      "major_code": {
        "type": "keyword"
      },
      "batch": {
        "type": "keyword"
      },
      "hot": {
        "type": "integer"
      },
      "course_duration": {
        "type": "integer"
      },
      "course_duration_str": {
        "type": "keyword"
      },
      "degree": {
        "type": "keyword"
      },
      "woman_ratio": {
        "type": "integer"
      },
      "man_ratio": {
        "type": "integer"
      },
      "description": {
        "type": "text"
      }
    }
  }
}
`
}

func (Major) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if len(v) < 21 {
		fill := make([]string, 21-len(v))
		v = append(v, fill...)
	}

	hot, _ := strconv.Atoi(v[7])
	courseDuration, _ := strconv.Atoi(v[8])
	womanRatio, _ := strconv.Atoi(v[11])
	manRatio, _ := strconv.Atoi(v[12])
	return &Major{
		ID:                k,
		Course:            v[0],
		CourseCode:        v[1],
		Category:          v[2],
		CategoryCode:      v[3],
		Major:             v[4],
		MajorCode:         v[5],
		Batch:             v[6],
		Hot:               hot,
		CourseDuration:    courseDuration,
		CourseDurationStr: v[9],
		Degree:            v[10],
		WomanRatio:        womanRatio,
		ManRatio:          manRatio,
		Description:       v[13],
		TrainingTarget:    v[14],
		TrainingRequire:   v[15],
		CourseRequire:     v[16],
		Capacity:          v[17],
		ExamDirection:     v[18],
		MajorPrograms:     v[19],
		Celebrity:         v[20],
	}, nil
}
