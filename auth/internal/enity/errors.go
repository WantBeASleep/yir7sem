package enity

import "errors"

/*
- codes.OK: Операция прошла успешно.
- codes.Canceled: Операция была отменена.
- codes.Unknown: Неизвестная ошибка.
- codes.InvalidArgument: Передан неверный аргумент.
- codes.DeadlineExceeded: Истекло время выполнения операции.
- codes.NotFound: Ресурс не найден.
- codes.AlreadyExists: Ресурс уже существует.
- codes.PermissionDenied: Доступ к ресурсу запрещен.
- codes.ResourceExhausted: Ресурсы исчерпаны.
- codes.FailedPrecondition: Операция не выполнена из-за несоблюдения предварительного условия.
- codes.Aborted: Операция была прервана.
- codes.OutOfRange: Аргумент вышел за допустимые пределы.
- codes.Unimplemented: Операция не поддерживается или отсутствует.
- codes.Internal: Внутренняя ошибка.
- codes.Unavailable: Служба в настоящее время недоступна.
- codes.DataLoss: Безвозвратная потеря данных.
- codes.Unauthenticated: Клиент не авторизован.
*/

var (
	ErrNotFound = errors.New("not found")
)
