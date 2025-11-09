package main

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name
// （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB 初始化数据库
func InitDB(models ...interface{}) *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
	return db
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

type Student struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Grade string `gorm:"column:grade"`
}

func q1Demo() {
	InitDB(&Student{})

	db := ConnectDB()

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	fmt.Println(students)
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Where("age < ?", 15).Delete(&Student{})

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

}

// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
// （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
// 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	gorm.Model
	Balance int64
}

type Transactions struct {
	gorm.Model
	From_account_id uint
	To_account_id   uint
	Amount          int64
}

func transactionAmount(db *gorm.DB, accountAId uint, accountBId uint, amount int64) {
	db.Transaction(func(tx *gorm.DB) error {
		var accountA Account
		err := tx.Model(&Account{}).Where("id = ?", accountAId).Where("balance >= ?", amount).Find(&accountA).Error
		if err != nil {
			return err
		}

		tx.Model(&Account{}).Where("id = ?", accountAId).Update("balance", gorm.Expr("balance - ?", amount))
		tx.Model(&Account{}).Where("id = ?", accountBId).Update("balance", gorm.Expr("balance + ?", amount))

		tx.Create(&Transactions{From_account_id: accountAId, To_account_id: accountBId, Amount: amount})

		return nil
	})
}

func q2Demo() {
	InitDB(&Account{}, &Transactions{})

	db := ConnectDB()

	db.Create(&Account{Balance: 20000})
	db.Create(&Account{Balance: 30000})

	transactionAmount(db, 1, 2, 10000)
	var accountA Account
	db.Where("id = ?", 1).Find(&accountA)
	fmt.Println("account A", accountA)
	var accountB Account
	db.Where("id = ?", 2).Find(&accountB)
	fmt.Println("account B", accountB)

	db.Where("id = ?", 2).Find(&accountB)
	var allTransaction Transactions
	db.Find(&allTransaction)
}

// Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：

type Employees struct {
	Id         uint
	Name       string
	Department string
	Salary     int64
}

func ConnectDBbySqlX() *sqlx.DB {
	// 连接 SQLite 数据库（也可以连接 MySQL、PostgreSQL 等）
	db, err := sqlx.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接测试失败:", err)
	}

	fmt.Println("数据库连接成功!")
	return db
}

func CreateEmployeesTable(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		salary INTEGER NOT NULL,
		department TEXT
	);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal("创建 employees 表失败:", err)
	}

	fmt.Println("employees 表已就绪")
}

func q3demo() {
	db := ConnectDBbySqlX()
	CreateEmployeesTable(db)
	var initEmployees []Employees = []Employees{{
		Name: "张三", Department: "技术部", Salary: 10000000,
	}, {
		Name: "lisi", Department: "技术部", Salary: 20000000,
	}}

	for _, emp := range initEmployees {
		query := `INSERT INTO employees (name, salary, department)
					VALUES (:name, :salary, :department)`
		_, err := db.NamedExec(query, emp)
		if err != nil {
			log.Printf("insert emp %s 失败: %v", emp.Name, err)
		}
	}

	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employees []Employees
	db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	fmt.Println(employees)

	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var maxEmployees []Employees
	db.Select(&maxEmployees, "SELECT * FROM employees where salary = (SELECT max(salary) from employees)")
	fmt.Println(maxEmployees)

}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	Id     uint
	Title  string
	Author string
	Price  int64
}

func CreateBookTable(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS book (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT ,
		author TEXT ,
		price INTEGER NOT NULL
	);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal("创建 book 表失败:", err)
	}

	fmt.Println("book 表已就绪")
}

func q4Demo() {
	db := ConnectDBbySqlX()
	CreateBookTable(db)

	var initBooks []Book = []Book{{
		Title: "鲁滨逊", Author: "张三", Price: 10000,
	}, {
		Title: "飘", Author: "lisi", Price: 200000,
	}}

	for _, book := range initBooks {
		query := `INSERT INTO book (title, author, price)
					VALUES (:title, :author, :price)`
		_, err := db.NamedExec(query, book)
		if err != nil {
			log.Printf("insert emp %s 失败: %v", book.Title, err)
		}
	}
	var books []Book
	db.Select(&books, "SELECT * from book where price > ?", 50000)
	fmt.Println(books)

}

// 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func main() {
	fmt.Println("main go")
	//q1Demo()
	//q2Demo()
	//q3demo()
	q4Demo()
}
