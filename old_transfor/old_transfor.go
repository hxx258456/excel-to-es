package old_transfor

import (
	"context"
	"excel-to-es/esmodel"
	"excel-to-es/transfor"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/xuri/excelize/v2"
)

var UniversityMap = make(map[string]string)

func ReadExcel[T transfor.Indexer](esCli *elastic.Client, filepath string, docType T, chunkSize int, flagReverse bool, flagOffset int, ctx context.Context) error {
	exists, err := esCli.IndexExists(docType.Index()).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		_, err := esCli.CreateIndex(docType.Index()).BodyJson(docType.Mapping()).Do(ctx)
		if err != nil {
			return err
		}
	}
	file, err := excelize.OpenFile(filepath)
	if err != nil {
		return err
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		return err
	}

	for i := 0; i < len(rows); i += chunkSize {

		end := i + chunkSize

		// 防止索引越界
		if end > len(rows) {
			end = len(rows)
		}
		chunk := rows[i:end]
		bulk := esCli.Bulk().Index(docType.Index()).Retrier(elastic.NewBackoffRetrier(elastic.NewConstantBackoff(time.Second * 5))).Refresh("true")
		for k, v := range chunk {
			if i != 0 {
				k = k + i
			}
			if k == 0 {
				continue
			}

			if flagReverse {
				universityName := strings.Replace(v[1], " ", "", -1)
				universityName = strings.Replace(universityName, "）", ")", -1)
				universityName = strings.Replace(universityName, "（", "(", -1)
				_, ok := UniversityMap[universityName]
				if !ok {

					res, err := esCli.Search().Index(esmodel.University{}.Index()).Query(elastic.NewTermQuery("name.keyword", universityName)).Do(ctx)
					if err != nil {
						log.Println(err)
						return err
					}
					for _, v := range res.Hits.Hits {
						UniversityMap[universityName] = v.Id
					}
				}
				v[0] = UniversityMap[universityName]

			}
			if flagOffset > 0 {
				k += flagOffset
			}
			doc, err := docType.GenDoc(k, v)
			if err != nil {
				log.Println(err)
				break
			}

			req := elastic.NewBulkUpdateRequest().Id(doc.GetId()).Doc(doc).Upsert(doc)
			bulk.Add(req)
		}
		res, err := bulk.Do(ctx)
		if err != nil {
			return err
		}
		if len(res.Failed()) > 0 {
			for _, v := range res.Failed() {
				fmt.Println(v.Error.Reason, v.Id, v.Result)
			}
		}
	}
	return nil
}
