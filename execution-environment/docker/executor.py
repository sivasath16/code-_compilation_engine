import os

# Fetch the user code from an environment variable
user_code = os.getenv("CODE")

try:
    exec(user_code)
except Exception as e:
    print(f"Error: {str(e)}")

