import pytest
import redis
import requests
import os

REDIS_HOST = os.environ.get("REDIS_HOST", "localhost")
REDIS_PORT = int(os.environ.get("REDIS_PORT", 6379))
API_HOST = os.path.join(os.environ.get("API_HOST", "http://localhost:8080"), "api", "v1")

def test_project():
    red = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True)
    username = "Mick"
    password = "Jagger"
    email = "mick@gmail.com"
    response = requests.post(os.path.join(API_HOST, "user/register"), 
                            json={ "username": username, "email": email, "password": password, "password_repeat": password }
                        )
    
    assert response.status_code == 200
    
    print(f"path is {os.path.join(API_HOST, 'user/login')}")
    response = requests.get(os.path.join(API_HOST, "user/login"), json={"username": username, "password": password})
    assert response.status_code == 200
    
if __name__ == "__main__":
    pytest.main(["-c", "full_cycle.py"])
