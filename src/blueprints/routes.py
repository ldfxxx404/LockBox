import os
import logging
from flask import Blueprint, render_template, redirect, url_for, request, flash, session, send_from_directory, abort, current_app
from werkzeug.utils import secure_filename
from src.services.auth_service import AuthService
from functools import wraps
from urllib.parse import unquote

def allowed_file(filename: str) -> bool:
    """Проверяет, разрешено ли загружать данный файл."""
    allowed = current_app.config.get('ALLOWED_EXTENSIONS', set())
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in allowed

def login_required(f):
    """Декоратор для проверки авторизации пользователя."""
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'user' not in session:
            return redirect(url_for('auth.login', next=request.url))
        return f(*args, **kwargs)
    return decorated_function

def get_user_directory(user_id: int) -> str:
    """Возвращает путь к директории пользователя."""
    user_dir = os.path.join(current_app.config['UPLOAD_FOLDER'], str(user_id))
    os.makedirs(user_dir, exist_ok=True)
    return user_dir

def save_uploaded_file(file, user_id: int):
    """Сохраняет загруженный файл в директорию пользователя."""
    filename = secure_filename(file.filename)
    filepath = os.path.join(get_user_directory(user_id), filename)
    file.save(filepath)
    logging.info(f'File saved to {filepath}')
    return filename

def get_user_files(user_id: int):
    """Возвращает список файлов пользователя."""
    user_dir = get_user_directory(user_id)
    try:
        return os.listdir(user_dir)
    except OSError:
        return []

main_blueprint = Blueprint('main', __name__)
auth_blueprint = Blueprint('auth', __name__, url_prefix='/auth')

# Authentication Routes
@auth_blueprint.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        email = request.form.get('email')
        password = request.form.get('password')
        user = AuthService.login(email, password)
        if user:
            session['user'] = user['email']
            session['user_id'] = user['id']
            return redirect(url_for('main.home'))
        flash('Invalid email or password', 'error')
    return render_template('login.html')

@auth_blueprint.route('/signup', methods=['GET', 'POST'])
def signup():
    if request.method == 'POST':
        email = request.form.get('email')
        password = request.form.get('password')
        name = request.form.get('name')
        user = AuthService.register(email, password, name)
        if user:
            get_user_directory(user['id'])
            flash('Registration successful! You can now log in.', 'success')
            return redirect(url_for('auth.login'))
        flash('Email address already exists. Please choose a different one.', 'error')
    return render_template('signup.html')

@auth_blueprint.route('/logout')
def logout():
    session.pop('user', None)
    session.pop('user_id', None)
    return redirect(url_for('auth.login'))

# Main Routes
@main_blueprint.route("/")
def home():
    return render_template('index.html')

@main_blueprint.route("/success")
def success():
    return render_template("success.html")

@main_blueprint.route('/upload', methods=['GET', 'POST'])
@login_required
def upload_file():
    if request.method == 'POST':
        file = request.files.get('file')
        if not file or file.filename == '':
            return "No selected file", 400

        if allowed_file(file.filename):
            filename = save_uploaded_file(file, session['user_id'])
            return render_template('success.html', filename=filename), 200

        return render_template('error.html'), 400
    return render_template('upload.html')

@main_blueprint.route('/storage')
@login_required
def storage():
    files = get_user_files(session['user_id'])
    return render_template('storage.html', files=files)

@main_blueprint.route('/storage/<path:filename>')
@login_required
def download_file(filename: str):
    user_id = session['user_id']
    safe_filename = secure_filename(unquote(filename))
    file_path = os.path.join(get_user_directory(user_id), safe_filename)
    if os.path.exists(file_path):
        return send_from_directory(get_user_directory(user_id), safe_filename, as_attachment=True)
    abort(404)

@main_blueprint.route('/profile')
@login_required
def profile():
    return render_template('profile.html')
