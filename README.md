# Gator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## 📋 Prerequisites

  Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database.

- [Go](https://go.dev/doc/install) 
- [Postgres](https://www.postgresql.org/download/)

###💾  Installation

### Via `go install` (Recommended)

With go 1.25 or higher:

```bash
>>> go install github.com/Joe2918/boot-dev-gator@latest
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## 🚀 Usage

1. Create a new user

```bash
>>> boot-dev-gator register <name>
```

2. Add a feed

```bash
>>> boot-dev-gator addfeed <name> <url>
```

3. Start the aggregator 

```
>>> boot-dev-gator agg 30s
```

4. View the posts

```
>>> boot-dev-gator browse [limit]
```

There are a few other commands you'll need as well:

- `boot-dev-gator  login <name>`  - Log in as a user that already exists
- `boot-dev-gator  users` - List all users
- `boot-dev-gator  feeds` - List all feeds
- `boot-dev-gator follow <url>` - Follow a feed that already exists in the database
- `boot-dev-gator unfollow <url>` - Unfollow a feed that already exists in the database
- `boot-dev-gator following` - List all the feeds that the current user is following