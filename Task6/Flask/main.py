import emoji
import urllib.request, json 
from flask import Flask
from threading import Thread
from flask.globals import request
from flask.templating import render_template
app = Flask(__name__)
myname = "%your name";


def initiateapp8080():
    app.run(debug=True,threaded=True,host="0.0.0.0",port=8080,use_reloader=False)

#Fetching emoju from openAPI
def fetchemoji():
    with urllib.request.urlopen("https://ranmoji.herokuapp.com/emojis/api/v.1.0/") as url:
        data = json.loads(url.read().decode())
        if data is not None:
            return data['emoji'].split(';',1)[0];
        else:
             return None    


@app.route('/',methods=['GET','POST'], defaults={'path': ''})
@app.route('/<path:path>',methods=['GET','POST'])
def request80(path):
    if(request.method=="POST"):
        jsonrequest = request.get_json(True);
        if 'animal' not in jsonrequest:
            print('animal not found in request'+'\n')
            return('animal not found in request'+'\n')
        elif 'sound' not in jsonrequest:
            print('sound not found in request'+'\n')
            return('sound not found in request'+'\n')
        elif 'count' not in jsonrequest: 
            print('count not found in request'+'\n')
            return('count not found in request'+'\n')  
        else:
            animal = jsonrequest['animal']
            sound = jsonrequest['sound']
            count = jsonrequest['count']
            if type(count) != int:
                print("int count expected"+'\n')
                return("int count expected"+'\n')
            else:
                returnstring = ''
                animal_emoji= emoji.emojize(':'+animal+':')
                if animal_emoji[0] == ':':
                    animal_emoji = ''
                for x in range(count):
                    returnstring+=animal_emoji + animal + ' says ' + sound + '\n'
                    print(animal_emoji+animal + ' says ' + sound);
                returnstring+='Made with ' + emoji.emojize(':kissing_face:') + ' by '+ myname;
                print('Made with ' + myname+'\n');    
                return returnstring+'\n';    
    elif(request.method=="GET"):

        randomEmoji = fetchemoji();
        if randomEmoji is None:
            randomEmoji = ""
        if len(path) > 0:
            return render_template("home.html",value=path+' section',emoji=randomEmoji[1:])
        else:
            return render_template("home.html",value="Home Page",emoji=randomEmoji[1:])


if __name__=="__main__":
   thread1 = Thread(target=initiateapp8080)
   thread1.start()
