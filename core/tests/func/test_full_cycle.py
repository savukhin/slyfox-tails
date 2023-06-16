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
    
    response = requests.get(os.path.join(API_HOST, "user/login"), json={"username": username, "password": password})
    assert response.status_code == 200
    token = response.json()["token"]

    response = requests.get(os.path.join(API_HOST, "restricted"), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    
    # PROJECT
    
    project_title = "Pizzeria"
    response = requests.post(os.path.join(API_HOST, "project"), json={"title": project_title}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    project_id = response.json()["id"]
    
    response = requests.get(os.path.join(API_HOST, "project", str(project_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    assert response.json()["title"] == project_title
    
    project_title = "Pizzeria updated"
    response = requests.patch(os.path.join(API_HOST, "project", str(project_id)), json={"title": project_title}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 202
    assert response.json()["title"] == project_title
    
    response = requests.delete(os.path.join(API_HOST, "project", str(project_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    
    response = requests.get(os.path.join(API_HOST, "project", str(project_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 403
    
    project_title = "Pizzeria"
    response = requests.post(os.path.join(API_HOST, "project"), json={"title": project_title}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    project_id = response.json()["id"]
    
    # JOB
    
    job_title = "Pepperoni"
    response = requests.post(os.path.join(API_HOST, "job"), json={"title": job_title, "project_id": project_id}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    job_id = response.json()["id"]
    
    response = requests.get(os.path.join(API_HOST, "job", str(job_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    assert response.json()["title"] == job_title
    
    job_title = "Pepperoni Mega"
    response = requests.patch(os.path.join(API_HOST, "job", str(job_id)), json={"title": job_title}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 202
    assert response.json()["title"] == job_title
    
    response = requests.delete(os.path.join(API_HOST, "job", str(job_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    
    response = requests.get(os.path.join(API_HOST, "job", str(job_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 403
    
    job_title = "Margarita"
    response = requests.post(os.path.join(API_HOST, "job"), json={"title": job_title, "project_id": project_id}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    job_id = response.json()["id"]
    
    # STAGE
    
    # STAGE EXECUTION
    
    # POINT
    
    # DELETIONS

    
if __name__ == "__main__":
    pytest.main(["-c", "full_cycle.py"])
