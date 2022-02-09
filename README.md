## Домашняя работа по курсу "Безопасность интернет-приложений"
Выполнил: Ершов Виталий

### Запуск
Запуск прокси-сервера
```
make run
```
Пример использования
```
curl -v -x http://127.0.0.1 http://mail.ru
```
Ответ
```
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 1080 (#0)
> GET http://mail.ru/ HTTP/1.1
> Host: mail.ru
> User-Agent: curl/7.64.1
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 301 Moved Permanently
< Connection: keep-alive
< Content-Length: 185
< Content-Type: text/html
< Date: Wed, 09 Feb 2022 09:22:39 GMT
< Location: https://mail.ru/
< Server: nginx/1.14.1
```
