import sqlite3
from werkzeug.security import generate_password_hash, check_password_hash
from src.repositories.user_repository import UserRepository

class AuthService:
    """Сервис аутентификации и регистрации пользователей."""

    @staticmethod
    def login(email: str, password: str):
        """
        Пытается выполнить вход пользователя по email и паролю.
        Возвращает объект пользователя или None, если аутентификация неуспешна.
        """
        user = UserRepository.get_user_by_email(email)
        if user and check_password_hash(user['password'], password):
            return user
        return None

    @staticmethod
    def register(email: str, password: str, name: str):
        """
        Регистрирует нового пользователя.
        Возвращает объект пользователя, если регистрация успешна, или None, если email уже используется.
        """
        if UserRepository.get_user_by_email(email):
            return None  # Пользователь с таким email уже существует

        hashed_password = generate_password_hash(password)
        return UserRepository.create_user(email, hashed_password, name)
