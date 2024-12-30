import os
from flask import Blueprint, render_template, request, send_from_directory, abort, current_app
from werkzeug.utils import secure_filename
from urllib.parse import unquote

# Assuming this is imported correctly
from .config import *

main = Blueprint('main', __name__)

# Function to check allowed file extensions
def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

# Home route
@main.route('/')
def home():
    return render_template('index.html')

# Successful file upload route
@main.route('/success', methods=['POST'])
def success():
    if 'file' not in request.files:
        return "No file part in the request", 400

    file = request.files['file']
    if file.filename == '':
        return "No selected file", 400

    if file and allowed_file(file.filename):
        filename = secure_filename(file.filename)
        filepath = os.path.join(current_app.config['UPLOAD_FOLDER'], filename)
        file.save(filepath)
        return render_template('success.html', filename=filename), 200

    return render_template('error.html'), 400

@main.route('/storage')
def storage():
    # Use current_app to get the config
    files = os.listdir(current_app.config['UPLOAD_FOLDER'])
    return render_template('storage.html', files=files)

# File download route
@main.route('/storage/<path:filename>')
def download_file(filename):
    filename = unquote(filename)
    safe_filename = secure_filename(filename)
    file_path = os.path.join(current_app.config['UPLOAD_FOLDER'], safe_filename)
    if os.path.exists(file_path):
        return send_from_directory(current_app.config['UPLOAD_FOLDER'], safe_filename, as_attachment=True)
    else:
        abort(404)

@main.route('/profile')
def profile():
    return render_template('profile.html')
