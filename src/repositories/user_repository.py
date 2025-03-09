import sqlite3
from config import DATABASE_PATH
from werkzeug.security import generate_password_hash
from src.resource.dbresource import *

class UserRepository:
    
    @staticmethod
    def _execute_query(query, params=None, fetchone=False, commit=False):
        """Вспомогательная функция для выполнения SQL-запросов."""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        try:
            cursor.execute(query, params or ())
            if commit:
                conn.commit()
            if fetchone:
                return cursor.fetchone()
            else:
                return cursor.fetchall()
        except Exception as e:
            conn.rollback()
            raise
        finally:
            conn.close()

    @staticmethod
    def get_user_by_email(email: str):
        """Возвращает пользователя по email или None, если не найден."""
        query = SELECT_USER_BY_EMAIL
        result = UserRepository._execute_query(query, (email,), fetchone=True)
        if result:
            return {
                'id': result[0],
                'email': result[1],
                'name': result[2],
                'password': result[3],  # Make sure password is there
                'is_admin': bool(result[4])
            }
        return None

    @staticmethod
    def create_user(email: str, password: str, name: str):
        """Создает нового пользователя и сохраняет его в базе данных."""
        query = INSERT_USER
        UserRepository._execute_query(query, (email, password, name), commit=True)
        return {'id': cursor.lastrowid, 'email': email, 'name': name, 'is_admin': False}

    @staticmethod
    def get_storage_info(user_id: int):
        """Получает информацию о хранилище пользователя"""
        query = SELECT_STORAGE_INFO
        result = UserRepository._execute_query(query, (user_id,), fetchone=True)
        if result:
            return {'current_usage': result[0], 'file_count': result[1], 'max_storage': result[2]}
        return None

    @staticmethod
    def check_storage_limit(user_id: int, file_size: int) -> bool:
        """Проверяет, можно ли загрузить файл без превышения лимита"""
        query = CHECK_STORAGE_LIMIT
        storage = UserRepository._execute_query(query, (user_id,), fetchone=True)
        if storage:
            return (storage[0] + file_size) <= storage[1]
        return False

    @staticmethod
    def update_storage(user_id: int, delta: int):
        """Обновляет статистику хранилища"""
        query = UPDATE_STORAGE
        UserRepository._execute_query(query, (delta, delta, user_id), commit=True)

    @staticmethod
    def increase_storage_limit(user_id: int, new_limit: int):
        """Увеличивает лимит хранилища для пользователя"""
        query = INCREASE_STORAGE_LIMIT
        UserRepository._execute_query(query, (new_limit, user_id), commit=True)

    @staticmethod
    def make_admin(user_id: int):
        """Назначает пользователя администратором"""
        UserRepository._set_admin_status(user_id, True)

    @staticmethod
    def revoke_admin(user_id: int):
        """Отзывает права администратора"""
        UserRepository._set_admin_status(user_id, False)

    @staticmethod
    def _set_admin_status(user_id: int, is_admin: bool):
        """Вспомогательный метод для установки статуса администратора"""
        query = UPDATE_USER_ADMIN_STATUS
        UserRepository._execute_query(query, (int(is_admin), user_id), commit=True)

    @staticmethod
    def is_admin(user_id: int) -> bool:
        """Проверяет, является ли пользователь администратором"""
        query = SELECT_USER_ADMIN_STATUS
        result = UserRepository._execute_query(query, (user_id,), fetchone=True)
        return bool(result[0]) if result else False

    @staticmethod
    def get_all_users():
        """Возвращает список всех пользователей"""
        query = SELECT_ALL_USERS
        users = UserRepository._execute_query(query)
        return [{'id': u[0], 'email': u[1], 'name': u[2], 'is_admin': bool(u[3])} for u in users]

    @staticmethod
    def update_user(user_id: int, email: str = None, name: str = None, password: str = None, is_admin: bool = None):
        """Обновляет информацию о пользователе (email, name, password, is_admin)"""
        updates = []
        params = []

        if email:
            updates.append("email = ?")
            params.append(email)
        if name:
            updates.append("name = ?")
            params.append(name)
        if password:
            hashed_password = generate_password_hash(password)
            updates.append("password = ?")
            params.append(hashed_password)
        if is_admin is not None:
            updates.append("is_admin = ?")
            params.append(int(is_admin))

        if updates:
            query = UPDATE_USER.format(updates=', '.join(updates))
            params.append(user_id)
            UserRepository._execute_query(query, params, commit=True)

    @staticmethod
    def get_user_info(user_id: int):
        """Возвращает информацию о пользователе, включая размер хранилища."""
        query = SELECT_USER_INFO
        result = UserRepository._execute_query(query, (user_id,), fetchone=True)
        if result:
            return {
                'email': result[0],
                'name': result[1],
                'is_admin': bool(result[2]),
                'current_usage': result[3] or 0,  # Handle None case
                'max_storage': result[4] or 0,   # Handle None case
                'password': result[5]
            }
        return None
