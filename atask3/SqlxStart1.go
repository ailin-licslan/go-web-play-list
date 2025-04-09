package main

/**
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// Employee 结构体定义
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"` //涉及到金融领域的精确计算时  可以考虑使用第三方decimal库  比如 github.com/shopspring/decimal  Or 使用字符串 string 转换
}

// 查询部门为“技术部”的员工信息
func queryEmployeesByDepartment(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE department =?"
	err := db.Select(&employees, query, "技术部")
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// 查询工资最高的员工信息
func queryHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
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

	// 查询部门为“技术部”的员工信息
	employees, err := queryEmployeesByDepartment(db)
	if err != nil {
		log.Printf("查询部门为“技术部”的员工信息时出错: %v", err)
	} else {
		log.Printf("部门为“技术部”的员工信息: %+v", employees)
	}

	// 查询工资最高的员工信息
	highestSalaryEmployee, err := queryHighestSalaryEmployee(db)
	if err != nil {
		log.Printf("查询工资最高的员工信息时出错: %v", err)
	} else {
		log.Printf("工资最高的员工信息: %+v", highestSalaryEmployee)
	}
}
