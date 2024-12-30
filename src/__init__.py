import os
from flask import Flask
from .config import UPLOAD_FOLDER, IP_ADDRESS, PORT
from .auth import auth as auth_blueprint
from .main import main as main_blueprint

def create_app():
    app = Flask(__name__, template_folder='../templates')
    app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER
    app.config['ALLOWED_EXTENSIONS'] = {'jpg', 'jpeg', 'webm', 'mp4', 'mp3', 'png', 'txt'}

    # Register blueprints
    app.register_blueprint(auth_blueprint)
    app.register_blueprint(main_blueprint)

    # Create upload folder if it doesn't exist
    if not os.path.exists(UPLOAD_FOLDER):
        os.makedirs(UPLOAD_FOLDER)

    return app
