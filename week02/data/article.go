package data

import (
	"context"
	"homework/week02/biz"
)

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		data: data,
	}
}

type articleRepo struct {
	data *Data
}

func (ar articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	rv := make([]*biz.Article, 0)
	db := ar.data.db.Model(&biz.Article{}).Select("id, title").Find(&rv)
	return rv, db.Error
}
