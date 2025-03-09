FROM python:3.12-slim

RUN apt-get update && apt-get install -y sqlite3

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r req.txt

EXPOSE 5000

RUN python createdb.py 
CMD [ "python", "main.py" ] 
