import os
import logging
from flask import Blueprint, render_template, redirect, url_for, request, flash, session, send_from_directory, abort, current_app
from werkzeug.utils import secure_filename
from src.services.auth_service import AuthService
from functools import wraps
from urllib.parse import unquote
from src.repositories.user_repository import UserRepository

def allowed_file(filename: str) -> bool:
    """Проверяет, разрешено ли загружать данный файл."""
    allowed = current_app.config.get('ALLOWED_EXTENSIONS', set())
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in allowed

def admin_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'user_id' not in session or not UserRepository.is_admin(session['user_id']):
            return redirect(url_for('auth.login', next=request.url))
        return f(*args, **kwargs)
    return decorated_function


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
    """Сохраняет загруженный файл и обновляет статистику хранилища"""
    # Получаем размер файла ДО сохранения
    file.seek(0, os.SEEK_END)
    file_size = file.tell()
    file.seek(0)
    
    # Проверяем лимиты хранилища
    try:
        if not UserRepository.check_storage_limit(user_id, file_size):
            raise ValueError("Превышен лимит хранилища")
        
        filename = secure_filename(file.filename)
        filepath = os.path.join(get_user_directory(user_id), filename)
        file.save(filepath)
        
        # Обновляем статистику после успешного сохранения
        UserRepository.update_storage(user_id, file_size)
        
        logging.info(f'File {filename} ({file_size} bytes) saved to {filepath}')
        return filename
    except Exception as e:
        logging.error(f"File upload error: {str(e)}")
        raise

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
    user_id = session.get('user_id')
    is_admin = UserRepository.is_admin(user_id) if user_id else False
    return render_template('index.html', is_admin=is_admin)

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

        try:
            if allowed_file(file.filename):
                filename = save_uploaded_file(file, session['user_id'])
                return render_template('success.html', filename=filename), 200
            return render_template('error.html', message="Недопустимый тип файла"), 400
        except ValueError as e:
            return render_template('error.html', message=str(e)), 400
        except Exception as e:
            logging.error(str(e))
            return render_template('error.html', message="Ошибка загрузки файла"), 500

    return render_template('upload.html')

@main_blueprint.route('/storage')
@login_required
def storage():
    user_id = session['user_id']
    files = get_user_files(user_id)
    storage_info = UserRepository.get_storage_info(user_id)
    return render_template(
        'storage.html',
        files=files,
        storage=storage_info,
        used_percent=round((storage_info['current_usage'] / storage_info['max_storage']) * 100, 2)
    )

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
    user_id = session['user_id']
    storage_info = UserRepository.get_storage_info(user_id)
    return render_template(
        'profile.html',
        storage=storage_info,
        used_percent=round((storage_info['current_usage'] / storage_info['max_storage']) * 100, 2)
    )

@main_blueprint.route('/delete/<path:filename>', methods=['POST'])
@login_required
def delete_file(filename: str):
    user_id = session['user_id']
    safe_filename = secure_filename(unquote(filename))
    file_path = os.path.join(get_user_directory(user_id), safe_filename)
    
    try:
        if os.path.exists(file_path):
            file_size = os.path.getsize(file_path)
            os.remove(file_path)
            UserRepository.update_storage(user_id, -file_size)
            flash('Файл успешно удален', 'success')
        else:
            flash('Файл не найден', 'error')
    except Exception as e:
        logging.error(f"File deletion error: {str(e)}")
        flash('Ошибка при удалении файла', 'error')
    
    return redirect(url_for('main.storage'))

def get_user_files(user_id: int):
    """Возвращает список файлов пользователя."""
    user_dir = get_user_directory(user_id)
    try:
        return os.listdir(user_dir)
    except OSError:
        return []

@main_blueprint.route('/admin/update_limit', methods=['GET', 'POST'])
@login_required
def update_storage_limit():
    if request.method == 'POST':
        user_id = int(request.form.get('user_id'))
        new_limit = int(request.form.get('new_limit'))
        
        try:
            UserRepository.increase_storage_limit(user_id, new_limit)
            flash('Лимит хранилища успешно обновлен', 'success')
        except Exception as e:
            flash('Ошибка при обновлении лимита', 'error')
    
    return render_template('update_limit.html')

@main_blueprint.route('/admin')
@admin_required
def admin_panel():
    users = UserRepository.get_all_users()
    return render_template('admin.html', users=users)

@main_blueprint.route('/admin/edit_user/<int:user_id>', methods=['GET', 'POST'])
@admin_required
def edit_user(user_id: int):
    user_info = UserRepository.get_user_info(user_id)
    if not user_info:
        abort(404)

    if request.method == 'POST':
        email = request.form.get('email')
        name = request.form.get('name')
        password = request.form.get('password') # TODO: Hash this!
        is_admin = request.form.get('is_admin') == 'on'
        max_storage = int(request.form.get('max_storage'))  # Получаем значение лимита

        try:
            UserRepository.update_user(user_id, email, name, password, is_admin)
            UserRepository.increase_storage_limit(user_id, max_storage) #Обновляем лимит
            return redirect(url_for('main.admin_panel'))
        except Exception as e:
            return render_template('edit_user.html', user_info=user_info, error=str(e))


    return render_template('edit_user.html', user_info=user_info)


@main_blueprint.route('/admin/make_admin/<int:user_id>')
@admin_required
def make_admin(user_id: int):
    UserRepository.make_admin(user_id)
    return redirect(url_for('main.admin_panel'))

@main_blueprint.route('/admin/revoke_admin/<int:user_id>')
@admin_required
def revoke_admin(user_id: int):
    UserRepository.revoke_admin(user_id)
    return redirect(url_for('main.admin_panel'))
