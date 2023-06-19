local http = require("socket.http")
local url = require("net.url")

print("Hello world")
print(os.getenv("HOME"))
print(os.getenv("HOME2") or "Hey")

REDIS_HOST = os.getenv("REDIS_HOST") or "localhost"
REDIS_PORT = tonumber(os.getenv("REDIS_PORT") or 6379)
API_HOST = os.getenv("API_HOST") or "http://localhost:8080" .. "api" .. "v1"
-- print(API_HOST)
print(API_HOST)
-- print(url.concat("a", "b"))
local u = url.parse("http://www.example.com/test/?start=10")
print(u)

-- http

-- print('a' / 'b')
-- local body, code, headers, status = http.request("http://google.com")
-- if body then
--     -- запрос выполнился успешно
--     print(body) -- в body тело ответа сервера
-- else
--     -- произошла ошибка
--     print(code) -- сообщение об ошибке (например, "сервер на найден")
-- end