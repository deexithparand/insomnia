/insomnia
├── /cmd
│   ├── main.go            # Entry point of the application (package main)
│   ├── setup.go           # Setup logic (package main)
│   ├── start.go           # Logic for starting the service (package main)
│   ├── stop.go            # Logic for stopping the service (package main)
│   └── status.go          # Logic for checking the status (package main)
├── /config
│   └── config.go          # Configuration management (package config)
├── /db
│   └── db.go              # Database connection & models (package db)
├── /handlers
│   ├── monitor.go         # Uptime monitoring logic (package handlers)
│   └── notifier.go        # Logic for notifications (package handlers)
├── /migrations
│   └── 001_initial_schema.sql  # Initial DB schema (SQL file, no Go package)
├── /docker
│   ├── Dockerfile         # Docker setup for the project
│   └── docker-compose.yml # Docker Compose setup for services       
├── /scripts
│   └── migration.sh       # Script for running database migrations
├── /docs
│   ├── index.md           # Main documentation file for Docsify
│   └── ...                # Additional documentation files as needed
├── /node_modules          # Automatically created by Docsify (after setup)
└── /package.json          # Docsify-related dependencies and configuration