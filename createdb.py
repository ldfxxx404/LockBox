import sqlite3
from config import DATABASE_PATH

def init_db():
    conn = sqlite3.connect(DATABASE_PATH)
    cursor = conn.cursor()
    
    # Таблица пользователей
    cursor.execute("""
    CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        name TEXT NOT NULL,
        is_admin INTEGER DEFAULT 0
    )""")
    
    # Таблица хранилища
    cursor.execute("""
    CREATE TABLE IF NOT EXISTS user_storage (
        user_id INTEGER PRIMARY KEY,
        storage_path TEXT NOT NULL UNIQUE,
        current_usage INTEGER DEFAULT 0,
        file_count INTEGER DEFAULT 0,
        max_storage INTEGER DEFAULT 1073741824,
        FOREIGN KEY(user_id) REFERENCES user(id)
    )""")
    
    # Триггер для автоматического создания записи хранилища
    cursor.execute("""
    CREATE TRIGGER IF NOT EXISTS create_user_storage 
    AFTER INSERT ON user
    BEGIN
        INSERT INTO user_storage(user_id, storage_path)
        VALUES (NEW.id, '/storage/' || NEW.id);
    END;
    """)
    
    conn.commit()
    conn.close()

if __name__ == "__main__":
    init_db()
