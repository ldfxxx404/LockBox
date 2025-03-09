import os
import logging
from flask import Flask
from src.blueprints.routes import main_blueprint, auth_blueprint
from config import ProductionConfig

def create_app(config_object=ProductionConfig):
    app = Flask(__name__)
    app.config.from_object(config_object)
    
    app.register_blueprint(auth_blueprint, url_prefix='/auth')
    app.register_blueprint(main_blueprint)
    
    # Ensure the storage directory exists
    if not os.path.exists(app.config['UPLOAD_FOLDER']):
        os.makedirs(app.config['UPLOAD_FOLDER'])
    
    # Set up logging
    logging.basicConfig(level=logging.DEBUG)
    
    return app

if __name__ == '__main__':
    app = create_app()
    host = os.getenv('IP_ADDRESS', '0.0.0.0')
    port = int(os.getenv('PORT', 5000))
    app.run(host=host, port=port)