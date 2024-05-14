FROM python:3.12.1

WORKDIR /app

RUN pip install poetry

COPY ./poetry.lock pyproject.toml ./

RUN poetry install --no-interaction --no-cache --no-root

COPY . .

CMD ["poetry", "run", "python", "-m", "bot"]