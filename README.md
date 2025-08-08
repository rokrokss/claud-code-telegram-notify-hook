# Claude Code Telegram Notification Hook

Get real-time Telegram notifications when Claude Code performs actions in your projects.

## Quick Start

### 1. Create Telegram Bot

1. Message [@BotFather](https://t.me/botfather) on Telegram
2. Send `/newbot` and follow prompts
3. Save the bot token (looks like: `1234567890:ABCdefGHI...`)

### 2. Get Your Chat ID

**Option A:** Message [@userinfobot](https://t.me/userinfobot) → Get your ID

**Option B:** 
1. Message your new bot
2. Visit: `https://api.telegram.org/botYOUR_TOKEN/getUpdates`
3. Find `"chat":{"id":YOUR_ID}`

### 3. Setup

```bash
# Clone repo
git clone https://github.com/yourusername/claud-code-telegram-notify-hook.git
cd claud-code-telegram-notify-hook

# Set environment variables
export CC_HOOK_TELEGRAM_BOT_TOKEN="your_bot_token"
export CC_HOOK_TELEGRAM_CHAT_ID="your_chat_id"

# Install hooks
cp -r .claude ~/.claude
```

### 4. Test

```bash
curl -X POST "https://api.telegram.org/botYOUR_TOKEN/sendMessage" \
     -d "chat_id=YOUR_CHAT_ID" \
     -d "text=Test message"
```

## How It Works

Claude Code triggers hooks → Python script reads event → Sends formatted message to Telegram

**Message Format:**
```
🤖 Project: `my-project`
⏰ 2024-01-20 15:30:45
✅ Event: `Notification`
📌 Stop Hook Active: `false`
```

## Supported Events

- `Notification` - General Claude Code notifications
- `Stop` - Operation completion
- `SubagentStop` - Sub-agent task completion

## Troubleshooting

| Issue | Solution |
|-------|----------|
| No notifications | Check token/ID, ensure bot conversation started |
| Token errors | Copy token exactly, it's case-sensitive |
| Permission denied | Run `chmod +x ~/.claude/hooks/notification.py` |
| Group chats | Use negative chat ID (e.g., `-1001234567890`) |

## Security

- Never commit tokens to git
- Add `.env` to `.gitignore`
- Rotate tokens with BotFather's `/revoke` if compromised

## License

MIT