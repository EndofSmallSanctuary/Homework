from os import truncate
from flask import Flask
from threading import Thread
from flask.globals import request
app = Flask(__name__)


def initiateapp80():
    app.run(debug=True,threaded=True,host="0.0.0.0",port=80,use_reloader=False)
def initiateapp8080():
    app.run(debug=True,threaded=True,host="0.0.0.0",port=443,use_reloader=False,ssl_context=('certificate/cert.pem', 'certificate/key.pem'))


@app.route("/",methods=['GET', 'POST'])
def request80():
    if(request.method=="POST"):
        print(request.get_json(True))
        return "dogs"
    elif(request.method=="GET"):
        return "wolfves"


if __name__=="__main__":
   thread1 = Thread(target=initiateapp80)
   thread2 = Thread(target=initiateapp8080)
   thread1.start()
   thread2.start()
   thread1.join()