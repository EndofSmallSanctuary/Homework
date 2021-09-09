import subprocess

class ShellProvider:

    @classmethod
    def executeNoHup(*args):
        if len(args) >= 3:
                subprocess.run(args[1],shell=args[2])
        else : 
            subprocess.run(args[1])    

    @classmethod
    def prepareCommand(c_context):
        command = ""

    @classmethod
    def execute(*args):
        shellresult = subprocess.run(args[1], stdout=subprocess.PIPE, stderr=subprocess.DEVNULL, shell=True).stdout.decode('utf-8')
        return shellresult