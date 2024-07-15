## Название проекта
Сервис сбора статистики

## Описание проекта
Микросервис на языке Go для сбора статистики.

## Цель проекта
Создать сервис с возможностью хранения статистики в базе данных.

## Ручки
1. GetOrderBook(exchangename, pair string)
    - URL: `/api/get_order_book`
    - Метод: `GET`
    - Описание: Получить стакан заказов для указанной биржи и торговой пары
    - Возвращаемое значение: Массив DepthOrder и error
   
   Формат json запроса:
      ```
      type GetOrderBookRequest struct {
          ExchangeName string `json:"exchange_name"`
          Pair         string `json:"pair"`
      }
      

2. SaveOrderBook(exchangename, pair string, orderBook DepthOrder)
    - URL: `/api/save_order_book`
    - Метод: `POST`
    - Описание: Сохранить стакан заказов для указанной биржи и торговой пары
    - Возвращаемое значение: error

   Формат json запроса:
      ```
      type SaveOrderBookRequest struct {
          ExchangeName string       `json:"exchange_name"`
	         Pair         string       `json:"pair"`
          OrderBook    []DepthOrder `json:"orderBook"`
      }

3. GetOrderHistory(client Client)
    - URL: `/api/get_order_history`
    - Метод: `GET`
    - Описание: Получить историю заказов для клиента
    - Возвращаемое значение: массив HistoryOrder и error

   Формат json запроса:
      ```
      type GetOrderHistoryRequest struct {
          Client Client `json:"client"`
      }
   

4. SaveOrder(client Client, order HistoryOrder)
    - URL: `/api/save_order`
    - Метод: `POST`
    - Описание: Сохранить заказ для указанного клиента
    - Возвращаемое значение: error

   Формат json запроса:
      ```
      type SaveOrderRequest struct {
	      Client Client       `json:"client"`
	      Order  HistoryOrder `json:"order"`
      }
   
## Используемые библиотеки
1. Для PostgreSQL
   `"github.com/jmoiron/sqlx"`
2. Для HTTP-запросов
   `"github.com/gin-gonic/gin"`
3. Для поддержки файлов конфигурации
   `"github.com/joho/godotenv"` и `"github.com/spf13/viper"`

## Прочее
Сервер работает на локальном хосте. Конфигурацию запуска можно изменить в файле `config.yml`

Пароль для postgres: qwerty
