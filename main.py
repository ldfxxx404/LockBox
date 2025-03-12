import os
import logging
from flask import Flask
from src.blueprints.routes import api_blueprint
from config import ProductionConfig

def create_app(config_object=ProductionConfig):
    app = Flask(__name__)
    app.config.from_object(config_object)
    
    app.register_blueprint(api_blueprint)
    
    # Ensure the storage directory exists
    os.makedirs(app.config['UPLOAD_FOLDER'], exist_ok=True)
    
    # Set up logging
    logging.basicConfig(level=logging.DEBUG)
    
    return app

if __name__ == '__main__':
    app = create_app()
    host = os.getenv('IP_ADDRESS', '127.0.0.1')
    port = int(os.getenv('PORT', 5000))
    app.run(host=host, port=port)
