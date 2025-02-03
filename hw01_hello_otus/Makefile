# Цель по умолчанию (выполняется при запуске `make` без аргументов)
.DEFAULT_GOAL := all

# Запуск всех проверок
all: lint test

# Запуск линтера
lint:
	golangci-lint run .

# Запуск тестов
test:
	go test -v -count=1 -race -timeout=1m .

# Запуск скрипта test.sh (если он существует)
test-script:
	if [ -f ./test.sh ]; then ./test.sh; fi

# Запуск всех команд: линтер, тесты и скрипт
check: lint test test-script

# Очистка (если нужно)
clean:
	go clean