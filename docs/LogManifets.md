# Logging Manifest for Project

## General Logging Rules

- Log **all important business logic** steps.
- Use **4 base log levels** consistently:

- `INFO` – Key events, process starts/stops, user actions.
- `WARNING` – Unexpected but non-fatal behavior.
- `DEBUG` – Detailed technical data for debugging.
- `ERROR` – Errors, failures, crashes.

- **Log only useful system utilities and actions** (avoid noise).
- Always log **maximum possible context for errors** (stack trace, inputs, IDs).
- Log messages must be **specific, concrete, and simple**. Avoid vague or abstract wording.

---

## Key Principles

- Time format: `UTC ISO8601 with milliseconds`.
- No logging of **passwords, secrets, tokens, sensitive user data**.
- Messages must help understand the system months later.

