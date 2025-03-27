import axios from 'axios';

const errorHandler = (
    err: unknown,
    defaultErrorMessage: string = "An unexpected error occurred"
): string => {
    if (axios.isAxiosError(err)) {
        const status = err.response?.status;

        switch (status) {
            case 200:
                return "";
            case 400:
                return "Invalid input. Please check your data and try again.";
            case 401:
                return "Invalid login credentials. Please check your email and password.";
            case 403:
                return "Access is forbidden. You do not have permission to perform this action.";
            case 404:
                return "The requested resource was not found. Please check the URL or contact support.";
            case 409:
                return "This email is already registered. Please log in or use a different email.";
            case 500:
                return "Internal server error. Please try again later.";
            default:
                return (err.response?.data as { message?: string })?.message || defaultErrorMessage;
        }
    } else if (err instanceof Error) {
        return err.message;
    }
    return defaultErrorMessage;
};

export default errorHandler;
