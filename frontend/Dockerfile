# Используем Node.js для сборки
FROM node:18 AS builder

WORKDIR /app

# Копируем package.json и устанавливаем зависимости
COPY package.json package-lock.json ./
RUN npm install

# Копируем весь код и собираем проект
COPY . .
RUN npm run build -- --no-lint

# Используем production-окружение с Node.js
FROM node:18 AS runner

WORKDIR /app

# Копируем собранный код
COPY --from=builder /app ./

# Устанавливаем переменные окружения
ENV NODE_ENV=production
ENV NEXT_PUBLIC_API_URL=http://backend:5000

# Открываем порт Next.js
EXPOSE 3000

# Запускаем Next.js сервер
CMD ["npm", "start"]
