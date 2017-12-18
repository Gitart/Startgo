# Start Go
### Simple service for Go lang and RethinkDB  

#### Инструкция по созданию рабочей среды для Golang + RethinkDB  
#### Автор : Савченко Артур

## Go
1. Скачать и установить язык разработки [Go](https://golang.org/doc/install?download=go1.9.2.windows-amd64.msi)
2. Скачать и установить для среду разработки например [VSCode](https://code.visualstudio.com/docs/?dv=win)  
3. Можно воспользоваться редактором [Sublime](https://www.sublimetext.com/3) или [Atom](https://atom.io/) 
4. Создать директорию для рабочего проекта (например: md c:/workgo)
5. Создать bat файл для компиляции проекта run.bat

## RethinkDB
1. Скачать и установить базу данных [Rethinkdb](https://download.rethinkdb.com/windows/rethinkdb-2.3.6.zip) для Windows	 
2. Запустить базу данных :  **rethinkdb.exe -n jarvis**
3. Панель управления (by default): **localhost:8080**
4. Адрес для подключения:  **localhost:28015**
6. Драйвер для базы [RethinkDb](https://github.com/GoRethink/gorethink) (инсталируется go get -u github.com/dancannon/gorethink) :
7. [Команды](https://rethinkdb.com/api/javascript/) управления базой

## Service
1. Подключение к базе данных
2. Опредление роутинга
3. Описание сервиса

## Testing 
1. Curl  
2. Insomnia
3. Postmen   
4. etc.
