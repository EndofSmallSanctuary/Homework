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

    def askIpLength(self):
        print("How many lines of selected ip's you would like me to keep?")
        print(self.ipList)
        try:
            minput = int(input())
        except Exception as e:
            print(e)
            return 0

        if(minput == 'show'):
            print(self.ipList)
            self.askIpLength()    
        elif(type(minput) is not int):
            print('Numeric value expected')
            self.askIpLength()
        else:
            return minput

            

    def executeIterations(self,tlength):
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
        print(tlength)
        print(self.ipList)
        if len(self.ipList) >= tlength:
            self.ipList = self.ipList[-tlength:]
        print(self.ipList)
        #grep -oP
        # is redutant

        for ip in self.ipList:
            command = "whois "+ip+" | awk '/^Organization/ {print $2}'"
            self.mShell.executeNoHup(command,True)