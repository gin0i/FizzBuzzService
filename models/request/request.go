package request

import (
	"FizzBuzzService/storage"
	"errors"
	"time"
	"log"
)

type Request struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Count     int        `json:"count"`
	MulA      int        `json:"int1"`
	MulB      int        `json:"int2"`
	StrA      string     `json:"str1"`
	StrB      string     `json:"str2"`
	Limit     int        `json:"limit"`
}

func DeleteAll() {
	storage.Storage().Delete(&Request{})
}

func Create(mulA, mulB int, strA, strB string, limit int) (*Request, error) {
	created, err := (&Request{
		MulA:  mulA,
		MulB:  mulB,
		StrA:  strA,
		StrB:  strB,
		Limit: limit,
	}).Save()

	return created, err
}

func FindOrCreate(mulA, mulB int, strA, strB string, limit int) (*Request, error) {
	findConditions := Request{
		MulA: mulA,
		MulB: mulB,
		StrA: strA,
		StrB: strB,
		Limit: limit,
	}
	r := FindByPattern(findConditions)
	if r == nil {
		newCapa, err := Create(mulA, mulB, strA, strB, limit )
		if err != nil {
			log.Println("Failed to create request record", mulA, mulB, strA, strB, limit)
		}

		return newCapa, err
	}

	return r, nil
}

func (req *Request) Save() (*Request, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	storage.Storage().Save(req)

	savedRequest := FindByID(req.ID)
	if savedRequest == nil {
		return nil, errors.New("Failed to create req")
	}
	return savedRequest, nil
}

func FindByID(id uint) *Request {
	req := Request{
		ID: id,
	}
	found := FindByPattern(req)
	return found
}

func FindWithHighestCount() *Request {
	req := &Request{}
	storage.Storage().FindMax(req, "count")
	if req.ID == 0 {
		req = nil
	}
	return req
}

func FindByPattern(pattern Request) *Request {
	req := &Request{}
	storage.Storage().FindByPattern(req, pattern)
	if req.ID == 0 {
		req = nil
	}
	return req
}

func (req *Request) Validate() error {
	if req.MulA == 0 || req.MulB == 0 {
		return errors.New("Invalid multiple value 0")
	}

	return nil
}

func (req *Request) Reload() error {
	newRequest := FindByID(req.ID)

	if newRequest == nil {
		return errors.New("Failed to reload")
	}

	req.ID = newRequest.ID
	req.MulB = newRequest.MulB
	req.MulA = newRequest.MulA
	req.StrB = newRequest.StrB
	req.StrA = newRequest.StrA
	req.CreatedAt = newRequest.CreatedAt
	req.UpdatedAt = newRequest.UpdatedAt
	return nil
}
