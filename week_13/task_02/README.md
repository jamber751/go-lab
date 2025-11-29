# Week 13, Task 02: REST API

## Описание
CRUD API для управления задачами.

## Эндпоинты
- `GET /tasks` - список задач
- `POST /tasks` - создать задачу
- `DELETE /tasks/{id}` - удалить задачу

## Запуск
```bash
go run main.go
```

## Тестирование
```bash
curl http://localhost:8080/tasks
curl -X POST -d '{"title":"Новая задача"}' http://localhost:8080/tasks
```

