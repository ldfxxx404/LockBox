# Use the official Python image from the Docker Hub
FROM python:3.12-slim

# Set the working directory in the container
WORKDIR /app

# Copy the requirements file into the container
COPY . .

# Install the required packages
RUN pip install --no-cache-dir -r req.txt

EXPOSE 1337

# Command to run your application
CMD ["python", "server_start.py"]
