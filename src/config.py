import os

BASE_DIR = os.path.abspath(os.path.dirname(__file__))  
DATABASE_PATH = os.path.join(BASE_DIR, '../instance/app.db')

class Config:
    SQLALCHEMY_DATABASE_URI = f'sqlite:///{DATABASE_PATH}'
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    UPLOAD_FOLDER = os.path.join(BASE_DIR, '../storage')  
    IP_ADDRESS = '0.0.0.0'
    PORT = 1337
    ALLOWED_EXTENSIONS = {'jpg', 'jpeg', 'webm', 'mp4', 'mp3', 'png', 'txt'}
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'filaret'  