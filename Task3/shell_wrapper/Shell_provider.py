import sys
import subprocess

class ShellProvider:

    @classmethod
    def executeNoHup(*args):
        try:
            if len(args) >= 3:
                    subprocess.run(args[1],shell=args[2])
            else : 
                subprocess.run(args[1])    
        except Exception as e:
                print("You dont have requested module installed yet")
                sys.exit()

    @classmethod
    def execute(*args):
        shellresult = subprocess.run(args[1], stdout=subprocess.PIPE, stderr=subprocess.DEVNULL, shell=True).stdout.decode('utf-8')
        return shellresult