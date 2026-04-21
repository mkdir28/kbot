# kbot
devops application from scratch
# Kbot🤖: Rock, Paper, Scissors game
A fun, interactive Rock, Paper, Scissors game built for Telegram using Go, `telebot`, and the `cobra` CLI framework. 
This guide will walk you through how to run this bot on your own computer for testing and playing.

---

## Prerequisites📋
Before you start, make sure you have the following installed on your computer:
1. **Go (Golang):** [Download and install Go](https://go.dev/doc/install) (version 1.18 or higher recommended).
2. **Telegram App:** A Telegram account on your phone or desktop.

---

## How to Run the Bot🚀

### Step 1: Create a Bot and Get a Token
To run this code, you need your own unique Telegram Bot Token.
1. Open Telegram and search for **@BotFather**.
2. Send the command `/newbot` and follow the prompts to give your bot a name and username.
3. BotFather will give you an **HTTP API Token** (it looks something like `123456789:ABCdefGhIJKlmNoPQRsTUVwxyZ`). 
4. **Copy this token.** (⚠️ *Never share this token or upload it to GitHub!*)

### Step 2: Download the Code
Clone this repository or download the source code to your computer.
```bash
# Example if using git:
git clone <your-repository-url>
cd <your-folder-name>
