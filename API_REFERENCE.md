# API Reference

Base URL: `http://localhost:8040`

All protected endpoints require:
```
Authorization: Bearer <access_token>
```

Paginated list endpoints accept optional query params: `page` (default 1) and `page_size` (default 20, max 100).

Paginated responses follow the shape:
```json
{ "data": [...], "page": 1, "page_size": 20, "total": 42 }
```

---

## Auth

### Register
```
POST /register
```
```json
{
  "username": "john_doe",
  "password": "secret123",
  "full_name": "John Doe"
}
```
**Response 201**
```json
{
  "username": "john_doe",
  "full_name": "John Doe",
  "password_changed_at": "0001-01-01T00:00:00Z",
  "created_at": "2024-01-15T10:00:00Z"
}
```

---

### Login
```
POST /login
```
```json
{
  "username": "john_doe",
  "password": "secret123"
}
```
**Response 200**
```json
{
  "session_id": "550e8400-e29b-41d4-a716-446655440000",
  "access_token": "<jwt>",
  "access_token_expires_at": "2024-01-15T10:15:00Z",
  "refresh_token": "<jwt>",
  "refresh_token_expires_at": "2024-01-16T10:00:00Z",
  "user": {
    "username": "john_doe",
    "full_name": "John Doe",
    "password_changed_at": "0001-01-01T00:00:00Z",
    "created_at": "2024-01-15T10:00:00Z"
  }
}
```

---

### Renew Access Token
```
POST /tokens/renew_access
```
```json
{
  "refresh_token": "<refresh_jwt>"
}
```
**Response 200**
```json
{
  "access_token": "<jwt>",
  "access_token_expires_at": "2024-01-15T10:30:00Z"
}
```

---

## Own-User Routes
> Any authenticated role.

### Get Profile
```
GET /user/info
```
**Response 200**
```json
{
  "username": "john_doe",
  "full_name": "John Doe",
  "password_changed_at": "0001-01-01T00:00:00Z",
  "created_at": "2024-01-15T10:00:00Z"
}
```

---

### Update Profile
```
PUT /user/update
```
```json
{
  "full_name": "John Updated Doe"
}
```

---

### Change Password
```
POST /user/password_change
```
```json
{
  "password": "newpassword123"
}
```

---

## Tenants
> `adminUser` only.

### List Tenants
```
GET /api/v1/tenants?search=clinic&active=true&page=1&page_size=20
```
Paginated response.

---

### Create Tenant
```
POST /api/v1/tenants
```
```json
{
  "name": "Clinic Buenos Aires",
  "timezone": "America/Argentina/Buenos_Aires"
}
```
**Response 201**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Clinic Buenos Aires",
  "timezone": "America/Argentina/Buenos_Aires",
  "active": true,
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

---

### Get Tenant
```
GET /api/v1/tenants/:id
```

---

### Update Tenant
```
PUT /api/v1/tenants/:id
```
```json
{
  "name": "Clinic Buenos Aires Updated",
  "timezone": "America/Argentina/Buenos_Aires"
}
```

---

### Deactivate Tenant
```
DELETE /api/v1/tenants/:id
```

---

## Users
> `adminUser` only (except provider link/unlink which also allow `tenantUser`).

### List Users
```
GET /api/v1/users?page=1&page_size=20
```

---

### Create User (Admin)
```
POST /api/v1/users
```
```json
{
  "username": "provider_user",
  "password": "secret123",
  "full_name": "Maria Garcia",
  "role": "tenantUser",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000"
}
```
> `role` must be one of: `adminUser`, `tenantUser`, `user`.
> `tenant_id` is optional.

**Response 201**
```json
{
  "id": 7,
  "username": "provider_user",
  "full_name": "Maria Garcia",
  "role": "tenantUser",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2024-01-15T10:00:00Z"
}
```

---

### Get User
```
GET /api/v1/users/:id
```

---

### Update User Role / Tenant
```
PUT /api/v1/users/:id
```
```json
{
  "full_name": "Maria Garcia",
  "role": "user",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

---

### Delete User
```
DELETE /api/v1/users/:id
```

---

### Link User to Tenant
```
POST /api/v1/users/:id/tenant
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

