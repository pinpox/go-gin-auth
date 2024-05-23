
## Run and Test

1. **Register a User:**

	```bash
	curl -X POST -H "Content-Type: application/json" -d '{"username":"yourusername", "password":"yourpassword", "name":"Your Name", "email":"your.email@example.com"}' http://localhost:8080/api/auth/register
	```

2. **Login to Get a Token:**

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"username":"yourusername", "password":"yourpassword", "name": "name", "email": "email"}' http://localhost:8080/api/auth/login
    ```

    You will receive a token in the response.

3. **Access the Profile Using the Token:**

    ```bash
    export TOKEN="your_received_token"
    curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/profile
    ```

