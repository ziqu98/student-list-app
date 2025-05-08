# student-list-app

## How to Build Using Docker

1. Ensure you have Docker installed on your system.
2. Clone the repository:
    ```bash
    git clone https://github.com/ziqu98/student-list-app.git
    cd student-list-app
    ```
3. Build the Docker image:
    ```bash
    docker build -t student-list-app .
    ```
4. Run the Docker container:
    ```bash
    docker run -p 8080:8080 student-list-app
    ```

## How to Use Using cURL

1. Fetch the list of students:
    ```bash
    curl -X GET http://localhost:8080/students
    ```

2. Add a new student:
    ```bash
    curl -X POST http://localhost:8080/students -H "Content-Type: application/json" -d '{"nrp": 12345678, "name": "John Doe", "major": "Management", "gpa": 3.11}'
    ```

3. Fetch specific student based on NRP:
    ```bash
    curl -X GET http://localhost:8080/students/12345678
    ```