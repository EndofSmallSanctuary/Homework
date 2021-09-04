from flask import Flask
from flask.globals import request
app = Flask(__name__)


@app.route("/",methods=['GET', 'POST'])
def request80():
    if(request.method=="POST"):
        print(request.get_json(True))
        return "dogs"
    elif(request.method=="GET"):
        return "wolfves"


if __name__=="__main__":
    app.run("localhost",80)