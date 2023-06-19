
FROM python:3.11.4-bullseye

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

# CMD ["python3", "-m" , "flask", "run", "--host=0.0.0.0"]
# ENTRYPOINT ["bash", "run.sh"]
ENTRYPOINT [ "python3", "test_full_cycle.py" ]