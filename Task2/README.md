# Домашнее задание номер 2
____

Приветствую вас, уважаемые преподаватели.

**В данном задании мной были выполнены:**
    **- Web приложение Flask, работающее на портах 80 / 443.**
    **- Плэйбук, загружающий конфигурацию и само приложение на вирутальную машину**

Hello, dear teachers. I'd like to itroduce my homework task 2 project 

**While creating that task, those functions were accomplished:**
    **- Web Flask app, which listens both 80 / 443 ports.**
    **- Playbook, developed to upload required configuration and app itself into vm env**

____

## Часть 1. Flask
## Part 1. Flask

### Представление
### Introduction

Выполненное мной приложение принимает запросы сразу по двум портам : 80/443. Требуемый для https соединения сертификат создается в процессе загрузки ( об этом ниже ), но может использоваться и установленный мной по умолчанию (/certificate): в случае, если вы захотите запустить приложение на своей локальной машине.

The app, i've created listening both 80 and 443 ports to accept requests. SSL certificate will be created during playbook execution, but, it could also use default one: if you wouild like to launch this app on your localhost

### Функционал 

Была реализована обработка post запросов по шаблону *имя машины.localhost*, с использованием emoji: если требуемое животное может быть представлено в виде emoji, 
вы увидите его изображение. **но только в том случае, если ваш терминал их поддерживает**. 

The function of accepting post requests was implemented. Default post req addr template looks like : *vmname.localhost*. If your requested animal has it's presentation as emoji, you will see related one. **only if your terminal is familliar with emoji**

![Alt-текст](https://imgur.com/a/ymQv1a7 "скриншот")

Get секция. Что я ожидаю увидеть посещая случайную страницу?

Get section. What i expect when visiting any randrom page in the internet

![Alt-текст](https://imgur.com/a/vZXL5Eh "скриншот")


## Часть 2. Ansible 
## Part 2. Ansible

### Инициализация
### Initialization

Что требуется от вас для запуска плэйбука?
    - Отредактируйте файл inventory: 
        - Установите адрес вашей виртуальной машины в группе homegroup
        - Установите имя пользователя в качестве значения переменной ansible_user
    - Отредактируйте файл vault_76.yml
        - Расшифруйте указанный файл с помощью команды ansible-vault decrypt vault_76.yml --vault-password-file .vault_password
        - Установите значения sudo паролей, необходимых для отработки playbook
        - В целях безопасности зашифруйте указанный файл, используя комманду ansible-vault encrypt  vault_76.yml --vault-password-file .vault_password

Playbook execution essencials:
    - Edit inventory file:
        - Set your vm addr into placement inside homegroup section
        - Set your vm user into ansible_user variable
    - Edit vault_76.yml:
        - Decrypt it, by using ansible-vault decrypt vault_76.yml --vault-password-file .vault_password into terminal
        - Set your sudo passwords, required to accomplish all playbook commands
        - Encrypt this file back for security purposes, with command ansible-vault encrypt  vault_76.yml --vault-password-file .vault_password

### Персонализация
### Personalization

Вы можете изменить базовые параметры, используемые для создания сертификата ssl. Изменения осуществляется путем задания новых значений ssl содержащих переменных, непосредственно в delpoy_playbook.yml

You are able to change ssl parameters. To do this you should change ssl contained variables inside deploy_playbook.yml itself


### Запуск
После того как все необходимые действия выполнены, осуществите запуск плейбука командой ansible-playbook deploy_playbook.yml --vault-password-file .vault_password

When all of preparations done, execute your playbook with ansible-playbook deploy_playbook.yml --vault-password-file .vault_password command



Спасибо за уделенное время!
____
Thanks for the attention!
