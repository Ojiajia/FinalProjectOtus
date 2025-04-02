HTTP-сервер для выдачи статических файлов
Описание:
В этом проекте студентам предлагается написать HTTP-сервер, который обрабатывает запросы и выдает статические файлы (HTML, CSS, JS) на клиент. Сервер должен обрабатывать файлы из заданной директории и возвращать их на клиент в зависимости от запрошенного пути.

Требования:
Программа должна принимать настройки для указания директории с статическими файлами и порта, на котором запускать сервер.
HTTP-сервер должен правильно обрабатывать запросы на различные типы файлов (например, HTML, CSS, JS, изображения).
Реализована обработка ошибок, связанных с отсутствием запрашиваемых файлов или неверными запросами.
Сервер должен быть способен обслуживать несколько клиентов одновременно.
Развертывание
Развертывание сервиса должно осуществляться с использованием docker compose в директории с проектом.

Тестирование
Написаны юнит-тесты на core логику приложения. Плюсом будут тесты на транспортном уровне и на уровне хранения.

Критерии оценивания
Максимум - 15 баллов (при условии выполнения обязательных требований):

Реализован алгоритм - 2 балла.
Реализовано разделение на слои (транспортный, хранения и т.д.) - 2 балла.
Реализовано API сервиса - 2 балла.
Реализован консольный пользовательский интерфейс - 2 балла.
Написаны юнит-тесты - 1 балл.
Написаны интеграционные тесты - 2 балла.
Тесты адекватны и полностью покрывают функциональность - 1 балл.
Понятность и чистота кода - до 3 баллов.
Зачёт от 10 баллов
