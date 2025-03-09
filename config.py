import os

BASE_DIR = os.path.abspath(os.path.dirname(__file__))
DATABASE_PATH = os.path.join(BASE_DIR, 'instance', 'app.db')
UPLOAD_FOLDER = os.path.join(BASE_DIR, 'storage')

class BaseConfig:
    """Базовая конфигурация для Flask-приложения."""
    SECRET_KEY = os.getenv('SECRET_KEY', 'supersecretkey123')
    UPLOAD_FOLDER = UPLOAD_FOLDER
    ALLOWED_EXTENSIONS = {'png', 'jpg', 'jpeg', 'gif', 'webm', 'mp4', 'mp3', 'txt', 'sql'}

class DevelopmentConfig(BaseConfig):
    """Конфигурация для режима разработки."""
    DEBUG = True

class ProductionConfig(BaseConfig):
    """Конфигурация для продакшена."""
    DEBUG = False