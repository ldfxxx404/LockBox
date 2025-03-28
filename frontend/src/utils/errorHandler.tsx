import { NextResponse } from 'next/server';
import axios from 'axios';

const errorHandler = (
    err: unknown,
    defaultErrorMessage: string = "An unexpected error occurred"
): NextResponse => {  // Return a NextResponse instead of just a string
    let errorMessage = defaultErrorMessage;

    if (axios.isAxiosError(err)) {
        const status = err.response?.status;
        switch (status) {
            case 200:
                errorMessage = "";  // No error
                break;
            case 400:
                errorMessage = "Invalid input. Please check your data and try again.";
                break;
            case 401:
                errorMessage = "Invalid login credentials. Please check your email and password.";
                break;
            case 403:
                errorMessage = "Access is forbidden. You do not have permission to perform this action.";
                break;
            case 404:
                errorMessage = "The requested resource was not found. Please check the URL or contact support.";
                break;
            case 409:
                errorMessage = "This email is already registered. Please log in or use a different email.";
                break;
            case 500:
                errorMessage = "Internal server error. Please try again later.";
                break;
            default:
                errorMessage = (err.response?.data as { message?: string })?.message || defaultErrorMessage;
                break;
        }
    } else if (err instanceof Error) {
        errorMessage = err.message;
    }

    // Return NextResponse with error message as JSON
    return new NextResponse(
        JSON.stringify({ error: errorMessage }),
        { status: 400 } // Adjust status if necessary
    );
};

export default errorHandler;
