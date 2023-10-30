# flegetonte
This is a simple project in go for creating and email service

## Description

This project is a simple email service that uses a REST API to send emails. It uses a simple template system to create the email body and it can send emails to the recipient specified in the request body.

## Dependencies

This project uses the following dependencies:

- [gorilla/mux](https://github.com/gorilla/mux) for routing
- [gomail](https://github.com/go-gomail/gomail) for sending emails

## Requeriments

Set the environment variables in the `.env` file you can see the [`.env.example`](.env.example) file for reference.

## Usage

Run the following command to start the server:

```bash
go run main.go
```

The server will start on port 8080.

## API

### Send email

This endpoint sends an email to the recipient specified in the request body.

```bash
curl -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "email=receipt_example@gmail.com" \
  -F "subject=Test Email" \
  -F "firstName=John" \
  -F "lastName=Doe" \
  http://localhost:8080/send-email
```
