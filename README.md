
# Risk Service

This simple golang service exposes REST APIs for risks

## Prerequisites

Before running the service, ensure that you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or higher)
- [Git](https://git-scm.com/)

## Getting Started

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/akraturi/RiskService.git
cd RiskService
```

### 2. Install Dependencies

The service has external dependencies on third party packages, which can be installed via `go mod`:

```bash
go mod tidy
```

### 3. Build the Service

To build the service, run the following command:

```bash
go build -o riskservice .
```

This will compile the service and create an executable file named `riskservice` in the current directory.

### 4. Running the Service

To run the service locally, execute:

```bash
./riskservice
```

Alternatively, if you prefer to run it in the background, you can use:

```bash
nohup ./riskservice &
```

This starts the service on http://localhost:8080

### 5. Test the Service

In Memory database is pre populated with some sample risks

```bash
curl http://localhost:8080/v1/risks
```

## Troubleshooting

If you encounter any issues, ensure that:

- Go is properly installed and configured.
- Port 8080 is available
- All dependencies are installed using `go mod tidy`.
