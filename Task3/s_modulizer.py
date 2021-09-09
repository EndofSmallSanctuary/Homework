from Shell_provider import ShellProvider as mShell


def initialRun():
    print("Do you like me to show you all avaliable connections list? [y/n]")
    person_choise = input();
    if(person_choise is 'y'):
        mShell.executeNoHup(['netstat','-tunapl'])
        return prepareIPS(pidChoise())
    elif(person_choise is 'n'):
        return prepareIPS(pidChoise())
    else:
        initialRun();   

def pidChoise():
    print("Now it's up to you to choose: i'm familiar with either pid or procesname, so type any")
    minput = input()
    mShell.executeNoHup(['clear']);
    return minput


def prepareIPS(obj):
    command = "netstat -tunapl | awk '/"+obj+"/ {print $5}' "
    ret = mShell.execute(command)
    ret = ret[:-1].split('\n');
    setUPTheLadder(ret)



def setUPTheLadder(ret):
    #cut -d: -f1
    for idx,ip_port in enumerate(ret):
        ret[idx] = ret[idx].split(':')[0]
    #sort
    ret.sort()
    #uniq -c
    ret = list(set(ret))   
    #sort
    ret.sort()
    #tail -n5
    if len(ret) >= 5:
        ret = ret[-5:]
    print(ret)
    #grep -oP
    # is redutant

    for ip in ret:
        print(ip)
        command = "whois "+ip+" | awk '/^Organization/ {print $2}'"
        mShell.executeNoHup(command,True)


# def shellTest():
#     subprocess.run(["ls", "-l"])

if __name__ == "__main__" :
    initialRun()