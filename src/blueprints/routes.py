import os
import logging
from flask import Blueprint, request, jsonify, session, send_from_directory, abort, current_app
from werkzeug.utils import secure_filename
from functools import wraps
from urllib.parse import unquote
from src.services.auth_service import AuthService
from src.repositories.user_repository import UserRepository

def allowed_file(filename: str) -> bool:
    """Проверяет, разрешено ли загружать данный файл."""
    allowed = current_app.config.get('ALLOWED_EXTENSIONS', set())
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in allowed

def login_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'user' not in session:
            return jsonify({"error": "Unauthorized"}), 401
        return f(*args, **kwargs)
    return decorated_function

def admin_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'user_id' not in session or not UserRepository.is_admin(session['user_id']):
            return jsonify({"error": "Forbidden"}), 403
        return f(*args, **kwargs)
    return decorated_function

def get_user_directory(user_id: int) -> str:
    """Возвращает путь к директории пользователя."""
    user_dir = os.path.join(current_app.config['UPLOAD_FOLDER'], str(user_id))
    os.makedirs(user_dir, exist_ok=True)
    return user_dir

def save_uploaded_file(file, user_id: int):
    """Сохраняет загруженный файл и обновляет статистику хранилища"""
    file.seek(0, os.SEEK_END)
    file_size = file.tell()
    file.seek(0)
    
    if not UserRepository.check_storage_limit(user_id, file_size):
        raise ValueError("Превышен лимит хранилища")
    
    filename = secure_filename(file.filename)
    filepath = os.path.join(get_user_directory(user_id), filename)
    file.save(filepath)
    UserRepository.update_storage(user_id, file_size)
    return filename

api_blueprint = Blueprint('api', __name__, url_prefix='/api')

@api_blueprint.route('/login', methods=['POST'])
def login():
    data = request.json
    user = AuthService.login(data.get('email'), data.get('password'))
    if user:
        session['user'] = user['email']
        session['user_id'] = user['id']
        return jsonify({"message": "Login successful", "user": user})
    return jsonify({"error": "Invalid email or password"}), 401

@api_blueprint.route('/logout', methods=['POST'])
@login_required
def logout():
    session.clear()
    return jsonify({"message": "Logout successful"})

@api_blueprint.route('/upload', methods=['POST'])
@login_required
def upload_file():
    file = request.files.get('file')
    if not file or file.filename == '':
        return jsonify({"error": "No file selected"}), 400
    try:
        if allowed_file(file.filename):
            filename = save_uploaded_file(file, session['user_id'])
            return jsonify({"message": "File uploaded", "filename": filename})
        return jsonify({"error": "Invalid file type"}), 400
    except ValueError as e:
        return jsonify({"error": str(e)}), 400
    except Exception as e:
        logging.error(str(e))
        return jsonify({"error": "File upload failed"}), 500

@api_blueprint.route('/storage', methods=['GET'])
@login_required
def storage():
    user_id = session['user_id']
    files = os.listdir(get_user_directory(user_id))
    storage_info = UserRepository.get_storage_info(user_id)
    return jsonify({"files": files, "storage": storage_info})

@api_blueprint.route('/storage/<path:filename>', methods=['GET'])
@login_required
def download_file(filename: str):
    user_id = session['user_id']
    safe_filename = secure_filename(unquote(filename))
    file_path = os.path.join(get_user_directory(user_id), safe_filename)
    if os.path.exists(file_path):
        return send_from_directory(get_user_directory(user_id), safe_filename, as_attachment=True)
    abort(404)

@api_blueprint.route('/delete/<path:filename>', methods=['DELETE'])
@login_required
def delete_file(filename: str):
    user_id = session['user_id']
    safe_filename = secure_filename(unquote(filename))
    file_path = os.path.join(get_user_directory(user_id), safe_filename)
    if os.path.exists(file_path):
        file_size = os.path.getsize(file_path)
        os.remove(file_path)
        UserRepository.update_storage(user_id, -file_size)
        return jsonify({"message": "File deleted"})
    return jsonify({"error": "File not found"}), 404

@api_blueprint.route('/admin/users', methods=['GET'])
@admin_required
def admin_get_users():
    users = UserRepository.get_all_users()
    return jsonify(users)

@api_blueprint.route('/admin/update_limit', methods=['POST'])
@admin_required
def update_storage_limit():
    data = request.json
    user_id = data.get('user_id')
    new_limit = data.get('new_limit')
    try:
        UserRepository.increase_storage_limit(user_id, new_limit)
        return jsonify({"message": "Storage limit updated"})
    except Exception as e:
        return jsonify({"error": str(e)}), 400
