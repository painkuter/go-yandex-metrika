go-yandex-metrika
=================

Библиотека для Go, работающая с API Yandex Метрики.

Для работы необходим логин с паролем, код или токен (oauth_token).

Необходимые для работы параметры можно получить зарегистрировав приложение вот тут https://oauth.yandex.ru/client/new.
Как получить отладочный токен рассказано тут https://tech.yandex.ru/oauth/doc/dg/tasks/get-oauth-token-docpage/

Библиотека работает только с JSON форматом.

Пример использования
--------------------

```go
package main

import (
	"log"

	metrika "github.com/xboston/go-yandex-metrika"
)

func main() {

	log.Println("Start")

	token := ""

	metrika := metrika.NewMetrikaFromToken(token)
	metrika.SetDebug(true)

	metrika.Authorize()
	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
```


Основано на проекте yametrikapy: https://github.com/pikhovkin/yametrikapy/blob/master/yametrikapy/core.py
