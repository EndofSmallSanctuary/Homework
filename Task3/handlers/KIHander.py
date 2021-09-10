class KeyboardIterruptedHandler:
    @classmethod
    def onKeyBoardIterrupted(*args):
        if len(args>1):
            print('\n',args[1])
        else:     
            print('\n'+'Goodbye')