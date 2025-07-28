export const NEXT_PUBLIC_API_BASE_URL = process.env.NEXT_PUBLIC_SERVER

export const REGISTER_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/register`
export const LOGIN_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/login`
export const PROFILE_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/v2/profile`
export const LOGOUT_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/logout`
export const UPLOAD_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/upload`
export const DOWNLOAD_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/storage/{filename}`
export const DELETE_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/delete/{filename}`
export const GET_USERS_URL = `${NEXT_PUBLIC_API_BASE_URL}/api/admin/users/api/admin/users`
