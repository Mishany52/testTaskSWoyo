**Сборка и запуск api без использования базы**:

BACKEND_CMD="/backend" docker-compose -f docker-compose-dev.yml up --build -d

**Чтобы запустить api с подключением к базе**:

BACKEND_CMD="/backend -d" docker-compose -f docker-compose-dev.yml up --build -d

**Если собрали один из вариантов и хотите собрать другой, то нужно удалить контейнеры**:

docker-compose -f docker-compose-dev.yml down --volumes

**curl на создание и получение короткой ссылки**:
```
curl --location 'http://localhost:8080' \
 --header 'Content-Type: application/json' \
 --header 'Cookie: Cookie_1=value' \
 --data '{
"longUrl": "http://google.com/1111"
}'
```


**Полученная короткая ссылка при тестирование**
   http://localhost:8080/xQZVBf2

**Запрос на получение длинной по короткой ссылки, полученной из предыдущего запроса**:
``` 
curl --location 'http://localhost:8080' \
 --header 'Content-Type: application/json' \
 --header 'Cookie: Cookie_1=value' \
 --data '{
"longUrl": "http://google.com/1111"
}'
```

