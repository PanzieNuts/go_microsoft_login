## Step 1
Go to the `.env` file and set your credentials, including Microsoft:

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
##Step 2
Go to terminal and run:
```bash
go run cmd/migration/main.go
```
##Step 3
Then run:
```bash
go run cmd/api/main.go
