
# Risk Service API Documentation

## Overview
This API allows you to manage risks by performing CRUD operations (Create, Read, Update, Delete). You can fetch all risks, retrieve a specific risk by ID, or add a new risk.

### Base URL
```
http://localhost:8080
```

---

## Endpoints

### 1. Get All Risks

#### **Request:**
```
GET http://localhost:8080/v1/risks
```

#### **cURL Example:**
```bash
curl -X GET "http://localhost:8080/v1/risks"
```

#### **Response:**
This will return a list of all risks in the system.

**Example Response:**
```json
{
  "data": [
    { "id": "2025-01-21T095937.200", "title": "Risk1", "state": "open", "description": "Description1" },
    { "id": "2025-01-21T095816.200", "title": "Risk2", "state": "closed", "description": "Description2" },
    ...
  ]
}
```

---

### 2. Post a New Risk

#### **Request:**
```
POST http://localhost:8080/v1/risks
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "Risk4",
  "state": "open",
  "description": "Description4"
}
```

#### **cURL Example:**
```bash
curl -X POST "http://localhost:8080/v1/risks" -H "Content-Type: application/json" -d '{
  "title": "Risk4",
  "state": "open",
  "description": "Description4"
}'
```

#### **Response:**
This will create a new risk and return the created risk data.

**Example Response:**
```json
{
  "data": {
    "id": "2025-01-21T100838.201",
    "title": "Risk4",
    "state": "open",
    "description": "Description4"
  }
}
```

---

### 3. Get Risk by ID

#### **Request:**
```
GET http://localhost:8080/v1/risks/{id}
```

**Example Request URL:**
```
GET http://localhost:8080/v1/risks/01efd757-e97c-66ed-bfbf-544810d07baf
```

#### **cURL Example:**
```bash
curl -X GET "http://localhost:8080/v1/risks/01efd757-e97c-66ed-bfbf-544810d07baf"
```

#### **Response:**
This will return a specific risk based on the provided ID.

**Example Response:**
```json
{
  "data": {
    "id": "01efd757-e97c-66ed-bfbf-544810d07baf",
    "title": "Risk1",
    "state": "open",
    "description": "Description1"
  }
}
```

If the risk is not found, the API will return a `404 Not Found` error.

**Example Error Response:**
```json
{
  "error": "Risk not found"
}
```

---

## Error Codes

- `404`: Not Found – The requested risk does not exist.
- `500`: Internal Server Error – Something went wrong on the server.

---

## Example Use Cases

1. **Retrieve all risks:**
   - Send a `GET` request to `/v1/risks` to get a list of all risks in the system.
   - **cURL Command:**
     ```bash
     curl -X GET "http://localhost:8080/v1/risks"
     ```

2. **Add a new risk:**
   - Send a `POST` request to `/v1/risks` with a JSON body to create a new risk.
   - **cURL Command:**
     ```bash
     curl -X POST "http://localhost:8080/v1/risks" -H "Content-Type: application/json" -d '{
       "title": "Risk4",
       "state": "open",
       "description": "Description4"
     }'
     ```

3. **Get risk by ID:**
   - Send a `GET` request with the risk ID to retrieve a specific risk.
   - **cURL Command:**
     ```bash
     curl -X GET "http://localhost:8080/v1/risks/01efd757-e97c-66ed-bfbf-544810d07baf"
     ```

---