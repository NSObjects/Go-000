package main

import (
	"database/sql"

	"github.com/opentracing/opentracing-go/log"

	xerrors "github.com/pkg/errors"
)

func main() {
	book := NewDaoBook()
	if _, err := book.GetBook(); xerrors.Is(err, sql.ErrConnDone) {
		log.Error(err)
	}
}

type DaoBook struct {
	db SqlDriver
}

func NewDaoBook() DaoBook {
	return DaoBook{}
}

//dao 层操作数据库， 捕获的是原始错误
//应该warp后返回， 方便定位到错误的具体位置
func (b DaoBook) GetBook() (Book, error) {
	if err := b.db.Exec("SELECT * FROM book"); err != nil {
		return Book{}, xerrors.Wrap(err, "db error")
	}
	return Book{}, nil
}

type Book struct {
	Id    string
	Name  string
	Title string
}

type SqlDriver interface {
	Exec(string string) error
}
