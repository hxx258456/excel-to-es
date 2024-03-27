package esmodel

import (
	"excel-to-es/transfor"
	"strconv"
)

type University struct {
	Code           string  `json:"code"`            //编码唯一
	Name           string  `json:"name"`            //名字
	Province       string  `json:"province"`        //省份
	City           string  `json:"city"`            //城市
	Category       string  `json:"category"`        //类别
	Nature         string  `json:"nature"`          //性质
	Belong         string  `json:"belong"`          //隶属
	Feature        string  `json:"feature"`         //特色
	Ranking        int     `json:"ranking"`         //排名
	CompositeIndex float32 `json:"composite_index"` //排名
	Heat           int     `json:"heat"`            //热度
	Description    string  `json:"description"`     //简介
	Logo           string  `json:"logo"`            //logo
	Batch          string  `json:"batch"`           //批次
}

func (University) Index() string {
	return "university_index"
}

func (University) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "code": {
        "type": "keyword"
      },
      "name": {
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
      "province": {
        "type": "keyword"
      },
      "batch": {
        "type": "keyword"
	  },
      "city": {
        "type": "keyword"
      },
      "category": {
        "type": "keyword"
      },
      "nature": {
        "type": "keyword"
      },
      "belong": {
        "type": "keyword"
      },
      "feature": {
        "type": "text",
        "analyzer": "comma",
        "search_analyzer": "comma"
      },
      "ranking": {
        "type": "integer"
      },
      "composite_index": {
        "type": "float"
      },
      "heat": {
        "type": "integer"
      },
      "description": {
        "type": "text"
      }
    }
  },
  "settings": {
    "analysis": {
        "analyzer": {
          "comma": {
            "type": "pattern",
            "pattern": ","
          }
        }
      }
  }
}
`
}

func (University) GenDoc(k int, v []string) (transfor.Indexer, error) {
	ranking, _ := strconv.Atoi(v[10])
	compositeIndex, _ := strconv.ParseFloat(v[11], 32)
	heat, _ := strconv.Atoi(v[12])
	return &University{
		Code:           v[0],
		Name:           v[1],
		Logo:           v[2],
		Batch:          v[3],
		Province:       v[4],
		City:           v[5],
		Category:       v[6],
		Nature:         v[7],
		Belong:         v[8],
		Feature:        v[9],
		Ranking:        ranking,
		CompositeIndex: float32(compositeIndex),
		Heat:           heat,
		Description:    v[13],
	}, nil
}

func (u University) GetId() string {
	return u.Code
}
