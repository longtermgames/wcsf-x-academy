# Как обновлять сайт без программиста

Всё делается прямо в браузере на github.com, в репозитории [longtermgames/wcsf-x-academy](https://github.com/longtermgames/wcsf-x-academy).

---

## Добавить товар в мерч

### 1. Загрузите фото товара

1. Откройте папку [assets/images](https://github.com/longtermgames/wcsf-x-academy/tree/main/assets/images)
2. **Add file → Upload files**
3. Перетащите фото (jpg или png), любое имя файла — например `sleeves.jpg`
4. **Commit changes**

### 2. Добавьте товар в список

1. Откройте [assets/products.json](https://github.com/longtermgames/wcsf-x-academy/blob/main/assets/products.json)
2. Нажмите карандаш (**Edit**) в правом верхнем углу
3. Возьмите последний товар в списке как образец:

```json
  {
    "name": "Wrist Wraps",
    "desc": "Кистевые бинты для турника и брусьев, длина 40 см",
    "price": "450",
    "image": "wristwraps.svg",
    "whatsapp": "996700000000",
    "message": "Хочу заказать Wrist Wraps"
  }
```

4. После закрывающей `}` последнего товара добавьте запятую `,` и новый блок:

```json
  ,
  {
    "name": "Название товара",
    "desc": "Короткое описание",
    "price": "1000",
    "image": "sleeves.jpg",
    "whatsapp": "996700000000",
    "message": "Хочу заказать Название товара"
  }
```

   `"image"` должно точно совпадать с именем файла из шага 1 (включая `.jpg`/`.png`).

5. **Commit changes**. Через 1–2 минуты товар появится на странице /merch.html

### Изменить или убрать товар

Так же через карандаш в `assets/products.json`:
- **Изменить** — поправьте текст/цену внутри блока `{ ... }`
- **Убрать** — удалите весь блок `{ ... }` вместе с одной из соседних запятых

---

## Добавить документ

### 1. Загрузите файл

1. Откройте папку [assets/docs](https://github.com/longtermgames/wcsf-x-academy/tree/main/assets/docs)
2. **Add file → Upload files**
3. Загрузите PDF, например `smeta-2026.pdf`
4. **Commit changes**

### 2. Добавьте запись в список

1. Откройте [assets/documents.json](https://github.com/longtermgames/wcsf-x-academy/blob/main/assets/documents.json)
2. Карандаш → **Edit**
3. Файл разбит на две группы — `"Для участников"` и `"Мероприятия и официальные партнёры"`. Внутри нужной группы, в массиве `"items"`, добавьте после последнего блока запятую и новый:

```json
  ,
  {
    "title": "Название документа",
    "desc": "Короткое описание",
    "file": "smeta-2026.pdf",
    "updated": "07.2026",
    "size": "150 КБ"
  }
```

   `"file"` должно точно совпадать с именем файла из шага 1.

4. **Commit changes**. Документ появится на странице /documents.html

### Изменить, убрать или добавить новую группу

- **Изменить/убрать** — так же редактируете нужный блок `{ ... }` внутри `"items"`
- **Новая группа** — скопируйте целиком блок `{ "group": "...", "items": [...] }` и вставьте перед ним запятую

---

## Включить оплату по QR (MBank)

1. В приложении MBank откройте раздел приёма платежей и сохраните свой QR-код как картинку
2. Загрузите её в [assets/images](https://github.com/longtermgames/wcsf-x-academy/tree/main/assets/images) (**Add file → Upload files**), например `mbank-qr.png`
3. Откройте [assets/payment.json](https://github.com/longtermgames/wcsf-x-academy/blob/main/assets/payment.json), карандаш → **Edit**, и впишите:

```json
{
  "enabled": true,
  "provider": "MBank",
  "qrImage": "mbank-qr.png",
  "phone": "0700 00 00 00",
  "note": "Отсканируйте QR в приложении MBank или переведите на номер, затем пришлите скрин оплаты в WhatsApp"
}
```

   `"qrImage"` должно точно совпадать с именем файла из шага 2. Пока `"enabled": false` или `"qrImage"` пустой — блок оплаты на сайте не показывается.

4. **Commit changes**. Блок с QR появится на странице /merch.html

---

## Если что-то сломалось

Если после Commit changes сайт показывает "Не получилось загрузить список…" — где-то пропущена запятая или скобка в JSON. Откатить: откройте [историю коммитов](https://github.com/longtermgames/wcsf-x-academy/commits/main) нужного файла, откройте предыдущую версию, скопируйте её содержимое обратно. Либо просто напишите об этом — поправим.
