## Оглавление
1. [Описание](#desc)
2. [Инструменты](#tools)
3. [Установка](#install)
4. [Тесты](#tests)

---

<a name="desc"></a>
### 1. Описание

Пакет имеет набор инструментов для удобного написания [табличных тестов](https://en.wikipedia.org/wiki/Data-driven_testing). 

<a name="tools"></a>
### 2. Инструменты

Список инструментов:
* `tools/assistant` - для работы с передачей данных через контекст.
* `tools/inventory` - базовые модели тестов и методы загрузки из JSON.

Список аддонов:
* `addons/behavior` - содержит расширенную модель теста с описанием поведения.

<a name="install"></a>
### 3. Установка
Для получения доступа к основным инструментам, достаточно импортировать пакет в проект:
```
go get github.com/auvitly/lab@latest
```

И импортировать один из инструментов. Например, `inventory`:

```go
import "github.com/auvitly/lab/tools/inventory"
```

<a name="tests"></a>
### 4. Тесты

#### 4.1 Тест функции
Рассмотрим на самом элементарном примере. Ниже представлена функция деления аргументов:
```go
func Multiply(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("сan't divide by zero")
	}
	
	return a/b, nil
}
```

Тест:
```go
//go:embed test
var data 

func TestMultiply(t *testing.T) {
    var tests = inventory.MustLoadTestsFromFS[inventory.Test[
        struct {
            A float64 `json:"a"`
            B float64 `json:"b"`
        },
        struct {
            C     float64          `json:"c"`
            Error *inventory.Error `json:"error"`
        },
    ]](fs, fmt.Sprintf("test/%s.json", t.Name()))

    for i := range tests {
        var test = tests[i]
        
        t.Run(tests[i].Title, func(tt *testing.T) {
            tt.Parallel()
            
            c, err := divide.Divide(test.In.A, test.In.B)
            if err != nil {
                assert.EqualError(tt, err, test.Out.Error.Error(), test.Title)
            
                return
            }
            
            assert.NoError(tt, test.Out.Error.Extract(), test.Title)
            assert.Equal(tt, c, test.Out.C, test.Title)
        })
    }	
}
```

Тестовые данные из файла `test/TestDivide.json`:
```json
[
  {
    "title": "success",
    "in": {
      "a": 1,
      "b": 2
    },
    "out": {
      "c": 0.5
    }
  },
  {
    "title": "error",
    "in": {
      "a": 2,
      "b": 0
    },
    "out": {
      "error": {
        "message": "сan't divide by zero"
      }
    }
  }
]
```

Все данные теста хранятся в файле названием теста, что позволяет сохранить простоту исходного файла 
с тестом и легко найти файл с данными, в случае необходимости. А при большом числе тестов можно разбить каталог `test` на подкаталоги `test/calc/TestDivide.json`.

Если уже имеется подготовленная модель с данными, то тест становится проще. 
Ниже представлен пример для случайного пакета `method` с подготовленными моделями для аргументов и результата.
```go
//go:embed test
var data 

func TestMethod(t *testing.T) {
    var tests = inventory.MustLoadTestsFromFS[inventory.Test[
        inventory.In[method.Arguments],
        inventory.Out[method.Results, *inventory.Error],
    ]](fs, fmt.Sprintf("test/%s.json", t.Name()))

    for i := range tests {
        var test = tests[i]
        
        t.Run(tests[i].Title, func(tt *testing.T) {
            tt.Parallel()
            
            result, err := method.Method(test.In.Arguments)
            if err != nil {
                assert.EqualError(tt, err, test.Out.Error.Error(), test.Title)
            
                return
            }
            
            assert.NoError(tt, test.Out.Error.Extract(), test.Title)
            assert.Equal(tt, result, test.Out.Result, test.Title)
        })
    }	
}
```
<!-- #### 4.2 Тест сущности -->
