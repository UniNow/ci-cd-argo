from locust import HttpUser, task


class CICDDemoUser(HttpUser):
    @task
    def hello_world(self):
        self.client.get("/")
        self.client.get("/contact")
        self.client.get("/users")
        self.client.get("/api/users")
