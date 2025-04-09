import os
import sys
from openai import OpenAI
from dotenv import load_dotenv

# Get user prompt from args
if len(sys.argv) < 2:
    print("Usage: python chatgpt.py [PROMPT]")
    sys.exit(1)
user_prompt = sys.argv[1]

# Read log file contents (if it exists)
log_path = "whole-repository/chatgpt_log.txt"
log_context = ""
if os.path.exists(log_path):
    with open(log_path, "r") as log_file:
        log_context = log_file.read()

# Initialize client
load_dotenv()
client = OpenAI(
    # This is the default and can be omitted
    api_key=os.environ.get("OPENAI_API_KEY"),
)

# Build the chat messages history
messages = []
if log_context:
    messages.append({
        "role": "system",
        "content": "Here is the previous conversation history:\n" + log_context
    })


# Add main system prompt
messages.append({
    "role": "system",
    "content": "You are a coding assistant that takes time to carefully read and understand each line of code."
})

# Add user prompt
messages.append({
    "role": "user",
    "content": user_prompt
})

response = client.chat.completions.create(
    model="gpt-4o",
    messages=messages,
)

output = response.choices[0].message.content
print(output)

# Append the response to a log file
with open("whole-repository/chatgpt_log.txt", "a") as log_file:
    log_file.write("\n---PROMPT---\n")
    log_file.write(user_prompt + "\n")
    log_file.write("\n---RESPONSE---\n")
    log_file.write(output + "\n")