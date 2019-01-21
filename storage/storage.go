package storage

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"FizzBuzzService/config"
)


type storage struct {
	db *gorm.DB
}

var (
	s    *storage
	once sync.Once
)

type Storager interface {
	Close()
	CreateTable(tableSpec interface{})
	Select(record interface{}) interface{}
	Insert(record interface{})
	Save(record interface{})
	Exec(query string)
	Update(record interface{}) interface{}
	Delete(record interface{})
	AutoMigrate(model interface{})
	Find(model interface{}) interface{}
	FindMax(model interface{}, field string) interface{}
	FindAllByPattern(model interface{}, pattern interface{}) interface{}
	FindBy(model interface{}, field string, value interface{}) interface{}
	ReplaceMember(record interface{}, field string, newMember interface{})
	FindAll(model interface{}, field string, value interface{}) interface{}
	FindByPattern(model interface{}, pattern interface{}) interface{}
}

func Storage() Storager {
	once.Do(func() {
		db := InitDB()
		s = &storage{
			db: db,
		}
	})

	return s
}

func (s *storage) Close() {
	s.db.Close()
}

func (s *storage) Exec(query string) {
	s.db.Exec(query)
}


func (s *storage) AutoMigrate(model interface{}){
	s.db.AutoMigrate(model)
}


func (s *storage) FindBy(model interface{}, field string, value interface{}) interface{} {
	s.db.Where(field+" = ?", value).First(model)
	return model
}

func (s *storage) FindAll(model interface{}, field string, value interface{}) interface{} {
	s.db.Where(field+" = ?", value).Find(model)
	return model
}

func (s *storage) Find(model interface{}) interface{} {
	s.db.Find(model)
	return model
}

func (s *storage) FindAllByPattern(model interface{}, pattern interface{}) interface{} {
	s.db.Where(pattern).Find(model)
	return model
}

func (s *storage) FindByPattern(model interface{}, pattern interface{}) interface{} {
	s.db.Where(pattern).First(model)
	return model
}

func (s *storage) CreateTable(tableSpec interface{}) {
	s.db.CreateTable(tableSpec)
}

func (s *storage) Select(record interface{}) interface{}{
	s.db.Select(record)
	return record
}

func (s *storage) Insert(record interface{}) {
	s.db.Create(record)
}

func (s *storage) ReplaceMember(record interface{}, field string, newMember interface{}){
	s.db.Model(record).Association(field).Replace(newMember)
}

func (s *storage) Save(record interface{}) {
	s.db.Save(record)
}

func (s *storage) Update(record interface{}) (interface{}) {
	s.db.Update(record)
	return record
}


func (s *storage) Delete(record interface{}) {
	s.db.Delete(record)
}

func (s *storage) FindMax(model interface{}, field string) interface{} {
	s.db.Order(field + " desc").First(model)
	return model
}


func InitDB() *gorm.DB {
	user := config.GetDatabaseUser()
	password := config.GetDatabasePassword()
	dbName := config.GetDatabaseName()
	server := config.GetDatabaseUrl()
	odb, err := gorm.Open("mysql", user + ":" + password + "@tcp(" + server + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Fatal("Could not init DB", err)
	}
	return odb
}
