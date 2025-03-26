import axios from 'axios';

const errorHandler = (
  err: unknown,
  defaultErrorMessage: string = "An unexpected error occurred"
): string | null => {
  if (axios.isAxiosError(err)) {
    const status = err.response?.status;

    switch (status) {
      case 200:
        return null;
      case 400:
        return "Invalid input. Please check your data and try again.";
      case 401:
        return "You are not authorized. Please log in.";
      case 403:
        return "Access is forbidden. This action is only available for administrators.";
      case 404:
        return "The requested resource was not found.";
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
