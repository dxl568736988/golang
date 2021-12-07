package service

import (
	"context"
	"database/sql"
	"errors"
	"homework/week02/biz"
)

func NewBlogService(article *biz.ArticleUsecase) *BlogService {
	return &BlogService{
		article: article,
	}
}

type BlogService struct {
	article *biz.ArticleUsecase
}

func (s *BlogService) ListArticle(ctx context.Context) []*biz.Article {
	ps, err := s.article.List(ctx)
	// 伪代码 ...
	if errors.Is(err, sql.ErrNoRows){
		// 可以做特殊业务逻辑处理, 如初始化后 降级返回...
		return []*biz.Article{
			{
				ID: 100,
				Title: "初始化",
			},
		}
	}
	return ps
}
