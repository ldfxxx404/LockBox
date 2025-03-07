import sqlite3
from config import DATABASE_PATH

class UserRepository:
    """Класс для работы с сущностью User в базе данных."""

    @staticmethod
    def get_user_by_email(email: str):
        """Возвращает пользователя по email или None, если не найден."""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM user WHERE email = ?", (email,))
        user = cursor.fetchone()
        conn.close()
        if user:
            return {
                'id': user[0],
                'email': user[1],
                'password': user[2],
                'name': user[3]
            }
        return None

    @staticmethod
    def create_user(email: str, password: str, name: str):
        """Создает нового пользователя и сохраняет его в базе данных."""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute(
            "INSERT INTO user (email, password, name) VALUES (?, ?, ?)", 
            (email, password, name)
        )
        conn.commit()
        user_id = cursor.lastrowid
        conn.close()
        return {
            'id': user_id,
            'email': email,
            'password': password,
            'name': name
        }