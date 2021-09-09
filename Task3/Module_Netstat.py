from Task3.s_modulizer import pidChoise


class NetstatProvider:
    
    def __init__(self,*args):
        self.pid = args[0]


    def prepareIPS(obj):
        command = "netstat -tunapl | awk '/"+obj+"/ {print $5}' "
        ret = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.DEVNULL, shell=True).stdout.decode('utf-8')
        print(ret)

