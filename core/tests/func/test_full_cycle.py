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
    
    stage_title = "Baking"
    response = requests.post(os.path.join(API_HOST, "stage"), json={"title": stage_title, "job_id": job_id}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    stage_id = response.json()["id"]
    
    response = requests.get(os.path.join(API_HOST, "stage", str(stage_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    assert response.json()["title"] == stage_title
    
    stage_title = "Baking X"
    response = requests.patch(os.path.join(API_HOST, "stage", str(stage_id)), json={"title": stage_title}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 202
    assert response.json()["title"] == stage_title
    
    response = requests.delete(os.path.join(API_HOST, "stage", str(stage_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    
    response = requests.get(os.path.join(API_HOST, "stage", str(stage_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 403
    
    stage_title = "Baking"
    response = requests.post(os.path.join(API_HOST, "stage"), json={"title": stage_title, "job_id": job_id}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    stage1_id = response.json()["id"]
    
    stage2_title = "Delivery"
    response = requests.post(os.path.join(API_HOST, "stage"), json={"title": stage2_title, "job_id": job_id}, headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    stage2_id = response.json()["id"]
    
    # POINT
    
    point_title = "Baking Point"
    response = requests.post(os.path.join(API_HOST, "point"), 
                             json={"title": point_title, "stage_ids": [stage1_id, stage2_id]},
                             headers={"Authorization": "Bearer " + token})
    assert response.status_code == 201
    point_id = response.json()["id"]
    
    response = requests.get(os.path.join(API_HOST, "point", str(point_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    assert response.json()["title"] == point_title
    assert response.json()["stages"][0]["id"] == stage1_id
    assert response.json()["stages"][1]["id"] == stage2_id
    
    point_title = "Baking Point 2"
    response = requests.patch(os.path.join(API_HOST, "point", str(point_id)), json={"title": point_title}, headers={"Authorization": "Bearer " + token})
    print(response.json())
    assert response.status_code == 202
    assert response.json()["title"] == point_title
    assert response.json()["stages"][0]["id"] == stage1_id
    assert response.json()["stages"][1]["id"] == stage2_id
    
    response = requests.delete(os.path.join(API_HOST, "point", str(point_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 200
    
    response = requests.get(os.path.join(API_HOST, "point", str(point_id)), headers={"Authorization": "Bearer " + token})
    assert response.status_code == 403
    
    # DELETIONS

    
if __name__ == "__main__":
    pytest.main(["-c", "full_cycle.py"])
