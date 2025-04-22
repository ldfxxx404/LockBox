interface Environment {
  NEXT_PUBLIC_API_URL: string
}

export const Env: Environment = {
  NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL as string,
}
