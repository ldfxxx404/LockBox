from src import create_app
from src.config import Config

app = create_app()

if __name__ == '__main__':
    app.run(host=Config.IP_ADDRESS, port=Config.PORT, debug=True)
