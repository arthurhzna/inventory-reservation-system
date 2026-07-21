# ARCHITECTURE.md

# Inventory Reservation System Architecture

## 1. Architectural Design & Synchronization

### Architecture

This project follows the principles of **Clean Architecture** to separate business logic from infrastructure concerns.

```
Presentation
    ↓
Application (Use Cases)
    ↓
Domain
    ↓
Infrastructure
```

- **Presentation** handles HTTP requests, validation, and response formatting.
- **Application** contains business rules such as reserving inventory, confirming reservations, checking stock, and expiring reservations.
- **Domain** defines entities, repository interfaces, policies, and business errors.
- **Infrastructure** implements PostgreSQL repositories, UUID generation, logging, and database transactions.

This separation makes the business logic independent from frameworks and databases, improving maintainability and testability.

---

### State Synchronization

Inventory consistency is maintained using **database transactions** and **row-level locking**.

For reservation requests:

1. Begin transaction.
2. Lock inventory row using `SELECT ... FOR UPDATE`.
3. Verify available stock.
4. Increase reserved stock.
5. Create reservation.
6. Commit transaction.

This prevents overselling when multiple users reserve the same inventory concurrently.

Reservation confirmation follows the same strategy:

1. Lock reservation row.
2. Validate reservation status.
3. Lock inventory row.
4. Decrease total stock.
5. Release reserved stock.
6. Mark reservation as confirmed.
7. Commit transaction.

Expired reservations are handled by a background worker that periodically scans expired active reservations and releases reserved stock inside a transaction.

---

### Error and Response Design

All API responses follow a consistent format.

Success

```json
{
    "message": "success",
    "data": {}
}
```

Validation Error

```json
{
    "message": "input validation error",
    "data": null,
    "errors": [
        {
            "field": "quantity",
            "message": "invalid input"
        }
    ]
}
```

Business Error

```json
{
    "message": "inventory item not found",
    "data": null
}
```

A consistent response structure simplifies frontend implementation and improves API usability.

---

# 2. Distributed Scaling & Failure Modes

The current implementation is designed to run on a **single application instance**. Inventory consistency is maintained using PostgreSQL transactions together with **row-level locking (`SELECT ... FOR UPDATE`)**, ensuring that concurrent reservation requests cannot oversell inventory.

When the application is scaled horizontally across multiple service instances, inventory consistency remains protected because database-level locking is shared across all instances connected to the same PostgreSQL database. However, scaling introduces several operational challenges.

## Failure Modes

The primary concern is the **background reservation expiration worker**. In the current implementation, every application instance starts its own worker. If the system is deployed with ten API instances, ten workers will periodically scan and attempt to process expired reservations.

Although PostgreSQL transactions and row-level locking prevent data corruption, multiple workers may still compete to process the same reservation. This does not compromise data consistency, but it introduces unnecessary database contention, duplicated work, and reduced operational efficiency.

## Proposed Improvements

To better support horizontal scaling, the API layer should remain completely **stateless**, while background processing is separated from request handling.

The reservation expiration logic should be moved into a dedicated worker service so that expiration is processed independently from the API servers. If multiple worker replicas are required for high availability, a coordination mechanism such as **leader election** or **PostgreSQL advisory locks** can ensure that only one worker processes a particular expiration task at any given time.

For future system growth, a message broker such as **RabbitMQ** or **Kafka** could be introduced to publish inventory-related events. This would allow additional services to consume inventory updates asynchronously without tightly coupling them to the API service.

With these improvements, the application can scale horizontally while preserving inventory consistency, eliminating duplicated background processing, and maintaining a stateless API architecture.
---

# 3. Engineering Trade-offs & AI Transparency

## Engineering Trade-offs

Since this assignment was intended to be completed within a **4–8 hour** timeframe, several engineering trade-offs were made to prioritize backend correctness and concurrency.

Most of the implementation effort was focused on the backend because it carries the highest weighting in the assessment. The primary objective was to ensure inventory consistency under concurrent requests by using PostgreSQL transactions together with **row-level locking (`SELECT ... FOR UPDATE`)** to prevent overselling.

For reservation expiration, a simple background worker was implemented to periodically scan and expire outdated reservations. This approach satisfies the assignment requirements while avoiding the additional complexity of a distributed scheduler or external job processing system.

On the frontend, the application was intentionally implemented as a lightweight single-page React application using React's built-in state management. Additional libraries such as Redux and React Router were intentionally omitted because they did not provide significant value for the scope of the assignment.

Testing efforts were primarily focused on the application's business logic (use cases), as this layer contains the core inventory reservation rules. Comprehensive integration and end-to-end testing were considered outside the available implementation time and would be appropriate for future development.

These decisions kept the implementation simple, maintainable, and aligned with the primary objective of delivering a correct and concurrent inventory reservation system within the given time constraints.

---

## AI Transparency

ChatGPT was used as a development assistant for discussing implementation approaches, reviewing code, and improving project documentation.

One example occurred during the implementation of the **reservation confirmation** workflow. ChatGPT suggested several implementation approaches that were functionally correct but did not fully align with the transaction boundaries and separation of responsibilities adopted in this project. The implementation was subsequently refined to remain consistent with the **Unit of Work** pattern, preserve clear transaction boundaries, and maintain the separation between business logic and the persistence layer in accordance with the project's **Clean Architecture**.

This experience reinforced that AI was used as a productivity tool for exploring implementation alternatives, while all architectural decisions, implementation validation, and final testing remained engineering responsibilities throughout the development process.