import json
import os
import sys
import urllib.request
import urllib.parse
from datetime import datetime

def send_telegram_message(bot_token, chat_id, message):
    """Send a Telegram message"""
    url = f"https://api.telegram.org/bot{bot_token}/sendMessage"
    data = {
        'chat_id': chat_id,
        'text': message,
        'parse_mode': 'Markdown'
    }
    
    try:
        req = urllib.request.Request(
            url,
            data=urllib.parse.urlencode(data).encode('utf-8'),
            headers={'Content-Type': 'application/x-www-form-urlencoded'}
        )
        with urllib.request.urlopen(req, timeout=5) as response:
            return response.status == 200
    except Exception:
        return False

def format_message(event_data, project_name):
    """Format event data into a Telegram message"""
    event = event_data.get('hook_event_name', 'Unknown')
    stop_hook_active = event_data.get('stop_hook_active', False)
    timestamp = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    
    # Base message format
    message = f"🤖 *Project: `{project_name}`*\n"
    message += f"⏰ {timestamp}\n"
    message += f"✅ Event: `{event}`\n"
    message += f"📌 Stop Hook Active: `{stop_hook_active}`"
    
    return message

def main():
    try:
        # Telegram configuration
        bot_token = os.getenv('CC_HOOK_TELEGRAM_BOT_TOKEN')
        chat_id = os.getenv('CC_HOOK_TELEGRAM_CHAT_ID')
        
        # Read input data
        input_data = json.loads(sys.stdin.read().strip())

        current_dir = os.getcwd()
        project_name = os.path.basename(current_dir)
        
        if input_data:
            # Format message
            message = format_message(input_data, project_name)
            
            # Send to Telegram
            send_telegram_message(bot_token, chat_id, message)
    
    except json.JSONDecodeError:
        sys.exit(0)
    except Exception:
        sys.exit(0)
    finally:
        sys.exit(0)

if __name__ == "__main__":
    main()
