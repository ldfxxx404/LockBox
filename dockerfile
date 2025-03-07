FROM python:3.12-slim

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r req.txt

EXPOSE 1337

CMD ["python", "main.py"]
