package managers

import (
	"FizzBuzzService/models/request"
	"log"
	"strconv"
)

func FizzBuzz(a, b int, strA, strB string, limit int) ([]string, error) {
	var result []string

	log.Println("Buzzing...")

	for i := 1; i <= limit; i++ {
		match := ""
		if i%a == 0 {
			match = match + strA
		}
		if i%b == 0 {
			match = match + strB
		}
		if match == "" {
			match = strconv.FormatInt(int64(i), 10)
		}
		result = append(result, match)
	}

	created, err := request.FindOrCreate(a, b, strA, strB, limit)
	if err != nil {
		return result, err
	}

	created.Count += 1
	_, err = created.Save()
	if err != nil {
		return result, err
	}

	return result, nil
}
