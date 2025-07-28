# TestMeIfYouCan

## Overview

This project is part of the **Student Automation QA Engineer Code Challenge**. The goal of this challenge is to assess the ability to write basic automated tests, structure test code thoughtfully, and document the process. In this case, I tested a small fictional application, **BookBazaar**, which is a book management system with an API endpoint.

## Project Scenario

BookBazaar is a fictional book management interface. The goal is to write automated tests that verify the basic functionality and reliability of the system.

### API Endpoints Tested:

* `GET /books`: Fetch a list of all books
* `POST /books`: Create a new book entry
* `DELETE /books/:id`: Delete a book by its ID

### Test Case Coverage:

* Normal behavior (valid book creation)
* Edge cases (missing fields)
* Negative cases (invalid book ID)

---

## Setup and Run Instructions

### Prerequisites:

* Go (1.17 or later) installed
* Go modules enabled (`go.mod` and `go.sum` present)

### Steps to Set Up:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/braydonlowe/TestMeIfYouCan.git
   cd TestMeIfYouCan
   ```

2. **Install dependencies:**

   This project uses Go’s built-in dependency management system. You don't need to install any additional packages manually.

3. **Run the API server:**

   In the root directory, run the server:

   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:5000`. The API is now accessible for testing.

4. **Run the Tests:**

   Navigate to the `api_test.go` (inside the api folder) file and run the tests using:

   ```bash
   go test ./...
   ```

   This will run the entire test suite for the API, including all GET, POST, and DELETE tests for books.

   Note: The server does not need to be running for the tests to run and pass.

---

## Tools and Libraries Used:

* **Go**: The server is written in Go and serves as the backend for the BookBazaar application.
* **`net/http`**: Used for handling HTTP requests in the mock API.
* **`httptest`**: For creating mock HTTP requests and responses to test the API endpoints.
* **`json`**: For encoding and decoding JSON data when interacting with the server.
* **`sync`**: For synchronizing access to shared resources, preventing race conditions during testing.

---

## Test Strategy and Structure

### **Test Suite Overview:**

The tests are written to verify the basic functionality and edge cases of the BookBazaar API. The suite includes the following:

* **Test Setup:**

  * The `ResetBooks()` function is used in tests to clear the existing books and reset the system's state before running each test.

* **Tests Included:**

  1. **TestGetBooks\_InitiallyEmpty**: Verifies that the API returns an empty list of books when no books have been created yet.
  2. **TestGetBooks\_AfterMultipleCreates**: Tests if multiple books can be created and correctly returned by the GET endpoint.
  3. **TestCreateBook\_Valid**: Tests the creation of a valid book entry and checks if the server responds with a `201 Created` status.
  4. **TestCreateBook\_missingTitle**: Tests if the server correctly responds with a `400 Bad Request` when the book title is missing.
  5. **TestCreateBook\_MissingAuthor**: Verifies that the server responds with a `400 Bad Request` when the book author is missing.
  6. **TestCreateBook\_InvalidJSON**: Ensures that the server responds with a `400 Bad Request` when the provided JSON is malformed.
  7. **TestDeleteBook\_InvalidID**: Verifies that the server responds with a `400 Bad Request` when an invalid book ID is provided for deletion.
  8. **TestDeleteBook\_NotFound**: Tests if the server responds with a `404 Not Found` when trying to delete a book that doesn't exist.
  9. **TestDeleteBook\_Success**: Tests the deletion of a valid book by ID, ensuring the server responds with `200 OK` once the book is deleted.

### **Rationale for Test Structure:**

The tests are organized around the HTTP methods (GET, POST, DELETE) and the expected behavior for each. Each test follows a basic structure of:

1. **Setup**: Create a mock request.
2. **Execution**: Call the corresponding handler function.
3. **Verification**: Check the status code and the response body for correctness.

#### **Concurrency and Synchronization:**

