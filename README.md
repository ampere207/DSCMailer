# DSC-Mailer

A Go-based email automation service for sending certificates and bulk emails with Excel parsing capabilities for Developer Students Club.

## Features

- SMTP email sending
- Excel file parsing for bulk operations
- Certificate generation
- Template-based email system
- Web interface for file uploads and email management

## Prerequisites

- Go 1.24.4 or higher
- SMTP server credentials (e.g., Gmail, SendGrid, etc.)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/vinyas-bharadwaj/DSCMailer.git
cd DSC-mailer
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
SERVER_PORT=3000

# SMTP Configuration (Required)
SMTP_USER=your-email@example.com
SMTP_PASS=your-app-password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587

# Certificate Generation (Required if using certificate features)
CERT_TEMPLATE=assets/certificate_base.png
CERT_FONT_PATH=assets/fonts/Montserrat-Bold.ttf
```

**Note:** The certificate-related variables are **required** if you plan to use the certificate generation feature. Make sure:
- `certificate_base.png` exists in the `assets/` directory
- The font file (e.g., `Montserrat-Bold.ttf`) exists in `assets/fonts/`
- If you don't need certificate generation, you can omit these variables, but certificate-related features won't work.

### SMTP Configuration Guide

**For Gmail:**
- Use `smtp.gmail.com` as SMTP_HOST
- Use port `587` for SMTP_PORT
- Generate an [App Password](https://myaccount.google.com/apppasswords) for SMTP_PASS
- Never use your actual Gmail password

**For other providers:**
- Consult your email provider's SMTP settings
- Common ports: 587 (TLS), 465 (SSL), or 25 (non-encrypted)

## Running the Application

### Option 1: Using the run script (Recommended)
```bash
chmod +x run.sh
./run.sh
```

### Option 2: Using Go directly
```bash
go run cmd/server/main.go
```

### Option 3: Build and run
```bash
go build -o dsc-mailer cmd/server/main.go
./dsc-mailer
```

The server will start on `http://127.0.0.1:3000`

## Project Structure

```
DSC-mailer/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── certificate/
│   │   └── generator.go         # Certificate generation logic
│   ├── mailer/
│   │   ├── smtp.go              # SMTP email sending
│   │   ├── template.go          # Email templates
│   │   └── worker.go            # Background worker for emails
│   ├── parser/
│   │   └── parser.go            # Excel file parsing
│   ├── server/
│   │   ├── router.go            # HTTP router setup
│   │   └── handlers/            # HTTP request handlers
│   └── utils/
│       └── create_message.go    # Email message builder
├── web/
│   ├── static/
│   │   └── styles.css           # CSS styles
│   └── templates/               # HTML templates
├── assets/
│   └── fonts/                   # Fonts for certificate generation
├── uploads/                     # Uploaded files directory
├── output/                      # Generated files directory
├── .env                         # Environment variables (create this)
└── go.mod                       # Go module dependencies
```

## Usage

1. Access the web interface at `http://127.0.0.1:3000`
2. Upload an Excel file with recipient information
3. Configure email template and subject
4. Send emails individually or in bulk

## API Endpoints

- `GET /` - Main upload page
- `POST /upload` - Handle file uploads
- Additional endpoints defined in `internal/server/router.go`

## Environment Variables

| Variable        | Required | Description                        | Example                           |
|----------------|----------|------------------------------------|-----------------------------------|
| SERVER_PORT    | No       | Server port (default: 3000)        | 3000                              |
| SMTP_USER      | Yes      | SMTP username/email                | user@gmail.com                    |
| SMTP_PASS      | Yes      | SMTP password/app token            | xxxx xxxx xxxx xxxx               |
| SMTP_HOST      | Yes      | SMTP server hostname               | smtp.gmail.com                    |
| SMTP_PORT      | Yes      | SMTP server port                   | 587                               |
| CERT_TEMPLATE  | No*      | Path to certificate template image | assets/certificate_base.png       |
| CERT_FONT_PATH | No*      | Path to font file for certificates | assets/fonts/Montserrat-Bold.ttf  |

**\*Required only if using certificate generation features**

## Troubleshooting

### "Error loading the .env file"
- Ensure `.env` file exists in the root directory
- Check file permissions

### SMTP Authentication Failed
- Verify SMTP credentials are correct
- For Gmail, ensure you're using an App Password
- Check if less secure app access is enabled (if applicable)



