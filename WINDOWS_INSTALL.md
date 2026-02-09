# Установка CHLR на Windows

## Способ 1: Скачать готовый бинарник (рекомендуется)

### Шаг 1: Скачайте бинарник

1. Откройте браузер и перейдите по ссылке:
   ```
   https://github.com/Chelaran/CHLR/releases/latest
   ```

2. Найдите файл `chlr-windows-amd64.exe` в списке Assets

3. Нажмите на файл для скачивания

### Шаг 2: Установка

**Вариант A: Использовать из любой папки**

1. Создайте папку для утилит (например, `C:\tools`)
2. Переместите скачанный `chlr-windows-amd64.exe` в эту папку
3. Переименуйте файл в `chlr.exe`
4. Добавьте папку в PATH:
   - Нажмите `Win + R`, введите `sysdm.cpl` и нажмите Enter
   - Перейдите на вкладку "Дополнительно"
   - Нажмите "Переменные среды"
   - В разделе "Системные переменные" найдите `Path` и нажмите "Изменить"
   - Нажмите "Создать" и добавьте путь к папке (например, `C:\tools`)
   - Нажмите OK во всех окнах
   - Перезапустите командную строку/PowerShell

**Вариант B: Использовать из текущей папки**

1. Переместите `chlr-windows-amd64.exe` в папку, где вы работаете
2. Переименуйте в `chlr.exe`
3. Используйте как `.\chlr.exe` или полный путь

### Шаг 3: Проверка

Откройте PowerShell или CMD и выполните:
```powershell
chlr --version
```

Должно вывести: `chlr version 0.1.2`

## Способ 2: Через PowerShell (автоматическая установка)

Откройте PowerShell и выполните:

```powershell
# Скачать последнюю версию
$url = "https://github.com/Chelaran/CHLR/releases/latest/download/chlr-windows-amd64.exe"
$output = "$env:USERPROFILE\chlr.exe"
Invoke-WebRequest -Uri $url -OutFile $output

# Добавить в PATH (для текущей сессии)
$env:Path += ";$env:USERPROFILE"

# Проверить
.\chlr.exe --version
```

Для постоянного добавления в PATH:
```powershell
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$env:USERPROFILE", "User")
```

## Использование

После установки вы можете использовать утилиту:

```powershell
# Создать простой проект
chlr init my-project

# Создать проект с PostgreSQL
chlr init my-app --db=postgres

# Создать монорепозиторий
chlr init my-platform --mono --db=postgres
```

## Помощь

Если что-то не работает:
- Проверьте, что файл скачан полностью
- Убедитесь, что антивирус не блокирует файл
- Попробуйте запустить от имени администратора
- Проверьте версию: `chlr --version`

