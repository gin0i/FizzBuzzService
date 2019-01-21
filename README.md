
# FizzBuzzService

## Le service necessite une base MySQL a indiquer via des variables d'environnement (voir .env.example)

1. Le service necessite les packages suivant:
 - github.com/gorilla
 - github.com/codegangsta/negroni
 - github.com/jinzhu/gorm
  
 
2. Les tests necessitent les packages suivants:
 - github.com/stretchr/testify/assert
 - github.com/stretchr/testify/suite
  
  
3. Usage: 
```go run main.go```

4. Routes:
 - ```POST``` /fizzbuzz/api/launch
   JSON:
   ```
   {
    "int1": 5,
    "int2": 3,
    "limit": 401,
    "str1": "Yvain",
    "str2": "Gauvain"
   }
   ````
 - ```GET``` /fizzbuzz/api/stat
