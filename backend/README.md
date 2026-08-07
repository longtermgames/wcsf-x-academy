# WCSF X Academy — backend регистраций

Маленький Go-сервис: принимает заявки на соревнования (BMX / воркаут-баттл / батут) и хранит их в Postgres. Список зарегистрированных смотрится на `/admin`.

## Локальный запуск

Нужен Postgres (можно локальный, можно сразу Neon).

```bash
cp .env.example .env
# впиши свой DATABASE_URL, ADMIN_USER, ADMIN_PASS
export $(cat .env | xargs)
go run .
```

Таблица `registrations` создаётся автоматически при старте (см. [migrate.go](migrate.go)).

## API

### `POST /api/register`

```json
{
  "full_name": "Иван Иванов",
  "phone": "+996700000000",
  "discipline": "bmx"
}
```

`discipline` — одно из `bmx`, `workout`, `trampoline`. Ответ `201` с созданной записью, `400` при невалидных данных.

### `GET /admin`

HTML-страница со списком заявок и счётчиками по дисциплинам. Защищена Basic Auth (`ADMIN_USER` / `ADMIN_PASS`).

### `GET /admin/export.csv`

То же самое, но CSV для скачивания.

## Деплой

**База — Neon** (neon.tech): создать проект, скопировать connection string (с `?sslmode=require`) в `DATABASE_URL`. Бессрочный free tier, без ограничения по времени как у Railway trial.

**Сервис — Render** (render.com):
1. New → Web Service → подключить репозиторий, Root Directory `backend`.
2. Runtime — Docker (подхватит [Dockerfile](Dockerfile) автоматически).
3. Environment → добавить `DATABASE_URL`, `ADMIN_USER`, `ADMIN_PASS`, `ALLOWED_ORIGIN=https://wcsfxacademy.com`.
4. Free-инстанс засыпает после ~15 минут простоя и первый запрос после сна будет с задержкой ~30-50 сек — это нормально для free tier.

## Дальше

Фронтенд ещё не подключён — на сайте пока нет формы, которая стучится в `/api/register`. Это отдельный шаг: добавить форму на `index.html` в блоки BMX/воркаут-баттл/батут с `fetch(...)` на этот эндпоинт.
