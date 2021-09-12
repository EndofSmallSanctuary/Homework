from shell_wrapper.Shell_provider import ShellProvider
from handlers.KIHander import KeyboardIterruptedHandler
class NetworkModule:
    
    

    def __init__(self):
        self.mShell = ShellProvider
        #relatedLines
        self.ipList = ''

    def pidChoise(self):
        print("Now it's up to you to choose: i'm familiar with either pid or procesname, so type any")
        try:
            minput = input()
        except:
            KeyboardIterruptedHandler.onKeyBoardIterrupted()
            return ''
        self.mShell.executeNoHup(['clear']);
        return minput
            

    def executeIterations(self):
    #cut -d: -f1
        for idx,ip_port in enumerate(self.ipList):
            self.ipList[idx] = self.ipList[idx].split(':')[0]
        #sort
        self.ipList.sort()
        #uniq -c
        self.ipList = list(set(self.ipList))   
        #sort
        self.ipList.sort()
        #tail -n5

        print("How many lines of selected ip's you would like me to cut at tail?")
        print(self.ipList)
        try:
            tlength = int(input())
        except Exception as e:
            print(e)
            tlength = 0


        
        if len(self.ipList) >= tlength:
            self.ipList = self.ipList[-tlength:]
        else: 
            print("That's much. I shall keep all connections i've got then")

        activeConnectionsNum = {}

        print("Which parameter you'd like me to find with whois? [all/param name]")
        try:
            person_choise = str(input()).lower();
        except:
            KeyboardIterruptedHandler.onKeyBoardIterrupted()

        if person_choise == "all":
            for ip in self.ipList:
                command = "whois " + ip
                self.mShell.executeNoHup(command,True)
        elif person_choise == "organization": 
            onOrganizationExpected(self)
        else:
            for ip in self.ipList:
                command = "whois "+ip+" | awk '/^"+person_choise+"/ {print $2}'"
                self.mShell.executeNoHup(command,True)         


        
def onOrganizationExpected(self):

        activeConnectionsNum = {}

        for ip in self.ipList:
            command = "whois "+ip+" | awk '/^Organization/ {print $2}'"
            organization = self.mShell.execute(command)[:-1]
            if len(organization) == 0:
                organization = 'No organization'
            if organization in activeConnectionsNum:
                activeConnectionsNum[organization] +=1
            else:
                activeConnectionsNum[organization] =1

        
        print("SUMMARY:")
        for conn in activeConnectionsNum:
            print(conn + " : " + str(activeConnectionsNum[conn]))