---

### Link User to Provider
> `adminUser` or `tenantUser`.
```
POST /api/v1/users/:id/provider
```
```json
{
  "provider_id": "660e8400-e29b-41d4-a716-446655440001"
}
```

---

### Unlink User from Provider
> `adminUser` or `tenantUser`.
```
DELETE /api/v1/users/:id/provider?provider_id=660e8400-e29b-41d4-a716-446655440001
```

---

### List User's Providers
```
GET /api/v1/users/:id/providers
```

---

## Providers
> `adminUser` or `tenantUser`.

### List Providers
```
GET /api/v1/providers?tenant_id=550e8400-e29b-41d4-a716-446655440000&search=dr&active=true&page=1&page_size=20
```
> `tenant_id` is required.

---

### Create Provider
```
POST /api/v1/providers
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Dr. Carlos Mendez"
}
```
**Response 201**
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440001",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Dr. Carlos Mendez",
  "active": true,
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

---

### Get Provider
```
GET /api/v1/providers/:id
```

---

### Update Provider
```
PUT /api/v1/providers/:id
```
```json
{
  "name": "Dr. Carlos Mendez Jr."
}
```

---

### Deactivate Provider
```
DELETE /api/v1/providers/:id
```

---

## Provider Availability
> `adminUser` or `tenantUser`.

### Add Availability Window
```
POST /api/v1/providers/:id/availability
```
```json
{
  "weekday": 1,
  "start_time": "09:00:00",
  "end_time": "17:00:00"
}
```
> `weekday`: 0 = Sunday, 1 = Monday, …, 6 = Saturday.

**Response 201**
```json
{
  "id": "770e8400-e29b-41d4-a716-446655440002",
  "provider_id": "660e8400-e29b-41d4-a716-446655440001",
  "weekday": 1,
  "start_time": "09:00:00",
  "end_time": "17:00:00"
}
```

---

### List Availability
```
GET /api/v1/providers/:id/availability
```
**Response 200**
```json
{
  "data": [
    {
      "id": "770e8400-e29b-41d4-a716-446655440002",
      "provider_id": "660e8400-e29b-41d4-a716-446655440001",
      "weekday": 1,
      "start_time": "09:00:00",
      "end_time": "17:00:00"
    }
  ]
}
```

---

## Provider Exceptions
> `adminUser` or `tenantUser`.

### Add Exception (Block Time)
```
POST /api/v1/providers/:id/exceptions
```
```json
{
  "start_at": "2024-01-20T12:00:00Z",
  "end_at": "2024-01-20T14:00:00Z",
  "reason": "Lunch break"
}
```
> `reason` is optional.

**Response 201**
```json
{
  "id": "880e8400-e29b-41d4-a716-446655440003",
  "provider_id": "660e8400-e29b-41d4-a716-446655440001",
  "start_at": "2024-01-20T12:00:00Z",
  "end_at": "2024-01-20T14:00:00Z",
  "reason": "Lunch break",
  "created_at": "2024-01-15T10:00:00Z"
}
```

---

### List Exceptions
```
GET /api/v1/providers/:id/exceptions
```
**Response 200**
```json
{
  "data": [...]
}
```

---

## Services
> `adminUser` or `tenantUser`.

### List Services
```
GET /api/v1/services?tenant_id=550e8400-e29b-41d4-a716-446655440000&search=consult&active=true&page=1&page_size=20
```
> `tenant_id` is required.

---

