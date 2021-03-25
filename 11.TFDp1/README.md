## Лекция 11. Про тесты

***Проблема*** - тестируем все вручную. Это не хорошо.

### Шаг 0. Термины
***TFD*** - это концпеция, подразумевающая написание модульных тестов еще ДО начала
реализации кода проекта (Test First Development)


### Шаг 1. Простейший пример с факториалом.
* Что оно должно уметь? Уметь вычислять факториал
* Как? ```func factorial(num int) int {}```
* Как проверить, что оно работает правильно?
```
0! = 1
1! = 1
2! = 2
3! = 6
5! = 120
6! = 720
....
Входной параметр меньше 10.
```
* В результате предыдущего пункта имеем ряд ограничений (граничные условия)

* Всегда с самого начала продумывайте граничные условия!!!

* Создаем файл с тестами
  ```main_test.go```
```
package main

import "testing"

type TestCase struct {
	InputData int // то, что будет подаваться на вход
	Answer    int // то, что вернет тестируемая функция
	Expected  int //то, что ожидаем получить
}

//Тестовый сценарий
var Cases []TestCase = []TestCase{
	{
		InputData: 0,
		Expected:  1,
	},
	{
		InputData: 1,
		Expected:  1,
	},
	{
		InputData: 3,
		Expected:  6,
	},
	{
		InputData: 5,
		Expected:  120,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("test case %d failed: result %v expected %v", id, test.Answer, test.Expected)
		}
	}
}

```

### Шаг 2. Реализация factoial()
```main.go```
```
func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	return ans
}

```

### Шаг 3. Запуск.
```go test -v``` (не забывайте про go mod)

### Шаг 4. Теперь попробуем создать простейший http тест
* Что оно должно уметь? Должно уметь вычислять факториал через типичный http запрос
* Как? ```POST /factorial?num=7``` - > 7!
* Как проверить, что оно работает правильно?
```
POST /factorial?num=0 => 1
POST /factorial?num=1 => 1
POST /factorial?num=2 => 2
POST /factorial?num=3 => 6
POST /factorial?num=5 => 120
...
Ограничимся тем, что факториал 10! - самое большое значение, которое в принципе будем вычислять

```

### Шаг 5. Простейший http тест
* Откроем postman и создать тестовый запрос: ```http://localhost:8080/factorial?num=6```
* Теперь добавим тесты на уровне приложения
```
func TestHandleFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandlerFactorial)
	for _, test := range HttpCases {
		//Подтест (суб-тест)
		t.Run(test.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder() // Куда писать ответ
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, err := http.NewRequest("GET", handlerData, nil) //Какой будет запрос
			// data := io.Reader([]byte(`{"num" : 5}`))
			// request, err := http.Post("http://localhost:8080/factorial?num=5", "application/json", data)
			if err != nil {
				t.Error(err)
			}
			handler.ServeHTTP(recorder, request) // Выполняем запрос и ответ записываем в recorder
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("test %s failed: input: %v! result: %v expected %v",
					test.Name,
					test.Numeric,
					string(recorder.Body.Bytes()),
					string(test.Expected),
				)
			}

		}) // Под-тестовый раннер
	}

}
```

* Теперь реализация:
```
package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	dataMap map[int]int
)

func init() {
	dataMap = make(map[int]int)
}

func main() {
	http.HandleFunc("/factorial", HandlerFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	if result, ok := dataMap[num]; ok {
		return result
	}
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	dataMap[num] = ans
	return ans
}

func HandlerFactorial(writer http.ResponseWriter, request *http.Request) {
	//http://localhost:8080/factorial?num=10
	num := request.FormValue("num")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(writer, err.Error(), 404) // msg := .....
		return
	}
	io.WriteString(writer, strconv.Itoa(factorial(n)))
}

```

### Шаг 6. Покрытие
Замер покрытия делаем через ```go test -cover``` (в будущем посмотрите на ```gotool coverage```).
По-хорошему нужно покрыть 70-85% вашего кода. 