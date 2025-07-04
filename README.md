
# 🧩 Strata.io Maveric Identity Orchestrator – Service Extension

## 📘 Overview

This project demonstrates a custom **Service Extension** for [Strata.io Maveric Identity Orchestrator](https://strata.io) that performs a dynamic HTTP header injection. Specifically, it extracts a user’s email address from a public REST API and sets it in the request header before forwarding it to the target application.

This use case simulates real-world scenarios where dynamic attributes (e.g., emails, roles, orgs) need to be fetched and enriched into the request pipeline via Maveric.

---

## 🎯 Use Case

- 🛡️ **Goal**: Intercept user traffic and dynamically set a custom header (`CUSTOM-EMAIL`) based on external identity data.
- 🔍 **Data Source**: [JSONPlaceholder API](https://jsonplaceholder.typicode.com/users)
- 🎯 **Target**: Extract the email address of the user with `id = 2`.
- 📨 **Resulting Header**:
  ```
  CUSTOM-EMAIL: Shanna@melissa.tv
  ```

---

## 📦 Project Structure

```bash
.
├── extension/
│   └── email_header.go        # Implementation of CreateEmailHeader function
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ How It Works

1. The `CreateEmailHeader` function is executed during request processing by Maveric.
2. It fetches all users from `https://jsonplaceholder.typicode.com/users`.
3. It filters for the user with `id=2`.
4. It extracts their `email` field.
5. Sets the HTTP header: `CUSTOM-EMAIL: user.email`.
6. Returns the header map to Maveric for injection downstream.

---

## 🔐 Dependencies

- `github.com/strata-io/service-extension/orchestrator` – Orchestrator SDK (installed with Maveric)
- Standard Go libraries (`net/http`, `encoding/json`, etc.)

---

## 🧪 Sample Code Snippet

```go
header := make(http.Header)
header.Set("CUSTOM-EMAIL", email)
return header, nil
```

---

## 📡 Sample API Response

```json
[
  {
    "id": 2,
    "name": "Ervin Howell",
    "email": "Shanna@melissa.tv",
    ...
  }
]
```

---

## ✅ Testing

You can simulate a call to `CreateEmailHeader()` by wiring it into your Maveric instance or by wrapping it in a test harness. You should confirm:

- API returns valid users.
- Header is correctly set when ID=2 is found.
- Function handles HTTP/network errors gracefully.
- Logging provides enough visibility during failures.

---

## 🚀 Deployment Instructions

> Requires Maveric Identity Orchestrator already installed and configured.

1. Copy `email_header.go` into your service-extension directory.
2. Register the extension in your Maveric config:
   ```yaml
   service-extensions:
     - path: /path/to/email_header.so
   ```
3. Rebuild the extension using Go:
   ```bash
   go build -buildmode=plugin -o email_header.so email_header.go
   ```
4. Restart Maveric.

---

## 🛠️ Example Log Output

```
[INFO] se: building email custom header
[DEBUG] se: retrieving email from mock endpoint..
[INFO] se: email extracted successfully: Shanna@melissa.tv
```

---

## 📌 Notes

- You can replace the JSONPlaceholder endpoint with any RESTful identity or attribute store (e.g., Azure AD, Okta, or your internal API).
- In production, make sure to secure outbound calls and handle rate-limiting, caching, or authorization headers if needed.

---

## 📞 Support

For issues with Maveric integration or plugin development:

- Strata.io Documentation: https://docs.strata.io
- Community Slack: [Join Strata Identity Slack](https://strata.io/slack)
- Email: support@strata.io
