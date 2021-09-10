from network_modules.NetworkModule import NetworkModule
from handlers.KIHander import KeyboardIterruptedHandler
from exceptions.UnregisteredModuleException import UnreginsteredModuleException


class NetSSProvider(NetworkModule):
    
    def __init__(self):
        NetworkModule.__init__(self)

    def initialRun(self,requested_module):
        print("Would you like me to show you all avaliable connections list? [y/n]")
        try:
            person_choise = input();
        except:
            KeyboardIterruptedHandler.onKeyBoardIterrupted()

        if(person_choise == 'y'):
            self.mShell.executeNoHup([requested_module,'-tunap'])
            return self.prepareIPS(self.pidChoise())
        elif(person_choise == 'n'):
            return self.prepareIPS(self.pidChoise())
        else:
            self.initialRun();  


    def prepareIPS(self,obj):
        command = "netstat -tunapl | awk '/"+obj+"/ {print $5}' "
        self.ipList = self.mShell.execute(command)
        self.ipList = self.ipList[:-1].split('\n');
        self.executeIterations(self.askIpLength())

    def onModuleChoose(self,requested_module):
        if requested_module != 'netstat' and requested_module != 'ss':
            raise UnreginsteredModuleException('netstat or ss modules expected')
        else:
            self.initialRun(requested_module)                 