# VK-internship
> Test assignment for an internship for the Go-developer (Summer 2023)

## Task description:
Создать простого чат бота на базе ВКонтакте АРI (допускается как альтернатива Telegram Bot АРI)

## Requirements:
Бот должен уметь отправлять приветственное сообщение, иметь минимум 4 кнопки на первом слое и минимум по 2 во втором.

Нужно реализовать на языке `Golang` с минимальным использованием готовых пакетов.
Код разместить на `github/bitbucket` и в поле ответа вставить ссылку на него с открытым доступом. По возможности также прислать ссылку на работающего бота (например через `Heroku`). 

## Features:
No third-party code was used
## Usage:
### N.B
TELEGRAM_TOKEN берется из OS.ENV
> Сборка:
> ```shell
> $ make build
> ```
> Сборка и запуск:
> ```shell
> $ make run
> ```
> Запуск в докере:
> ```shell
> $ make docker-run
> ```