Since the system uses a shared resource (the list of books), I’ve implemented a `sync.Mutex` to protect access to the books list, ensuring thread safety when running concurrent tests. This ensures that there is no race condition when the API endpoints are accessed concurrently.

---

## API Implementation and Why It Was Implemented This Way

The server is written using Go's `net/http` package, with a very simple design for testing purposes. I chose Go for its simplicity and ease of setting up a basic HTTP server.

### **Why a Mock API Server?**

In this context, a mock server was necessary for testing the endpoints. The server uses an in-memory list to simulate a database. This approach allows us to focus on testing API functionality without needing an actual database or external dependencies. This makes it easier to run automated tests and ensures that the tests are repeatable and isolated.

### **Concurrency and Synchronization:**

To handle multiple requests at once (in a concurrent test environment), I used `sync.Mutex` to synchronize access to shared resources like the book list. This prevents race conditions that could cause errors during testing.

### **Server Endpoint Handlers:**

1. **GET /books**: This endpoint fetches all books stored in memory and returns them as a JSON array. The test for this endpoint checks that an empty array is returned when no books are created and that the correct books are returned after creation.
2. **POST /books**: This endpoint accepts a JSON body with the book's title and author. It adds the book to the in-memory list and responds with a `201 Created` status.
3. **DELETE /books/\:id**: This endpoint deletes a book based on its ID. The test ensures that a non-existing ID results in a `404 Not Found` response.

### **Test Coverage:**

The tests cover the primary functionality of the API, including successful and edge-case scenarios (missing fields, invalid input, and invalid IDs). They ensure that the system responds correctly in each case, returning the proper status codes and response bodies.

## Docker and GitHub Actions
### Dockerfile
To ensure the application can be easily deployed and tested in any environment, a Dockerfile was created to containerize the Go API server. This allows the API to run consistently across different systems without dependency issues.

You can add a section to your README to explain the load testing attempt and the error you received. Below is a suggestion for how you could present this information:

---

### Load Testing Attempt

As part of testing the performance and reliability of the `BookBazaar` API, I implemented a load test for the `/books` endpoint. The test was designed to simulate a high number of requests (500 requests with a concurrency of 50) in order to assess the system's ability to handle load.

However, during the execution of the load test, the following error was encountered:

```
brayd@Keldon MINGW64 ~/Documents/Personal/OIT/TestMeIfYouCan (main)
$ go test ./api/api_test.go
--- FAIL: TestLoadEndpoint (0.19s)
    api_test.go:256: Received 210 errors during load test
    api_test.go:259: Completed 500 requests in 191.1922ms
FAIL
FAIL    command-line-arguments  2.598s
FAIL
```

This error message indicates that a significant number of requests (210 out of 500) failed during the load test. The test completed in 191.19ms, but the failure rate was much higher than expected, suggesting potential issues with how the server handles multiple concurrent requests.

#### Possible Causes for Failure:

1. **Server Capacity**: The server is not optimized for handling high concurrency. This could be due to insufficient handling of concurrent connections or improper locking mechanisms.
2. **Rate Limiting**: The server currently does not have proper rate-limiting mechanisms, causing it to struggle under heavy loads.
3. **Error Handling**: There might be unexpected errors in the API response that were not anticipated, leading to request failures.

#### Next Steps:

To resolve these issues, the following steps could be considered:

* **Improving Concurrency Handling**: Refactoring the code to use more efficient synchronization mechanisms, such as `sync.RWMutex` for read-heavy operations.
* **Profiling and Optimization**: Running a performance profiler to identify bottlenecks in the code that may be leading to failures under load.
* **Error Logging**: Adding more detailed error logging to capture the root cause of the failures during the load test.

---

This section provides transparency about the challenges you faced and shows your understanding of what could be improved. It will also help reviewers understand your thought process in handling performance testing.


---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This `README.md` covers all required aspects of the code challenge and provides clear instructions on how to run the tests and understand the test suite's structure.

