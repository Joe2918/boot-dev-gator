# Gator

## Setup

### 📋 Prerequisites

- Go](https://go.dev/doc/install) 
- [Postgres](https://www.postgresql.org/download/)

### 💾  Installation

### Via `go install` (Recommended)

With go 1.25 or higher:

```bash
go install github.com/Joe2918/boot-dev-gator@latest
```

create a config file in your home directory, ~/.gatorconfig.json, with the following content:
```
{
  "db_url": "postgres://example"
}
```

## 🚀 Usage

1. Register a user
```bash
>>> boot-dev-gator register rover
User created successfully!
 * ID:      55376e04-8fa4-4b9f-a0fb-b0ef07a331c7
 * Name:    rover
```

2. Add feed
```bash
>>> boot-dev-gator addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

3. Collect the posts from feeds
```
>>> boot-dev-gator agg 10s

```

4. List the posts
```
>>> boot-dev-gator browse 1
```





