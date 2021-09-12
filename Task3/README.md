# Домашнее задание номер 3
____

Здравствуйте, уважаемые преподаватели!
С вашего позволения представляю выполненный проект третьего домашнего задания.

Моя программа реализует задачу по поиску информации о хозяине ip-адресса(ов) ((зависит от количества выбранных строк)), к которому(ым) в текущий момент работы подключена подключена программа.
____
Hello, dear teachers
With your permission, i'd like to introduce my fully-completed Task3 project

My program was build to serve as wrapper over the standard, not so comfortable, terminal. You may decide to use one of two built-in modules: netstat or ss - both of them are fine. Overall purpose of my program is to collect 'whois' info of connections, established by single process.

## Инструкция

Для того, чтобы получить желаемый результат пользователю требуется перейти в корень папки "Task 3" путём использования терминала, IDE или графического интерфейса. Следующим шагом должен стать запуск программы, для этого с помощью консоли или среды разработки запустите файл whois.py. Программа не требует особых прав для запуска, поэтому вы вполне можете использовать команду вида "python3 whois.py"
____
First of all, user should move himself directly into root project folder by using terminal/ide or sysgui. Second step is all about execution 'whois.py' file. Program does not require sudo permissions, therefore it will be absolutely fine with classic "python3 whois.py" command


После запуска программа спросит вас о желаемом модуле получения списка активных соединений:
____
When lauched, program asks your decision about module to work with

![Alt-текст](https://i.imgur.com/j8quS2w.png "скриншот")

**Не переживайте, если модуль netstat у вас не установлен - скрипт уведомит вас об этом**
____
**If by some reasons, you dont have net-tools installed, script will notify about it**

На следующем шаге требуется ввести pid или имя процесса, статистика по которому вас интересует. 
Для полученного отсортированного списка адресов, предлагается выбрать количество строк для сбора информации, а также желаемый параметр (либо указать все параметры в качестве требуемых)
____
Right after action performed, it requires you to type pid or process name, which you would like to collect information about. You will also be asked about iplist length, once it is ready

Конечный результат зависит от выбранных параметров:
____
Result depends on your entries:

**Таким образом для 'all' вывод будет выглядеть так:**
![Alt-текст](https://i.imgur.com/v0kgXuj.png "скриншот")
____
**Case for 'all' parameters chosen**

**А для 'organization' вот так:**
![Alt-текст](https://i.imgur.com/w9dY9Jp.png "скриншот")
____
**And a special one for 'organization'**


Благодарю за внимание!
____
Thank you for your attention!