import ollama 

model = "gemma2:2b"

try:
    ollama.chat(model)
except ollama.ResponseError as e:
    if e.status_code == 404:
        ollama.pull(model)
stream = ollama.chat(
    model=model,
    messages=[{'role': 'user', 'content': 'Why is the sky blue?'}],
    stream=True,
)

for chunk in stream:
  print(chunk['message']['content'], end='', flush=True)
