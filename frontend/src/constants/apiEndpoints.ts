export const BASE_URL = process.env.NEXT_PUBLIC_SERVER

export const REGISTER_URL = `${BASE_URL}/api/register`
export const LOGIN_URL = `${BASE_URL}/api/login`
export const PROFILE_URL = `${BASE_URL}/api/v2/profile`
export const LOGOUT_URL = `${BASE_URL}/api/logout`
export const UPLOAD_URL = `${BASE_URL}/api/upload`
export const DOWNLOAD_URL = `${BASE_URL}/api/storage/{filename}`
export const DELETE_URL = `${BASE_URL}/api/delete/{filename}`
export const GET_USERS_URL = `${BASE_URL}/api/admin/users`
export const MAKE_ADMIN_URL = `${BASE_URL}/api/admin/make_admin`
export const REVOKE_ADMIN_URL = `${BASE_URL}/api/admin/revoke_admin`
export const UPDATE_LIMIT_URL = `${BASE_URL}/api/admin/update_limit`
export const GET_ADMINS_URL = `${BASE_URL}/api/admin/admins`
