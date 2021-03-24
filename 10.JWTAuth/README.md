## Лекция 10. Простейши механизм аутентификации

На данный момент у нас реализован API , с одной проблмой - кто угодно может получить доступ к элемента в БД через публичные запросы, и например, удалить все что там имеется.

***Идея*** : сделать так, чтобы пользователь, который собирается использовать наш API не был анонимным, а мог зарегестрироваться и пройти базовую аутентификацию.

### Шаг 0. Термины
***Аутентификация*** - процесс узнавания свой/чужой. (Подразуемвает под собой сопоставление данных стороннего пользователя с данными, которые уже имеются в бд.) 
***Авторизация*** - процесс выдачи прав доступа различного уровня.


### Шаг 1. Простейшая логика при аутентификации
* К нам пришел какой-то пользователь
* Пользователь должен пройти регистрацию 
* Пользователь переходит на ресурс аутентификации и получает какой-либо аутентификационный ключ
* Далее пользователь с этим ключом может ходить по всем ресурсам нашего api.

### Шаг 2. Аутентификация с помощью JWT токена
***JWT** - ```JsonWebToken``` - символьная строка с закодированным ключом.

### Шаг 3. Немного про то, где будут выполняться действия по работе с JWT
***Middleware*** - часть ПО (архитектурная часть), которая напрямую не взаимодействует ни с клиентом, ни с сервером, а осуществляет какие-либо команды или запросы во-время клиент-серврного общения.
Например:
* Пользователь вызывает ```POST /api/v1/article +.json``` 
* Auth Middleware - проверяет, может ли данный клиент данный запрос вообще выполнять или у него не ххватает прав? (Мы не знаем кто это)
* Сервер должен принять данные и обработать запрос (добавить в бд инфу про статью)

### Шаг 4. Реализация
Добавим 2 зависмости в проект:
* ```go get -u github.com/auth0/go-jwt-middleware```
* ```go get -u github.com/form3tech-oss/jwt-go```

В следующей директории :```internal/app/middleware/middleware.go```

```
package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey      []byte      = []byte("UltraRestApiSectryKey9000")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: emptyValidFunc,
	SigningMethod:       jwt.SigningMethodHS256,
})

```

### Шаг 5. Как пользователю получить этот токен?
Нам необходимо реализовать ресурс ```/auth``` или ```api/v1/user/auth```.
```
//func for configure Router
func (s *APIServer) configureRouter() {
	s.router.HandleFunc(prefix+"/articles", s.GetAllArticles).Methods("GET")
	s.router.HandleFunc(prefix+"/articles"+"/{id}", s.GetArticleById).Methods("GET")
	s.router.HandleFunc(prefix+"/articles"+"/{id}", s.DeleteArticleById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/articles", s.PostArticle).Methods("POST")
	s.router.HandleFunc(prefix+"/user/register", s.PostUserRegister).Methods("POST")
	//new pair for auth
	s.router.HandleFunc(prefix+"/user/auth", s.PostToAuth).Methods("POST")
}

```

### шАГ 6. Реализация PostToAuth
```
func (api *APIServer) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth POST /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	//Обрабатываем случай, если json - вовсе не json или в нем какие-либо пробелмы
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Необходимо попытаться обнаружить пользователя с таким login в бд
	userInDB, ok, err := api.store.User().FindByLogin(userFromJson.Login)
	// Проблема доступа к бд
	if err != nil {
		api.logger.Info("Can not make user search in database:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Если подключение удалось , но пользователя с таким логином нет
	if !ok {
		api.logger.Info("User with that login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Если пользователь с таким логином ест ьв бд - проверим, что у него пароль совпадает с фактическим
	if userInDB.Password != userFromJson.Password {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Теперь выбиваем токен как знак успешной аутентифкации
	token := jwt.New(jwt.SigningMethodHS256)             // Тот же метод подписания токена, что и в JwtMiddleware.go
	claims := token.Claims.(jwt.MapClaims)               // Дополнительные действия (в формате мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Время жизни токена
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	//В случае, если токен выбить не удалось!
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//В случае, если токен успешно выбит - отдаем его клиенту
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
```

### Шаг 7. Проверим, что токен выбивается
Для этого идем в postman

### Шаг 8. Завернем необходимые хендлеры в JWT-REQUIRED-декоратор
Для того, чтобы обозначит факт необходимости использования JWT токена перед выполнением какого-либо запроса - заверните его в декоратор ```middleware.JwtMiddleware```
```
//Теперь требует наличия JWT
	s.router.Handle(prefix+"/articles"+"/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.GetArticleById),
	)).Methods("GET")
	//
```

### Шаг 9. В postman
На вкладке ```Headers``` у данного запроса доавбляем пару параметров
```Authorization``` и ```Bearer <your_token_form_auth>```