### Create Service
```
POST /api/v1/services
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "General Consultation",
  "duration_minutes": 30
}
```
**Response 201**
```json
{
  "id": "990e8400-e29b-41d4-a716-446655440004",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "General Consultation",
  "duration_minutes": 30,
  "active": true,
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

---

### Get Service
```
GET /api/v1/services/:id
```

---

### Update Service
```
PUT /api/v1/services/:id
```
```json
{
  "name": "General Consultation",
  "duration_minutes": 45
}
```

---

### Deactivate Service
```
DELETE /api/v1/services/:id
```

---

## Tenant Channels
> `adminUser` or `tenantUser`.

### List Tenant Channels
```
GET /api/v1/tenant-channels?tenant_id=550e8400-e29b-41d4-a716-446655440000&channel_type=whatsapp&active=true
```
> `tenant_id` is required.

---

### Create Tenant Channel
```
POST /api/v1/tenant-channels
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "channel_type": "whatsapp",
  "external_id": "5491112345678",
  "external_key": "my-api-key"
}
```
> `external_key` is optional.

**Response 201**
```json
{
  "id": "aa0e8400-e29b-41d4-a716-446655440005",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "channel_type": "whatsapp",
  "external_id": "5491112345678",
  "external_key": "my-api-key",
  "active": true,
  "created_at": "2024-01-15T10:00:00Z"
}
```

---

### Get Tenant Channel
```
GET /api/v1/tenant-channels/:id
```

---

### Update Tenant Channel
```
PUT /api/v1/tenant-channels/:id
```
```json
{
  "channel_type": "whatsapp",
  "external_id": "5491112345678",
  "external_key": "updated-key",
  "active": true
}
```

---

### Deactivate Tenant Channel
```
DELETE /api/v1/tenant-channels/:id
```

---

## Customers
> All authenticated roles.

### List Customers
```
GET /api/v1/customers?tenant_id=550e8400-e29b-41d4-a716-446655440000&search=ana&page=1&page_size=20
```
> `tenant_id` is required.

---

### Create Customer
```
POST /api/v1/customers
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "first_name": "Ana",
  "last_name": "Lopez",
  "notes": "Prefers morning appointments"
}
```
> `first_name`, `last_name`, and `notes` are optional.

**Response 201**
```json
{
  "id": "bb0e8400-e29b-41d4-a716-446655440006",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "first_name": "Ana",
  "last_name": "Lopez",
  "notes": "Prefers morning appointments",
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

---

### Get Customer
```
GET /api/v1/customers/:id
```

---

### Update Customer
```
PUT /api/v1/customers/:id
```
```json
{
  "first_name": "Ana",
  "last_name": "Lopez Ruiz",
  "notes": "Updated notes"
}
```

---

## Slots

### Generate Slots
> All authenticated roles.
```
POST /api/v1/slot-generator
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "provider_id": "660e8400-e29b-41d4-a716-446655440001",
  "number_of_weeks": 4,
  "slot_minutes": 30,
  "timezone": "America/Argentina/Buenos_Aires"
}
```
> `slot_minutes` defaults to the service duration if omitted. Max `number_of_weeks` is 12.
> `timezone` defaults to UTC if omitted; use a valid IANA timezone string.

**Response 201**
```json
{
  "generated": 64
}
```

---

### List Available Slots
> All authenticated roles.
```
GET /api/v1/availability?tenant_id=550e8400-e29b-41d4-a716-446655440000&provider_id=660e8400-e29b-41d4-a716-446655440001&service_id=990e8400-e29b-41d4-a716-446655440004&date=2024-01-20&page=1&page_size=20
```
> All four query params are required.

**Response 200**
```json
{
  "data": [
    {
      "id": "cc0e8400-e29b-41d4-a716-446655440007",
      "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
      "provider_id": "660e8400-e29b-41d4-a716-446655440001",
      "start_at": "2024-01-20T09:00:00Z",
      "end_at": "2024-01-20T09:30:00Z",
      "status": "available",
      "created_at": "2024-01-15T10:00:00Z"
    }
  ]
}
```

---

## Appointments
> All authenticated roles.

### Create Appointment
```
POST /api/v1/appointments
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "customer_id": "bb0e8400-e29b-41d4-a716-446655440006",
  "service_id": "990e8400-e29b-41d4-a716-446655440004",
  "slot_id": "cc0e8400-e29b-41d4-a716-446655440007",
  "notes": "First visit"
}
```
> `notes` is optional.

**Response 201**
```json
{
  "id": "dd0e8400-e29b-41d4-a716-446655440008",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "customer_id": "bb0e8400-e29b-41d4-a716-446655440006",
  "provider_id": "660e8400-e29b-41d4-a716-446655440001",
  "service_id": "990e8400-e29b-41d4-a716-446655440004",
  "slot_id": "cc0e8400-e29b-41d4-a716-446655440007",
  "start_at": "2024-01-20T09:00:00Z",
  "end_at": "2024-01-20T09:30:00Z",
  "status": "confirmed",
  "notes": "First visit",
  "created_at": "2024-01-15T10:00:00Z",
  "updated_at": "2024-01-15T10:00:00Z"
}
```

---

### List Appointments
```
GET /api/v1/appointments?tenant_id=550e8400-e29b-41d4-a716-446655440000&provider_id=660e8400-e29b-41d4-a716-446655440001&customer_id=bb0e8400-e29b-41d4-a716-446655440006&status=confirmed&page=1&page_size=20
```
> `tenant_id` is required. `provider_id`, `customer_id`, and `status` are optional filters.

---

### Get Appointment
```
GET /api/v1/appointments/:id
```

---

### Update Appointment
```
PATCH /api/v1/appointments/:id
```
```json
{
  "status": "cancelled",
  "slot_id": "ee0e8400-e29b-41d4-a716-446655440009",
  "service_id": "990e8400-e29b-41d4-a716-446655440004",
  "notes": "Rescheduled"
}
```
> All fields are optional. Send only the fields you want to change.

---

### Cancel Appointment
```
DELETE /api/v1/appointments/:id
```

---

## Conversations
> `adminUser` or `tenantUser`.

### List Conversations
```
GET /api/v1/conversations?tenant_id=550e8400-e29b-41d4-a716-446655440000&page=1&page_size=20
```

---

### Get Conversation (with messages)
```
GET /api/v1/conversations/:id
```
**Response 200**
```json
{
  "id": "ff0e8400-e29b-41d4-a716-446655440010",
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "customer_id": "bb0e8400-e29b-41d4-a716-446655440006",
  "created_at": "2024-01-15T10:00:00Z",
  "messages": [
    {
      "id": "110e8400-e29b-41d4-a716-446655440011",
      "thread_id": "ff0e8400-e29b-41d4-a716-446655440010",
      "direction": "in",
      "message": "Hello, I'd like to book an appointment",
      "created_at": "2024-01-15T10:01:00Z"
    }
  ]
}
```

---

### Post Message to Conversation
```
POST /api/v1/conversations/message
```
```json
{
  "tenant_id": "550e8400-e29b-41d4-a716-446655440000",
  "customer_id": "bb0e8400-e29b-41d4-a716-446655440006",
  "direction": "out",
  "message": "Hello! Your appointment is confirmed for Jan 20 at 9am.",
  "metadata": {}
}
```
> `direction` must be `in` (customer → system) or `out` (system → customer).
> `metadata` is optional free-form JSON.

---

### Process Inbound Message
```
POST /api/v1/inbound-messages
```
```json
{
  "tenant_channel_external_id": "5491112345678",
  "external_customer_id": "5491198765432",
  "message": "Quiero un turno",
  "source": "whatsapp",
  "metadata": {}
}
```

---

## Webhooks

### Evolution / WhatsApp Webhook
> Public — no auth required. Called by Evolution API.
```
POST /api/v1/webhooks/evolution
```
```json
{
  "instance": "my-instance",
  "event": "messages.upsert",
  "data": {
    "key": {
      "remoteJid": "5491198765432@s.whatsapp.net",
      "fromMe": false,
      "id": "MSG_ID_123"
    },
    "message": {
      "conversation": "Hola, quiero un turno"
    },
    "pushName": "Ana Lopez",
    "messageType": "conversation"
  }
}
```
**Response 202** (no body)

---

## Error Responses

All errors return JSON with an `error` field:
```json
{ "error": "description of what went wrong" }
```

| Status | Meaning |
|--------|---------|
| 400 | Bad request / validation failed |
| 401 | Missing or invalid token |
| 403 | Insufficient role |
| 404 | Resource not found |
| 409 | Conflict (e.g. duplicate username) |
| 500 | Internal server error |
