package math

import "testing"

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg is negative-expected error not nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg is negative-expected error not nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg are negative-expected error not nil")
	}
}

func TestAddAllZeroes(t *testing.T) {
	_, err := Add(0, 1)
	if err == nil {
		t.Error("first arg is zero-expected error not nil")
	}
	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg is zero-expected error not nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("both args is zero-expected error not nil")
	}
}
/* go fmt*/
/*Объект *testing.T предоставляет доступ к нескольким базовым методам:
Error, Errorf — записывает сообщение в error-лог и помечает тест как непройденный. Исполнение теста продолжается.
Fatal, Fatalf — делает то же самое, но исполнение теста немедленно завершается. Этот метод часто используется в рабочих проектах при обработке ошибок. Очень удобен при отладке, когда тестируется какой-то конкретный участок кода.
Skip, Skipf — позволяет пропустить тест с сообщением. Используется, когда окружение для теста не задано. Типичный сценарий — прогон интеграционных тестов с внешним сервисом только на CI, где к нему есть доступы.
Log, Logf — позволяет выводить лог-сообщения внутри теста. Преимущество перед методами пакета fmt в том, что из лога сразу видно, к какому тесту относится сообщение.
Run(name string, testf func(t *testing.T) ) — запускает функцию в качестве теста, что удобно при выполнении нескольких запусков теста, например, с разными именами.
*/

/*go test-запуск теста
  go test add -cover  - степень покрытия тестами
  go test . -coverprofile=coverage.out - Запустить на нём утилиту go test и сохранить файл профиля покрытия тестами. Путь к файлу с профилем — это значение флага -coverprofile. В данном случае сохраним его в файл coverage.out текущей директории
  go tool cover -html=coverage.out  - Проанализировать полученный файл утилитой cover. Например, по собранному профилю можно получить HTML-представление исходного кода с дополнительной разметкой, связанной с покрытием тестами
*/
/* go fmt*/