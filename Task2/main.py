import emoji
import urllib.request, json 
from flask import Flask
from threading import Thread
from flask.globals import request
from flask.templating import render_template
app = Flask(__name__)
myname = "%your name";

def initiateapp80():
    app.run(debug=True,threaded=True,host="0.0.0.0",port=80,use_reloader=False)
def initiateapp8080():
    app.run(debug=True,threaded=True,host="0.0.0.0",port=443,use_reloader=False,ssl_context=('certificate/cert.pem', 'certificate/key.pem'))

#Fetching emoju from openAPI
def fetchemoji():
    with urllib.request.urlopen("https://ranmoji.herokuapp.com/emojis/api/v.1.0/") as url:
        data = json.loads(url.read().decode())
        if data is not None:
            return data['emoji'].split(';',1)[0];
        else:
             return None    


@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def request80(path):
    if(request.method=="POST"):
        jsonrequest = request.get_json(True);
        if 'animal' not in jsonrequest:
            print('animal not found in request')
            return('animal not found in request')
        elif 'sound' not in jsonrequest:
            print('sound not found in request')
            return('sound not found in request')
        elif 'count' not in jsonrequest: 
            print('count not found in request')
            return('count not found in request')  
        else:
            animal = jsonrequest['animal']
            sound = jsonrequest['sound']
            count = jsonrequest['count']
            if type(count) != int:
                print("int count expected")
                return("int count expected")
            else:
                returnstring = ''
                for x in range(count):
                    returnstring+=emoji.emojize(':'+animal+':') + animal + ' says ' + sound + '\n'
                    print(emoji.emojize(':'+animal+':')+animal + ' says ' + sound);
                returnstring+='Made with ' + emoji.emojize(':kissing_face:') + ' by '+ myname;
                print('Made with ' + myname);    
                return returnstring;    
    elif(request.method=="GET"):

        randomEmoji = fetchemoji();
        if randomEmoji is None:
            randomEmoji = ""
        if len(path) > 0:
            return render_template("home.html",value=path+' section',emoji=randomEmoji[1:])
        else:
            return render_template("home.html",value="Home Page",emoji=randomEmoji[1:])


if __name__=="__main__":
   thread1 = Thread(target=initiateapp80)
   thread2 = Thread(target=initiateapp8080)
   thread1.start()
   thread2.start()
