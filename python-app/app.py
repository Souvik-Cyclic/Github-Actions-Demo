from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello_world():
    return "Hello, Hi gg World! Hi, I am a Python Flask application."


if __name__ == "__main__":
    app.run(debug=True)
