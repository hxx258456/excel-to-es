package transfor

import (
	"context"
	"excel-to-es/reader"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/panjf2000/ants/v2"
	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/mmap"
	"runtime"
	"sync"
	"time"
)

type Indexer interface {
	Index() string
	Mapping() string
	GenDoc(int, []string) (Indexer, error)
	GetId() string
}

type TaskParams[T Indexer] struct {
	esCli      *elastic.Client
	docType    T
	rows       [][]string
	chunkStart int
}

func ReadExcel[T Indexer](esUrl, esUser, esPassword string, filepath string, docType T, chunkSize int, ctx context.Context) error {
	wg := &sync.WaitGroup{}
	pool, err := ants.NewPoolWithFunc(runtime.NumCPU(), func(i interface{}) {

		defer wg.Done()
		param := i.(*TaskParams[T])
		bulk := param.esCli.Bulk().Index(param.docType.Index()).Retrier(elastic.NewBackoffRetrier(elastic.NewConstantBackoff(time.Second * 5))).Refresh("true")

		for k, v := range param.rows {
			if k == 0 {
				continue
			}
			v := v
			doc, err := param.docType.GenDoc(k+param.chunkStart, v)
			if err != nil {
				fmt.Println(err)
				break
			}
			req := elastic.NewBulkUpdateRequest().Id(doc.GetId()).Doc(doc).Upsert(doc)
			bulk.Add(req)
		}
		res, err := bulk.Do(ctx)
		if err != nil {
			fmt.Println(err)
		}
		if len(res.Failed()) > 0 {
			for _, v := range res.Failed() {
				fmt.Println(v.Error.Reason, v.Id, v.Result)
			}
		}
	}, ants.WithPreAlloc(false))
	if err != nil {
		return err
	}
	defer pool.Release()
	esCli, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetBasicAuth(esUser, esPassword), elastic.SetSniff(false))
	if err != nil {
		return err
	}
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
	readAt, err := mmap.Open(filepath)
	if err != nil {
		return err
	}

	read := reader.NewReader(readAt)
	file, err := excelize.OpenReader(read)
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
		wg.Add(1)
		err := pool.Invoke(&TaskParams[T]{
			esCli:      esCli,
			docType:    docType,
			chunkStart: i,
			rows:       chunk,
		})
		if err != nil {
			return err
		}
	}
	wg.Wait()
	return nil
}
