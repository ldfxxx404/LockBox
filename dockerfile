FROM python:3.12-slim

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r req.txt

EXPOSE 5000

CMD ["python", "main.py"]
