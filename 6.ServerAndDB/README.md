## Лекция 6. Подключение к БД и стандартные схемы миграции


### Шаг 0. Общие соображения
* Определить модель данных (определить правило стыковки объекта в таблице вашей СУБД с объектом внутри языка) - как объект представлен в БД
* Обработчики модели (модельные методы) - как объект взаимодействует с БД
* Выделение публичных обработчиков и стыковка их с серверными запросами

***Миграция*** - это процесс изменения схемы хранения данных. (положительный - upgrade / up миграция) (отрицательная downgrade /down миграция)

***Data repository*** (репозиторий с обработчиками) - это и есть то место, где будут жить публичные обработчики (модельные методы).

### Шаг 1. Библиотеки для работы с бд
```database/sql```
```sqlx```
```gosql```

### Шаг 2. Инициализация хранилища
```storage/storage.go```
Цель данного моделя определить:
* инстанс хранилища
* конструктор хранилища
* публичный метод Open (установка соединения)
* публичный метод Close (закрытие соединения)


### Шаг 3. Инициализация Storage
```storage.go```
Главная проблема кроется внутри метода Open, т.к. по факту низкоуровненвый sql.Open "ленивый" (устанавливает соединение с бд только в момент осуществления первого запрос)
```config.go```
Содержит инстанс конфига и конструктор. Атрибутом конфига является лишь строка подключения вида :
```
"host=localhost port=5432 user=postgres password=postgres dbname=restapi sslmode=disable"
```

### Шаг 4. Добавление хранилища к API
Добавим новый атрибут к API
```
//Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавление поля для работы с хранилищем
	storage *storage.Storage
}
```

Добавим новый конфиуратор:
```
//Пытаемся отконфигурировать наше хранилище (storage API)
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}

```

### Шаг 5. Первичная миграция
Для начала установим ```scoop```
* Открываем PowerShell: ```Set-ExecutionPolicy RemoteSigned -scope CurrentUser``` и ```Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')```

* Для линукса/мака : https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

После установки ```scoop``` выполним: ```scoop install migrate```

### 5.1 Создание миграционного репозитория
В данном репозитории будут находится up/down пары sql миграционных запросов к бд.
```
migrate create -ext sql -dir migrations UsersCreationMigration
```

### 5.2 Создание up/down sql файлов
См. ```migrations/....up.sql``` и ```migrations/...down.sql```

### Шаг 5.3 Применить миграцию
```
migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up
```
