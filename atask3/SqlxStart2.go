package main

/**
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
import (
	_ "github.com/go-sql-driver/mysql" //无法连接数据库: sql: unknown driver "mysql" (forgotten import?)
	"github.com/jmoiron/sqlx"
	"log"
)

// Book 结构体定义，与books表字段对应
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 查询价格大于58元的书籍
func queryBooksByPrice(db *sqlx.DB) ([]Book, error) {
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price >?"
	err := db.Select(&books, query, 58)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	// 假设已经连接好数据库，这里只是示例，实际需要正确配置连接信息
	db, err := sqlx.Open("mysql", "root:123456@tcp(192.168.0.155:3306)/db1")
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("无法ping通数据库: %v", err)
	}

	// 查询价格大于58元的书籍
	books, err := queryBooksByPrice(db)
	if err != nil {
		log.Printf("查询价格大于58元的书籍时出错: %v", err)
	} else {
		log.Printf("价格大于58元的书籍信息: %+v", books)
	}
}
