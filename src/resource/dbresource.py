SELECT_USER_BY_EMAIL = "SELECT id, email, name, password, is_admin FROM user WHERE email = ?"
INSERT_USER = "INSERT INTO user (email, password, name, is_admin) VALUES (?, ?, ?, 0)"
SELECT_STORAGE_INFO = "SELECT current_usage, file_count, max_storage FROM user_storage WHERE user_id = ?"
CHECK_STORAGE_LIMIT = "SELECT current_usage, max_storage FROM user_storage WHERE user_id = ?"
UPDATE_STORAGE = """
    UPDATE user_storage
    SET 
        current_usage = MAX(current_usage + ?, 0),
        file_count = MAX(file_count + CASE WHEN ? > 0 THEN 1 ELSE -1 END, 0)
    WHERE user_id = ?
"""
INCREASE_STORAGE_LIMIT = "UPDATE user_storage SET max_storage = ? WHERE user_id = ?"
UPDATE_USER_ADMIN_STATUS = "UPDATE user SET is_admin = ? WHERE id = ?"
SELECT_USER_ADMIN_STATUS = "SELECT is_admin FROM user WHERE id = ?"
SELECT_ALL_USERS = "SELECT id, email, name, is_admin FROM user"
UPDATE_USER = "UPDATE user SET {updates} WHERE id = ?"
SELECT_USER_INFO = """
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
"""