from flask import Flask, render_template_string
import os

app = Flask(__name__)

HTML = '''
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello Instance</title>
    <style>
        body {
            background: linear-gradient(135deg, #4f8cff 0%, #a6ffcb 100%);
            height: 100vh;
            margin: 0;
            display: flex;
            align-items: center;
            justify-content: center;
            font-family: 'Segoe UI', Arial, sans-serif;
        }
        .center {
            background: rgba(255,255,255,0.9);
            border-radius: 24px;
            box-shadow: 0 8px 32px rgba(0,0,0,0.15);
            padding: 48px 64px;
            text-align: center;
        }
        .instance {
            font-size: 3.5rem;
            font-weight: bold;
            color: #2d3a4b;
            letter-spacing: 2px;
        }
        .label {
            font-size: 1.2rem;
            color: #4f8cff;
            margin-bottom: 16px;
        }
    </style>
</head>
<body>
    <div class="center">
        <div class="instance">{{ instance }}</div>
    </div>
</body>
</html>
'''

@app.route("/")
def index():
    instance = os.environ.get("INSTANCE_NAME", "unknown")
    return render_template_string(HTML, instance=instance)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
