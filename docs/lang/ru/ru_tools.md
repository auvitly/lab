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

Используя `tools/inventory` составим конструкцию для теста на основе функции `MustRun`.
Структура теста является композицией двух интерфейсов: `InPlaceholder`, `OutPlaceholder`. Этим интерфейсам
отвечают соответствующие структуры пакета `In` и `Out`. Также существует пустой заполнитель `Empty`, который можно
применить вместо `InPlaceholder` или `OutPlaceholder`. 

Тест:

```go
//go:embed test
var fs embed.FS

func TestDivide(t *testing.T) {
    t.Parallel()
    
    inventory.MustRun(t, fs, func(
        t *testing.T,
        test inventory.Test[
            *inventory.In[struct {
                A float64 `json:"a"`
                B float64 `json:"b"`
            }],
            *inventory.Out[float64, error],
        ],
    ) {
        result, err := divide.Divide(test.In.Arguments.A, test.In.Arguments.B)
        if err != nil {
            assert.EqualError(t, err, test.Out.Error.Error(), test.Title)
            
            return
        }
        
        assert.NoError(t, test.Out.Error, test.Title)
        assert.Equal(t, result, test.Out.Result, test.Title)
    })
}
```

Файл с тестовыми данными при использовании функции `Run` или `MustRun` должен находится внутри `embed.FS`, 
а название файла должно совпадать с именем теста (например, для текущего случая допустимое 
расположение: `test/TestDivide.json`). Пакет самостоятельно обнаружит файл внутри каталога и загрузит данные,
что позволяет разбивать каталоги на подкаталоги, чтобы обеспечить читаемость (например, `test/math/TestDivide.json`).
Тестовые данные заполняются по шаблону, на основании JSON тегов `Test` и переданных типов `In`, `Out`:

```json
[
  {
    "title": "ok",
    "in": {
      "arguments": {
        "a": 1,
        "b": 2
      }
    },
    "out": {
      "results": 0.5
    }
  },
  {
    "title": "not ok",
    "in": {
      "arguments": {
        "a": 2,
        "b": 0
      }
    },
    "out": {
      "error": "сan't divide by zero"
    }
  }
]
```

<!-- #### 4.2 Тест сущности -->
