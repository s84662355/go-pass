package model

import (
	"GoPass/lib/es"
	"GoPass/lib/helper"
	lm "GoPass/logic/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	_ "time"
)

type Article struct {
	lm.Article
	CateName string `json:"cate_name"`
}

func (l Article) GetIndex() string {
	return "article"
}

func (l Article) GetConn() string {
	return "default"
}

func (l Article) ModelW() *elastic.IndexService {
	return es.Es(l.GetConn()).Index().Index(l.GetIndex()).Type("_doc")
}

func (l Article) Model() *elastic.GetService {
	return es.Es(l.GetConn()).Get().Index(l.GetIndex()).Type("_doc")
}

func (l Article) ModelSearch() *elastic.SearchService {
	return es.Es(l.GetConn()).Search().Index(l.GetIndex()).Type("_doc")
}

func (l Article) ModelCount() *elastic.CountService {
	return es.Es(l.GetConn()).Count().Index(l.GetIndex()).Type("_doc")
}

func (l Article) Get(id interface{}) Article {
	service := l.Model()

	get1, err := service.Id(fmt.Sprintf("%s", id)).Do(context.Background())
	var bb Article
	if err != nil {
		panic(err)
	}
	if get1.Found {
		err := json.Unmarshal(get1.Source, &bb)
		if err != nil {
			fmt.Println(err)
		}

	}
	return bb
}

func (l Article) Search(data map[string]interface{}, page int, size int) map[string]interface{} {
	service := l.ModelSearch()
	queryOne := elastic.NewBoolQuery()

	if v, ok := data["title"]; ok {
		queryOne.Must(elastic.NewMultiMatchQuery(v, "title", "content.title", "cate_name", "summary", "content.content"))
	}

	if v, ok := data["cate_id"]; ok && v.(string) != "" && v.(string) != "0" {
		queryOne.Must(elastic.NewTermQuery("cate_id", v))

	}

	if v1, ok1 := data["start_date"]; ok1 {
		if v2, ok2 := data["end_date"]; ok2 {
			dateS := elastic.NewRangeQuery("updated_at")
			dateS.Format("yyyy-MM-dd HH:mm:ss").Gte(v1).Lte(v2)
			queryOne.Must(dateS)
		}
	}

	if v, ok := data["status"]; ok && (v.(string) == "0" || v.(string) == "1") {
		queryOne.Must(elastic.NewTermQuery("status", v))
	}

	count, err := l.ModelCount().Query(queryOne).Do(context.Background())
	if err != nil {
		fmt.Println(err, count)
	}
	paginatorMap := helper.Paginator(page, size, count)
	sourceContext := elastic.NewFetchSourceContext(true)
	sourceContext.Exclude("content")

	res, err := service.
		FetchSourceContext(sourceContext).
		Query(queryOne).
		Size(size).
		From((page-1)*size).
		Sort("updated_at", false).
		Do(context.Background())
	if err != nil {
		fmt.Println(err, res)
	}

	results := make([]Article, 0)

	for _, hit := range res.Hits.Hits {
		//*hit.Source
		var bb Article
		json.Unmarshal(hit.Source, &bb)
		results = append(results, bb)

	}

	///fmt.Println(results)

	return map[string]interface{}{"list": results, "paginator": paginatorMap}
}

func (l Article) PostDataById(id interface{}) (*Article, error) {
	articleModel := lm.Article{}.GetUnscoped(id)
	if articleModel.Id == 0 {
		return nil, errors.New(fmt.Sprintf("id:%d数据不存在", id))
	}
	return l.PostData(articleModel)
}

func (l Article) PostData(articleModel lm.Article) (*Article, error) {
	cate := lm.Cate{}.Get(articleModel.CateId)
	mm := Article{
		articleModel,
		cate.Name,
	}
	_, err := l.ModelW().Id(fmt.Sprintf("%d", articleModel.Id)).
		VersionType("external").
		Version(articleModel.UpdatedAt.Unix()).
		BodyJson(mm).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &mm, nil
}
