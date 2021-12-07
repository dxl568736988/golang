package main

import (
	"context"
	"fmt"
	"homework/week02/biz"
	"homework/week02/data"
	"homework/week02/service"
)

/*
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：dao层不应该wrap这个error, 应该直接往上层抛。因为dao层代码可能是复用的,如果wrap了error,
   如果上层接着wrap 那么就会出现多次调用栈信息。像sql.ErrNoRows这类错误(类似于io.EOF),跟传统的error不太一样，
   上层代码可能会利用这个Sentinel error做一些特殊的业务逻辑, 对错误进行降级处理。所以我觉得应该在service层wrap error。
*/

func main() {
	// 深刻体会到 wire 的便捷之处
	db := data.NewData()
	// dao层
	dao := data.NewArticleRepo(db)
	// biz层
	article := biz.NewArticleUsecase(dao)
	// service层
	blog := service.NewBlogService(article)
	//模拟调用service层ListArticle接口
	ps := blog.ListArticle(context.Background())
	fmt.Println(ps)
}