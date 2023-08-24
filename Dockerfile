FROM python:3.11-alpine3.18

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt
RUN rm requirements.txt

COPY src src

CMD ["python3", "-m", "src"]
