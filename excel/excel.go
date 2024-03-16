package excel

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

type School struct {
	ID             int     `json:"id"`              // es id
	Code           string  `json:"code"`            //编码
	Name           string  `json:"name"`            //名字
	Province       string  `json:"province"`        //省份
	City           string  `json:"city"`            //城市
	Category       string  `json:"category"`        //类别
	Nature         string  `json:"nature"`          //性质
	Belong         string  `json:"belong"`          //隶属
	Feature        string  `json:"feature"`         //特色
	Ranking        int     `json:"ranking"`         //排名
	CompositeIndex float32 `json:"composite_index"` //排名
	Heat           int     `json:"heat"`            // 热度
	Description    string  `json:"description"`     // 简介
}

func (School) Index() string {
	return "school_index"
}

func (School) Mapping() string {
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
        "search_analyzer": "ik_max_word"
      },
      "province": {
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

func ReadExcel(esCli *elastic.Client, filepath string, ctx context.Context) {
	exists, err := esCli.IndexExists(School{}.Index()).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := esCli.CreateIndex(School{}.Index()).BodyJson(School{}.Mapping()).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	log.Println(exists)
	file, err := excelize.OpenFile(filepath)
	if err != nil {
		panic(err)
	}
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}
	for k, v := range rows {
		if k == 0 {
			continue
		}
		ranking, err := strconv.Atoi(v[8])
		if err != nil {
			panic(err)
		}
		compositeIndex, err := strconv.ParseFloat(v[9], 32)
		if err != nil {
			panic(err)
		}
		heat, err := strconv.Atoi(v[10])
		if err != nil {
			panic(err)
		}
		doc := &School{
			ID:             k,
			Code:           v[0],
			Name:           v[1],
			Province:       v[2],
			City:           v[3],
			Category:       v[4],
			Nature:         v[5],
			Belong:         v[6],
			Feature:        v[7],
			Ranking:        ranking,
			CompositeIndex: float32(compositeIndex),
			Heat:           heat,
			Description:    v[11],
		}
		query := elastic.NewTermQuery("code", doc.Code)
		result, err := esCli.Search(doc.Index()).Query(query).Do(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		if result.TotalHits() <= 0 {
			_, err := esCli.Index().Index(doc.Index()).BodyJson(doc).Do(ctx)
			if err != nil {
				log.Println(err)
				//continue
			}
		} else {
			log.Println("更新", result.Hits.Hits[0].Id)
			_, err := esCli.Update().Index(doc.Index()).Id(result.Hits.Hits[0].Id).Doc(doc).Do(ctx)
			if err != nil {
				log.Println(err)
				//continue
			}
		}
		esCli.Flush()
	}
}
