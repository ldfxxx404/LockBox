# Используем Node.js для сборки
FROM node:18 AS builder

WORKDIR /app

# Копируем package.json и устанавливаем зависимости
COPY package.json package-lock.json ./
RUN npm install

# Копируем весь код и собираем проект
COPY . .
RUN npm run build

# Используем production-окружение с Node.js
FROM node:18 AS runner

WORKDIR /app

# Копируем собранный код
COPY --from=builder /app ./

# Открываем порт Next.js
EXPOSE 3000

# Запускаем Next.js сервер
CMD ["npm", "start"]
