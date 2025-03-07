import os
import sqlite3
import logging
from flask import Blueprint, render_template, redirect, url_for, request, flash, session, send_from_directory, abort, current_app
from werkzeug.utils import secure_filename
from src.services.auth_service import AuthService
from functools import wraps
from urllib.parse import unquote

main_blueprint = Blueprint('main', __name__)
auth_blueprint = Blueprint('auth', __name__, url_prefix='/auth')

def allowed_file(filename: str) -> bool:
    allowed = current_app.config.get('ALLOWED_EXTENSIONS', set())
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in allowed

def login_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'user' not in session:
            return redirect(url_for('auth.login', next=request.url))
        return f(*args, **kwargs)
    return decorated_function

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
            # Create user directory
            user_dir = os.path.join(current_app.config['UPLOAD_FOLDER'], str(user['id']))
            os.makedirs(user_dir, exist_ok=True)
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
@main_blueprint.route("/", methods=['GET'])
def home():
    return render_template('index.html')

@main_blueprint.route("/success", methods=['GET', 'POST'])
def success():
    return render_template("success.html")

@main_blueprint.route('/upload', methods=['GET', 'POST'])
@login_required
def upload_file():
    if request.method == 'POST':
        if 'file' not in request.files:
            return "No file part in the request", 400

        file = request.files['file']
        if not file or file.filename == '':
            return "No selected file", 400

        if allowed_file(file.filename):
            filename = secure_filename(file.filename)
            user_id = session['user_id']
            user_dir = os.path.join(current_app.config['UPLOAD_FOLDER'], str(user_id))
            os.makedirs(user_dir, exist_ok=True)  # Ensure the directory exists
            filepath = os.path.join(user_dir, filename)
            file.save(filepath)
            logging.info(f'File saved to {filepath}')
            return render_template('success.html', filename=filename), 200

        return render_template('error.html'), 400
    return render_template('upload.html')

@main_blueprint.route('/storage')
@login_required
def storage():
    user_id = session['user_id']
    user_dir = os.path.join(current_app.config['UPLOAD_FOLDER'], str(user_id))
    try:
        files = os.listdir(user_dir)
    except OSError:
        files = []
    return render_template('storage.html', files=files)

@main_blueprint.route('/storage/<path:filename>')
@login_required
def download_file(filename: str):
    user_id = session['user_id']
    decoded_filename = unquote(filename)
    safe_filename = secure_filename(decoded_filename)
    user_dir = os.path.join(current_app.config['UPLOAD_FOLDER'], str(user_id))
    file_path = os.path.join(user_dir, safe_filename)
    if os.path.exists(file_path):
        return send_from_directory(user_dir, safe_filename, as_attachment=True)
    abort(404)

@main_blueprint.route('/profile')
@login_required
def profile():
    return render_template('profile.html')