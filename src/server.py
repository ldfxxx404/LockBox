from flask import Flask, render_template, request, send_from_directory
from fileinput import filename
import os
from config import *

# Конфиги
app = Flask(__name__, template_folder='../templates')
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

# Проверка файлв
def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

# Маршруты

# Домашняя страница
@app.route('/')
def home(): 
    return render_template ('index.html')

# Удачный аплоад файла
@app.route('/success', methods=['POST'])
def success():
    if 'file' not in request.files:
        return "Request without file", 400
    
    file = request.files['file']
    
    if file.filename == '':
        return render_template('warning.html'), 400
    
    if file and allowed_file(file.filename):
        filepath = f"{UPLOAD_FOLDER}/{file.filename}"
        file.save(filepath)
        return render_template('success.html'), 200  
    
    return render_template('error.html'), 400

# Загрузка файлов
@app.route('/Storage')
def storage():
    files = os.listdir(UPLOAD_FOLDER)
    return render_template('storage.html', files=files)

# Скачивание файла
@app.route('/Storage/<filename>')
def download_file(filename):
    return send_from_directory(UPLOAD_FOLDER, filename)

if __name__ == '__main__': 
    app.run(debug=True, host=IP_ADDRESS, port=PORT)
