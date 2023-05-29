# Transaction API

This is a RESTful API built using Golang, PostgreSQL, and Docker to manage and retrieve user transactions. It provides endpoints for creating transactions, listing transactions with pagination and filtering, and retrieving transaction details.

## Prerequisites

Before running the application, make sure you have the following installed:

- Docker
- Docker Compose

## Getting Started

To start the application, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/ViniciusReno/transaction-api.git
   cd transaction-api
   ```

2. Configure the environment variables:

   - Create a `.env` file in the root directory.
   - Set the required environment variables such as database credentials and API port.

     ```
    POSTGRES_HOST=127.0.0.1
    POSTGRES_PORT=5432
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_DB=payment_user
     ```

3. Start the application:

   ```shell
   docker-compose up -d
   ```

   This command will build the Docker containers for the API and PostgreSQL database.

4. Usage

   Once the application is up and running, you can interact with it using the following endpoints:

   - **Create a Transaction**

     - URL: `POST /transactions`
     - Body: JSON payload containing transaction details
     - Example:

       ```json
       {
         "origin": "desktop-web",
         "userID": "c8e25ff6-df04-4c45-b93a-841cd76eae0a",
         "amount": 100.00,
         "operationType": "credit"
       }
       ```

   - **List Transactions**

     - URL: `GET /transactions`
     - Query Parameters:
       - `page` (optional): Page number for pagination (default: 1)
       - `limit` (optional): Number of transactions per page (default: 10)
       - `filters` (optional): JSON object with filter criteria (e.g., `{"operationType": "debit"}`)
     - Example: `GET /transactions?page=1&limit=10&filters={"operationType": "debit"}`

   - **Get Transaction Details**

     - URL: `GET /transactions/{transactionID}`
     - Example: `GET /transactions/01885e77-8cf2-58f5-abad-f6a0ddb1ed5b`


7. Troubleshooting

   If you encounter any issues while running the application or have any questions, please feel free to create an issue in the [GitHub repository](https://github.com/ViniciusReno/transaction-api/issues).

8. Contributing

   Contributions are welcome! If you'd like to contribute to this project, please fork the repository and create a pull request. Ensure that you follow the established coding guidelines and write appropriate unit tests.

9. License

   This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

The application is now up and running. You can access it at `http://localhost:8080`.

