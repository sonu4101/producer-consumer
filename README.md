# Golang Producer-Consumer Application

A high-performance **Producer-Consumer** system built with **Go**.  
This project demonstrates:
- Controlled message production rate (RPS)
- Concurrent consumers with database writes
- Graceful shutdown with `sync.WaitGroup`
- Shared RPS limiting across producers

---

## 📂 Project Structure

```
.
├── cmd
│   └── app                # Main application entry point
├── internal
│   ├── consumer           # Consumer logic (DB writes)
│   ├── producer           # Producer logic (rate limiting)
│   └── model              # Data model for messages
└── go.mod / go.sum        # Dependencies
```

---

## ⚙️ How It Works

### **Producers**
- Multiple producer goroutines generate messages.
- A **shared rate limiter** controls the total production rate (`RPS`).
- Messages are sent into a buffered channel.

### **Consumers**
- Multiple consumer goroutines read from the channel.
- Each message is inserted into a database table.
- Uses `sync/atomic` to track the number of messages consumed.

### **Flow**
1. **Start producers** → generate messages at the configured rate.
2. **Push to channel** → channel acts as a buffer.
3. **Consumers** → read messages from the channel and insert them into DB.
4. **WaitGroups** → ensure all goroutines complete before shutdown.

---

## 🗄 Database Schema

```sql
CREATE TABLE messages (
    id INT PRIMARY KEY AUTO_INCREMENT,
    message VARCHAR(256) NOT NULL,
    created_at DATETIME NOT NULL
);
```

---

## 📦 Installation

```bash
git clone <your-repo-url>
cd <your-repo-folder>
go mod tidy
```

---

## ▶️ Running the App

```bash
go run ./cmd/app --producers=10 --consumers=5 --rps=500 --duration=10
```

### **Arguments**
| Flag        | Type   | Description                                  |
|-------------|--------|----------------------------------------------|
| `--producers` | int    | Number of producer goroutines               |
| `--consumers` | int    | Number of consumer goroutines               |
| `--rps`       | int    | Total requests per second (shared)          |
| `--duration`  | int    | Duration in seconds to run the app          |

---

## 📊 Example Output

```text
[2025-08-13T22:17:44+05:30] 🚀 Starting application...
[2025-08-13T22:17:44+05:30] 🏁 Starting consumers...
[2025-08-13T22:17:44+05:30] 🏁 Starting producers...
[2025-08-13T22:17:54+05:30] ✅ All producers finished, closing channel...

✅ Produced: 5000 | Consumed: 5000
⏱ Time taken to consume all messages: 1.002s
```

---

## 🔧 Configuration

You can change:
- **DB connection** inside `db` package.
- **Rate limiter behavior** in `producer` package.
- **Message format** inside `model` package.

---

