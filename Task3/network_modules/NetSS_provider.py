from network_modules.NetworkModule import NetworkModule
from handlers.KIHander import KeyboardIterruptedHandler
import sys
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
            return self.prepareIPS(requested_module,self.pidChoise())
        elif(person_choise == 'n'):
            return self.prepareIPS(requested_module,self.pidChoise())
        else:
            self.initialRun(requested_module);  


    def prepareIPS(self,requested_module,obj):
        if requested_module == 'netstat':
            command = "netstat -tunapl | awk '/"+obj+"/ {print $5}' "
        else:
            command = "ss -tunap | awk '/"+obj+"/ {print $6}' "
        self.ipList = self.mShell.execute(command)
        self.ipList = self.ipList[:-1].split('\n');
        if len(self.ipList) ==1  and self.ipList[0] == '':
                print("no matching pid/processes found")
                sys.exit()
        self.executeIterations()

    def onModuleChoose(self,requested_module):
        if requested_module != 'netstat' and requested_module != 'ss':
            raise UnreginsteredModuleException('netstat or ss modules expected')
        else:
            self.initialRun(requested_module)                 