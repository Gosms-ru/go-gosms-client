# GoSMS Client

[![Go Reference](https://pkg.go.dev/badge/github.com/gosms-ru/go-gosms-client.svg)](https://pkg.go.dev/github.com/gosms-ru/go-gosms-client)
[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/gosms-ru/go-gosms-client)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/version-v0.0.2-blue.svg)](https://github.com/gosms-ru/go-gosms-client/releases/tag/v0.0.2)

[![Русский](https://img.shields.io/badge/Русский-✓-blue)](#russian)
[![English](https://img.shields.io/badge/English-✓-green)](#english)

<a name="russian"></a>
# GoSMS Client

Go-клиент для работы с API GoSMS. Этот SDK позволяет отправлять SMS, получать информацию о них, удалять их, а также управлять устройствами через API GoSMS.

## Требования
- Go >= 1.21

## Содержание

- [Установка](#установка)
- [Настройка](#настройка)
- [Использование](#использование)
  - [Отправка SMS](#отправка-sms)
  - [Получение информации об SMS](#получение-информации-об-sms)
  - [Удаление SMS](#удаление-sms)
  - [Получение списка SMS](#получение-списка-sms)
  - [Работа с устройствами](#работа-с-устройствами)
    - [Получение информации об устройстве](#получение-информации-об-устройстве)
    - [Редактирование устройства](#редактирование-устройства)
    - [Удаление устройства](#удаление-устройства)
- [Обработка ошибок](#обработка-ошибок)
- [Тестирование](#тестирование)
- [Документация](#документация)
- [Лицензия](#лицензия)

## Установка

```bash
# Установка последней версии
go get github.com/gosms-ru/go-gosms-client

# Установка с обновлением всех зависимостей до последних версий
go get -u github.com/gosms-ru/go-gosms-client
```

## Настройка

Для использования SDK вам потребуется токен доступа к API GoSMS. Получить его можно в [панели управления GoSMS](https://cms.gosms.ru/).

```go
import "github.com/gosms-ru/go-gosms-client"

// Создаем клиент с вашим токеном
client := gosms.NewClient("ваш-токен")
```

## Использование

### Отправка SMS

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на отправку SMS
    req := gosms.SendSMSRequest{
        Message:     "Тестовое сообщение",
        PhoneNumber: "79999999999",
        DeviceID:    "device-id",     // опционально
        ToSim:       1,               // опционально
        CallbackID:  "callback-id",   // опционально
    }
    
    // Отправляем SMS
    response, err := client.SendSMS(req)
    if err != nil {
        log.Fatalf("Ошибка отправки SMS: %v", err)
    }
    
    fmt.Printf("SMS отправлено, ID: %s\n", response.ID)
}
```

### Получение информации об SMS

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на получение информации об SMS
    req := gosms.GetSMSRequest{
        ID: "6654a4e8f1527149588c89f2",
    }
    
    // Получаем информацию об SMS
    response, err := client.GetSMS(req)
    if err != nil {
        log.Fatalf("Ошибка получения информации об SMS: %v", err)
    }
    
    fmt.Printf("Информация об SMS:\n")
    fmt.Printf("ID: %s\n", response.ID)
    fmt.Printf("Сообщение: %s\n", response.Message)
    fmt.Printf("Статус: %d\n", response.Status)
    fmt.Printf("Статус сообщения: %s\n", response.MessageStatus)
    fmt.Printf("Номер телефона: %s\n", response.PhoneNumber)
    fmt.Printf("ID устройства: %s\n", response.DeviceID)
    fmt.Printf("ID callback: %s\n", response.CallbackID)
    fmt.Printf("Время создания: %d\n", response.TimeCreate)
    if response.ToSim != nil {
        fmt.Printf("Номер SIM: %d\n", *response.ToSim)
    }
}
```

### Удаление SMS

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на удаление SMS
    req := gosms.DeleteSMSRequest{
        ID: "6654a4e8f1527149588c89f2",
    }
    
    // Удаляем SMS
    err := client.DeleteSMS(req)
    if err != nil {
        log.Fatalf("Ошибка удаления SMS: %v", err)
    }
    
    fmt.Println("SMS успешно удалено")
}
```

### Получение списка SMS

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на получение списка SMS
    req := gosms.ListSMSRequest{
        Limit:  5,                // обязательное поле, от 1 до 100
        Offset: 1,                // опционально, по умолчанию 1
        Search: "79999999999",    // опционально, поиск по номеру телефона
    }
    
    // Получаем список SMS
    response, err := client.ListSMS(req)
    if err != nil {
        log.Fatalf("Ошибка получения списка SMS: %v", err)
    }
    
    fmt.Printf("Всего записей: %d\n", response.Pagination.TotalRecords)
    fmt.Printf("Текущая страница: %d\n", response.Pagination.Offset)
    fmt.Printf("Записей на странице: %d\n\n", response.Pagination.Limit)
    
    for _, sms := range response.SMSList {
        fmt.Printf("ID: %s\n", sms.ID)
        fmt.Printf("Сообщение: %s\n", sms.Message)
        fmt.Printf("Статус: %d\n", sms.Status)
        fmt.Printf("Статус сообщения: %s\n", sms.MessageStatus)
        fmt.Printf("Номер телефона: %s\n", sms.PhoneNumber)
        fmt.Printf("ID устройства: %s\n", sms.DeviceID)
        fmt.Printf("ID callback: %s\n", sms.CallbackID)
        fmt.Printf("Время создания: %d\n", sms.TimeCreate)
        if sms.ToSim != nil {
            fmt.Printf("Номер SIM: %d\n", *sms.ToSim)
        }
        fmt.Println("---")
    }
}
```

### Работа с устройствами

#### Получение информации об устройстве

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на получение информации об устройстве
    req := gosms.GetDeviceInfoRequest{
        DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
    }
    
    // Получаем информацию об устройстве
    response, err := client.GetDeviceInfo(req)
    if err != nil {
        log.Fatalf("Ошибка получения информации об устройстве: %v", err)
    }
    
    fmt.Printf("ID устройства: %s\n", response.DeviceID)
    fmt.Printf("Заряд батареи: %d%%\n", response.DeviceBatteryState)
    fmt.Printf("Название устройства: %s\n", response.DeviceName)
    fmt.Printf("Активно: %v\n", response.IsActive)
    fmt.Printf("Заряжается: %v\n", response.IsCharging)
    fmt.Printf("Последний онлайн: %s\n", response.LastOnlineDate)
    fmt.Printf("Тип устройства: %s\n", response.DeviceNameType)
    fmt.Printf("Оповещение о низком заряде: %v\n", response.LowBatteryAlert)
    fmt.Printf("SIM по умолчанию: %d\n", response.ToSim)
    
    fmt.Println("\nСписок SIM-карт:")
    for _, sim := range response.SimList {
        fmt.Printf("Слот %d: %s\n", sim.SlotIndex, sim.DisplayName)
    }
}
```

#### Редактирование устройства

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на редактирование устройства
    req := gosms.EditDeviceRequest{
        DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
        IsActive: true, // включить отправку SMS
    }
    
    // Редактируем устройство
    err := client.EditDevice(req)
    if err != nil {
        log.Fatalf("Ошибка редактирования устройства: %v", err)
    }
    
    fmt.Println("Устройство успешно отредактировано")
}
```

#### Удаление устройства

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Создаем клиент с вашим токеном
    client := gosms.NewClient("ваш-токен")
    
    // Создаем запрос на удаление устройства
    req := gosms.DeleteDeviceRequest{
        DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
    }
    
    // Удаляем устройство
    err := client.DeleteDevice(req)
    if err != nil {
        log.Fatalf("Ошибка удаления устройства: %v", err)
    }
    
    fmt.Println("Устройство успешно удалено")
}
```

## Обработка ошибок

SDK возвращает ошибки в формате `error`, которые можно обработать с помощью стандартных механизмов Go:

```go
response, err := client.SendSMS(req)
if err != nil {
    // Обработка ошибки
    log.Fatalf("Ошибка: %v", err)
}
```

## Тестирование

Для запуска тестов используйте команду:

```bash
go test -v
```

## Документация

Для получения токена и подробной информации об API посетите [документацию GoSMS](https://docs.gosms.ru/).

## Лицензия

[MIT](LICENSE)

---

<a name="english"></a>
# GoSMS Client

Go client for working with the GoSMS API. This SDK allows you to send SMS, retrieve information about them, delete them, and manage devices through the GoSMS API.

## Requirements

- Go >= 1.21

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Sending SMS](#sending-sms)
  - [Getting SMS Information](#getting-sms-information)
  - [Deleting SMS](#deleting-sms)
  - [Getting SMS List](#getting-sms-list)
  - [Working with Devices](#working-with-devices)
    - [Getting Device Information](#getting-device-information)
    - [Editing Device](#editing-device)
    - [Deleting Device](#deleting-device)
- [Error Handling](#error-handling)
- [Testing](#testing)
- [Documentation](#documentation)
- [License](#license)

## Installation

```bash
# Install the latest version
go get github.com/gosms-ru/go-gosms-client

# Install with updating all dependencies to their latest versions
go get -u github.com/gosms-ru/go-gosms-client
```

## Configuration

To use the SDK, you need an access token for the GoSMS API. You can get it in the [GoSMS control panel](https://cms.gosms.ru/).

```go
import "github.com/gosms-ru/go-gosms-client"

// Create a client with your token
client := gosms.NewClient("your-token")
```

## Usage

### Sending SMS

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to send SMS
    req := gosms.SendSMSRequest{
        Message:     "Test message",
        PhoneNumber: "79999999999",
        DeviceID:    "device-id",     // optional
        ToSim:       1,               // optional
        CallbackID:  "callback-id",   // optional
    }
    
    // Send SMS
    response, err := client.SendSMS(req)
    if err != nil {
        log.Fatalf("Error sending SMS: %v", err)
    }
    
    fmt.Printf("SMS sent, ID: %s\n", response.ID)
}
```

### Getting SMS Information

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to get SMS information
    req := gosms.GetSMSRequest{
        ID: "6654a4e8f1527149588c89f2",
    }
    
    // Get SMS information
    response, err := client.GetSMS(req)
    if err != nil {
        log.Fatalf("Error getting SMS information: %v", err)
    }
    
    fmt.Printf("SMS Information:\n")
    fmt.Printf("ID: %s\n", response.ID)
    fmt.Printf("Message: %s\n", response.Message)
    fmt.Printf("Status: %d\n", response.Status)
    fmt.Printf("Message Status: %s\n", response.MessageStatus)
    fmt.Printf("Phone Number: %s\n", response.PhoneNumber)
    fmt.Printf("Device ID: %s\n", response.DeviceID)
    fmt.Printf("Callback ID: %s\n", response.CallbackID)
    fmt.Printf("Creation Time: %d\n", response.TimeCreate)
    if response.ToSim != nil {
        fmt.Printf("SIM Number: %d\n", *response.ToSim)
    }
}
```

### Deleting SMS

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to delete SMS
    req := gosms.DeleteSMSRequest{
        ID: "6654a4e8f1527149588c89f2",
    }
    
    // Delete SMS
    err := client.DeleteSMS(req)
    if err != nil {
        log.Fatalf("Error deleting SMS: %v", err)
    }
    
    fmt.Println("SMS successfully deleted")
}
```

### Getting SMS List

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to get SMS list
    req := gosms.ListSMSRequest{
        Limit:  5,                // required field, from 1 to 100
        Offset: 1,                // optional, default is 1
        Search: "79999999999",    // optional, search by phone number
    }
    
    // Get SMS list
    response, err := client.ListSMS(req)
    if err != nil {
        log.Fatalf("Error getting SMS list: %v", err)
    }
    
    fmt.Printf("Total records: %d\n", response.Pagination.TotalRecords)
    fmt.Printf("Current page: %d\n", response.Pagination.Offset)
    fmt.Printf("Records per page: %d\n\n", response.Pagination.Limit)
    
    for _, sms := range response.SMSList {
        fmt.Printf("ID: %s\n", sms.ID)
        fmt.Printf("Message: %s\n", sms.Message)
        fmt.Printf("Status: %d\n", sms.Status)
        fmt.Printf("Message Status: %s\n", sms.MessageStatus)
        fmt.Printf("Phone Number: %s\n", sms.PhoneNumber)
        fmt.Printf("Device ID: %s\n", sms.DeviceID)
        fmt.Printf("Callback ID: %s\n", sms.CallbackID)
        fmt.Printf("Creation Time: %d\n", sms.TimeCreate)
        if sms.ToSim != nil {
            fmt.Printf("SIM Number: %d\n", *sms.ToSim)
        }
        fmt.Println("---")
    }
}
```

### Working with Devices

#### Getting Device Information

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to get device information
    req := gosms.GetDeviceInfoRequest{
        DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
    }
    
    // Get device information
    response, err := client.GetDeviceInfo(req)
    if err != nil {
        log.Fatalf("Error getting device information: %v", err)
    }
    
    fmt.Printf("Device ID: %s\n", response.DeviceID)
    fmt.Printf("Battery Level: %d%%\n", response.DeviceBatteryState)
    fmt.Printf("Device Name: %s\n", response.DeviceName)
    fmt.Printf("Is Active: %v\n", response.IsActive)
    fmt.Printf("Is Charging: %v\n", response.IsCharging)
    fmt.Printf("Last Online: %s\n", response.LastOnlineDate)
    fmt.Printf("Device Type: %s\n", response.DeviceNameType)
    fmt.Printf("Low Battery Alert: %v\n", response.LowBatteryAlert)
    fmt.Printf("Default SIM: %d\n", response.ToSim)
    
    fmt.Println("\nSIM Cards List:")
    for _, sim := range response.SimList {
        fmt.Printf("Slot %d: %s\n", sim.SlotIndex, sim.DisplayName)
    }
}
```

#### Editing Device

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to edit device
    req := gosms.EditDeviceRequest{
        DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
        IsActive: true, // enable SMS sending
    }
    
    // Edit device
    err := client.EditDevice(req)
    if err != nil {
        log.Fatalf("Error editing device: %v", err)
    }
    
    fmt.Println("Device successfully edited")
}
```

#### Deleting Device

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/gosms-ru/go-gosms-client"
)

func main() {
    // Create a client with your token
    client := gosms.NewClient("your-token")
    
    // Create a request to delete device
    req := gosms.DeleteDeviceRequest{
        DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
    }
    
    // Delete device
    err := client.DeleteDevice(req)
    if err != nil {
        log.Fatalf("Error deleting device: %v", err)
    }
    
    fmt.Println("Device successfully deleted")
}
```

## Error Handling

The SDK returns errors in the `error` format, which can be handled using standard Go mechanisms:

```go
response, err := client.SendSMS(req)
if err != nil {
    // Error handling
    log.Fatalf("Error: %v", err)
}
```

## Testing

To run tests, use the command:

```bash
go test -v
```

## Documentation

To get a token and detailed information about the API, visit the [GoSMS documentation](https://docs.gosms.ru/).

## License

[MIT](LICENSE)