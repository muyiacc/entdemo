package main

import (
	"context"
	"entdemo/ent"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// 连接数据库
	client, err := ent.Open("sqlite3", "file:rargb_db.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to sqlite: %v", err)
	}
	defer client.Close()

	// 在 context.Background() 中执行查询
	ctx := context.Background()

	// 查询 items 表中的记录数
	count, err := client.Item.
		Query().Count(ctx)
	if err != nil {
		log.Fatalf("failed querying item count: %v", err)
	}

	log.Printf("Number of items: %d", count)
}
