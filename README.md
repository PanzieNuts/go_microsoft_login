## Step 1: Set Up Your Credentials
Open the `.env` file and set your credentials, including Microsoft. Your `.env` file should look like this:

```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_db_name

CLIENT_ID=your_client_id
CLIENT_SECRET=your_client_id_secret
REDIRECT_URL=http://localhost:8000/auth/microsoft/callback
TENANT_ID=your_tenant_id
JWT_SECRET=your_jwt_secret

```

## Step 2: Run Database Migrations
In your terminal, run the following command to apply the database migrations:
```bash
go run cmd/migration/main.go

```
## Step 3: Start the API Server
Once migrations are applied, run the following to start the API server:
```bash
go run cmd/api/main.go
```
## Step 4: Authenticate via Microsoft
Open your browser and visit the following URL to trigger the Microsoft login:
```bash
http://localhost:8000/auth/microsoft
```
This should bring up the Microsoft login page for authentication.
