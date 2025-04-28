# Folder Backup Utility

A lightweight Go utility for creating scheduled backups of any folder with automatic retention management.

## Overview

This tool creates timestamped ZIP archives of any specified folder and maintains a configurable backup history. While originally created for backing up Obsidian vaults, it works with any directory you need to preserve.

## Features

- Creates compressed ZIP backups of any folder
- Names backups with precise timestamps for easy identification
- Manages backup history automatically (configurable retention)
- Simple to configure, build, and run
- Minimal dependencies
- Cross-platform (Windows, macOS, Linux)

## Requirements

- Go 1.16 or higher

## Installation

```bash
# Clone the repository
git clone https://github.com/binge-coder/Folder-Backup-Utility.git
cd Folder-Backup-Utility

# Install dependencies (automatically fetches required packages)
go mod download
```

## Configuration

Modify these variables in [`backup-script.go`](backup-script.go) to match your requirements:

```go
sourceFolder := `full\path\to\folder\needing\backup` // Any folder you want to back up
backupFolder := `full\path\to\backup\storage`   // Where backups will be stored
```

By default, the script keeps the 4 most recent backups. To change this, modify the condition:

```go
if len(backupFiles) > 4 { // Change to your desired number of backups to keep
    // ...
}
```

## Building and Running

```bash
# Build the executable
go build -o backup-script.exe

# Run the backup
./backup-script.exe
```

## Automation

### Windows
Use Task Scheduler:
1. Open Task Scheduler
2. Create a new task
3. Set the trigger (daily/weekly/etc.)
4. Action: Start a program
5. Browse to your [`backup-script.exe`](backup-script.exe) location

### macOS/Linux
Use cron:
```bash
# Edit crontab
crontab -e

# Add a line to run daily at 2 AM, for example:
0 2 * * * /path/to/backup-script
```

## How It Works

1. Creates the backup folder if it doesn't exist
2. Generates a timestamped ZIP file of the source folder
3. Counts existing backups and removes the oldest ones if necessary

## Customization

The script uses the [mholt/archiver](https://github.com/mholt/archiver) library for creating ZIP archives. All dependencies will be automatically downloaded when you run `go mod download`.

## License

This project is open source and available under the MIT License.