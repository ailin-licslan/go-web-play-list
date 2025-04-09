package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/**
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键，
from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A
扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

// 定义转账函数
func transfer(db *sql.DB, fromAccountID, toAccountID int) error {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	// 检查转出账户余额是否足够
	var balance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id =?", fromAccountID).Scan(&balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("转出账户 %d 不存在", fromAccountID)
		}
		return err
	}

	//判断余额
	if balance < 100 {
		return fmt.Errorf("转出账户 %d 余额不足", fromAccountID)
	}

	// 从账户A扣除100元
	_, err = tx.Exec("UPDATE accounts SET balance = balance - 100 WHERE id =?", fromAccountID)
	if err != nil {
		return err
	}

	// 向账户B增加100元
	_, err = tx.Exec("UPDATE accounts SET balance = balance + 100 WHERE id =?", toAccountID)
	if err != nil {
		return err
	}

	// 在transactions表中记录转账信息
	_, err = tx.Exec("INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?,?,?)", fromAccountID, toAccountID, 100)
	if err != nil {
		return err
	}

	// 提交事务
	return tx.Commit()
}

func main() {
	// 连接数据库  root:123456@(192.168.0.155:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.155:3306)/db1")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// 测试转账
	fromAccountID := 1
	toAccountID := 2
	err = transfer(db, fromAccountID, toAccountID)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
}
