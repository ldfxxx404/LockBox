from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from .config import Config
import os

# Инициализация базы данных
db = SQLAlchemy()

def create_app():
    app = Flask(__name__, instance_relative_config=True, template_folder='../templates')  # Укажите путь к шаблонам
    
    # Настройка конфигурации приложения
    app.config.from_object(Config)  # Загружаем конфигурацию из config.py
    app.secret_key = Config.SECRET_KEY 
    # Инициализация базы данных с приложением
    db.init_app(app)

    # Регистрация блюпринтов
    from .auth import auth as auth_blueprint
    from .main import main as main_blueprint
    app.register_blueprint(auth_blueprint)
    app.register_blueprint(main_blueprint)

    # Создание папки для загрузок, если она не существует
    if not os.path.exists(app.config['UPLOAD_FOLDER']):
        os.makedirs(app.config['UPLOAD_FOLDER'])

    with app.app_context():
        db.create_all()  # Создание всех таблиц

    return app
