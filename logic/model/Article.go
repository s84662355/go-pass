package model

import (
    mysqlConfig "GoPass/config/mysql"
    "GoPass/lib/helper"
    "GoPass/lib/mysql"
    "GoPass/lib/redis"
    "fmt"
    "github.com/jinzhu/gorm"
)

var redisClient = redis.GetRedis()

type Article struct {
    Id         uint64              `json:"id"`
    Title      string              `gorm:"type:varchar(100);not null;" json:"title"`   //文章标题
    Content    helper.JsonSqlValue `gorm:"type:longtext;not null;" json:"content"`     //文章内容
    Image      string              `gorm:"type:varchar(200);not null;" json:"image"`   //图片
    Summary    string              `gorm:"type:varchar(500);not null;" json:"summary"` //简短介绍
    CateId     uint32              `gorm:"not null;default:0;index" json:"cate_id"`    //分类id
    ReadAmount uint64              `gorm:"not null;default:0;" json:"read_amount"`     //阅读量
    Sort       uint32              `gorm:"not null;default:0;index" json:"sort"`       //分类id
    Status     uint32              `gorm:"not null;default:1;" json:"status"`          //0待发布 1发布
    CreatedAt  helper.JSONTime     `gorm:"not null;default:current_timestamp" json:"created_at"`
    UpdatedAt  helper.JSONTime     `gorm:"not null;default:current_timestamp; " json:"updated_at"`
    DeletedAt  *helper.JSONTime    `json:"deleted_at"`
}

type ArticleData struct {
    Article
    CateName string `json:"cate_name"`
}

func (m Article) Model() *gorm.DB {
    return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*Article) Connection() string {
    return mysqlConfig.Config.Default
}

func (*Article) TableName() string {
    return "article"
}

func (a Article) Get(id interface{}) Article {
    res := Article{}
    a.Model().Where("id = ?", id).Where("deleted_at is null").First(&res)
    return res
}

func (a Article) GetUnscoped(id interface{}) Article {
    res := Article{}
    a.Model().Where("id = ?", id).Unscoped().First(&res)
    return res
}

func (a Article) Info(id interface{}) ArticleData {
    var results ArticleData
    query := a.Model()

    query.Select("article.* , cate.name as cate_name").Joins("LEFT JOIN cate on article.cate_id = cate.id").Where("article.id = ?", id).Where("article.deleted_at is null").First(&results)
    return results
}

func (a *Article) SetReadAmount() {
    redisClient.IncrBy(fmt.Sprintf("ReadAmount_%d", a.Id), 1)
}

func (a Article) SetReadAmountBy(id interface{}) {
    redisClient.IncrBy(fmt.Sprintf("ReadAmount_%d", id), 1)
}

func (a *Article) UpdateReadAmount(count int) {
    a.UpdatedAt = helper.JSONTime{}.Create()
    a.Model().Model(&a).Where("read_amount = ?", a.ReadAmount).Update("read_amount", a.ReadAmount+uint64(count))
}

func (a *Article) Create() bool {
    a.CreatedAt = helper.JSONTime{}.Create()
    a.UpdatedAt = helper.JSONTime{}.Create()
    a.Model().Create(a)
    if a.Id == 0 {
        return false
    }
    return true
}

func (a *Article) Update() {
    a.UpdatedAt = helper.JSONTime{}.Create()
    a.Model().Save(a)
}

func (a *Article) Delete() {
    deletedAt := helper.JSONTime{}.Create()
    a.DeletedAt = &deletedAt
    a.UpdatedAt = helper.JSONTime{}.Create()
    a.Model().Save(a)
}

func (a Article) List(data map[string]interface{}, page int, size int) map[string]interface{} {
    query := a.Model()
    var count int64 = 0

    query = query.Select(" article.id ,  article.title , article.image ,article. summary,   article.cate_id ,    article.read_amount,    article.sort ,  article.created_at , article.updated_at ,  article.deleted_at ,   article.status,   cate.name as cate_name").Joins("LEFT JOIN cate on article.cate_id = cate.id")

    if v, ok := data["title"]; ok {
        query = query.Where("article.title  like  '%" + v.(string) + "%' ")
    }

    if v, ok := data["cate_id"]; ok && v.(string) != "" && v.(string) != "0" {
        query = query.Where("article.cate_id = ?", v)
    }

    if v1, ok1 := data["start_date"]; ok1 {
        if v2, ok2 := data["end_date"]; ok2 {
            query = query.Where("article.updated_at BETWEEN ? AND ?", v1, v2)
        }
    }

    if v, ok := data["status"]; ok && (v.(string) == "0" || v.(string) == "1") {
        query = query.Where("article.status = ?", v)

    }

    query.Where("article.deleted_at is null")

    var results []ArticleData

    query.Count(&count).Offset((page - 1) * size).Limit(size).Order("article.updated_at  desc").Scan(&results)
    paginatorMap := helper.Paginator(page, size, count)

    //fmt.Println(results)

    return map[string]interface{}{"list": results, "paginator": paginatorMap}
}

func (a Article) GetByCateId(cate_id interface{}) []Article {
    query := a.Model()
    query = query.Select(" article.*").Joins("LEFT JOIN cate on article.cate_id = cate.id")
    query = query.Where("article.cate_id = ?", cate_id)
    var results []Article
    query.Unscoped().Scan(&results)

    return results
}
