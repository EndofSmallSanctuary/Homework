from handlers.KIHander import KeyboardIterruptedHandler
from network_modules.NetSS_provider import NetSSProvider
# def shellTest():
#     subprocess.run(["ls", "-l"])

if __name__ == "__main__" :
    netstat = NetSSProvider()
    print('which module you want me do use? [netstat/ss]')
    netstat.onModuleChoose(input())
    # try:
    #     netstat.onModuleChoose(input())
    # except:
    #     KeyboardIterruptedHandler.onKeyBoardIterrupted()

