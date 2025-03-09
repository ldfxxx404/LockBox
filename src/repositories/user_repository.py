import sqlite3
from config import DATABASE_PATH

class UserRepository:
    @staticmethod
    def get_user_by_email(email: str):
        """Возвращает пользователя по email или None, если не найден."""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute("SELECT id, email, name, password, is_admin FROM user WHERE email = ?", (email,))
        user = cursor.fetchone()
        conn.close()
        if user:
            return {
                'id': user[0],
                'email': user[1],
                'name': user[2],
                'password': user[3], # Make sure password is there
                'is_admin': bool(user[4])
            }
        return None  # Return None if user not found

    @staticmethod
    def create_user(email: str, password: str, name: str):
        """Создает нового пользователя и сохраняет его в базе данных."""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute("INSERT INTO user (email, password, name, is_admin) VALUES (?, ?, ?, 0)", (email, password, name))
        conn.commit()
        user_id = cursor.lastrowid
        conn.close()
        return {'id': user_id, 'email': email, 'name': name, 'is_admin': False}  # Явное возвращение is_admin

    @staticmethod
    def get_storage_info(user_id: int):
        """Получает информацию о хранилище пользователя"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute("SELECT current_usage, file_count, max_storage FROM user_storage WHERE user_id = ?", (user_id,))
        result = cursor.fetchone()
        conn.close()
        if result:
            return {'current_usage': result[0], 'file_count': result[1], 'max_storage': result[2]}
        else:
            return None  # Важно вернуть None, если записи нет

    @staticmethod
    def check_storage_limit(user_id: int, file_size: int) -> bool:
        """Проверяет, можно ли загрузить файл без превышения лимита"""
        storage = UserRepository.get_storage_info(user_id)
        if not storage:
            return False  # Если нет инфы о хранилище, считаем, что нельзя

        return (storage['current_usage'] + file_size) <= storage['max_storage']

    @staticmethod
    def update_storage(user_id: int, delta: int):
        """Обновляет статистику хранилища"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        try:
            cursor.execute("""
                UPDATE user_storage
                SET 
                    current_usage = MAX(current_usage + ?, 0),
                    file_count = MAX(file_count + CASE WHEN ? > 0 THEN 1 ELSE -1 END, 0)
                WHERE user_id = ?
            """, (delta, delta, user_id))
            conn.commit()
        except Exception as e:
            conn.rollback()
            raise
        finally:
            conn.close()

    @staticmethod
    def increase_storage_limit(user_id: int, new_limit: int):
        """Увеличивает лимит хранилища для пользователя"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        try:
            cursor.execute("UPDATE user_storage SET max_storage = ? WHERE user_id = ?", (new_limit, user_id))
            conn.commit()
        except Exception as e:
            conn.rollback()
            raise
        finally:
            conn.close()

    @staticmethod
    def make_admin(user_id: int):
        """Назначает пользователя администратором"""
        UserRepository._set_admin_status(user_id, True)  # Используем приватный метод

    @staticmethod
    def revoke_admin(user_id: int):
        """Отзывает права администратора"""
        UserRepository._set_admin_status(user_id, False) # Используем приватный метод

    @staticmethod
    def _set_admin_status(user_id: int, is_admin: bool):
        """Вспомогательный метод для установки статуса администратора"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        try:
            cursor.execute("UPDATE user SET is_admin = ? WHERE id = ?", (int(is_admin), user_id))
            conn.commit()
        except Exception as e:
            conn.rollback()
            raise
        finally:
            conn.close()

    @staticmethod
    def is_admin(user_id: int) -> bool:
        """Проверяет, является ли пользователь администратором"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute("SELECT is_admin FROM user WHERE id = ?", (user_id,))
        result = cursor.fetchone()
        conn.close()
        return bool(result[0]) if result else False

    @staticmethod
    def get_all_users():
        """Возвращает список всех пользователей"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        cursor.execute("SELECT id, email, name, is_admin FROM user")
        users = cursor.fetchall()
        conn.close()
        return [{'id': u[0], 'email': u[1], 'name': u[2], 'is_admin': bool(u[3])} for u in users]

    @staticmethod
    def update_user(user_id: int, email: str = None, name: str = None, password: str = None, is_admin: bool = None):
        """Обновляет информацию о пользователе (email, name, password, is_admin)"""
        conn = sqlite3.connect(DATABASE_PATH)
        cursor = conn.cursor()
        try:
            updates = []
            params = []

            if email:
                updates.append("email = ?")
                params.append(email)
            if name:
                updates.append("name = ?")
                params.append(name)
            if password:  # Хэшируйте пароль перед сохранением!
                updates.append("password = ?")
                params.append(password)
            if is_admin is not None:
                updates.append("is_admin = ?")
                params.append(int(is_admin))

            if updates:
                query = f"UPDATE user SET {', '.join(updates)} WHERE id = ?"
                params.append(user_id)
                cursor.execute(query, params)
                conn.commit()

        except Exception as e:
            conn.rollback()
            raise
        finally:
            conn.close()


    @staticmethod
    def get_user_info(user_id: int):
         """Возвращает информацию о пользователе, включая размер хранилища."""
         conn = sqlite3.connect(DATABASE_PATH)
         cursor = conn.cursor()

         cursor.execute("""
             SELECT 
                 u.email, 
                 u.name, 
                 u.is_admin, 
                 us.current_usage, 
                 us.max_storage,
                 u.password
             FROM user u
             LEFT JOIN user_storage us ON u.id = us.user_id
             WHERE u.id = ?
         """, (user_id,))

         result = cursor.fetchone()
         conn.close()

         if result:
             return {
                 'email': result[0],
                 'name': result[1],
                 'is_admin': bool(result[2]),
                 'current_usage': result[3] or 0,  # Handle None case
                 'max_storage': result[4] or 0,   # Handle None case
                 'password': result[5]
             }
         else:
             return None
