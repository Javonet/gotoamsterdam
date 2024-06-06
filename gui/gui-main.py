import tkinter as tk
import subprocess
import os

def run_script(language):
    script_path = os.path.abspath(f'./{language}/{language}.bat')
    subprocess.call([script_path], shell=True)

app = tk.Tk()
app.title("Javonet - Speed of Integration Challenge")
app.attributes('-fullscreen', True)

button_dotnet = tk.Button(app, text=".Net Core", command=lambda: run_script('dotnet'))
button_dotnet.pack(pady=10)

button_java = tk.Button(app, text="Java", command=lambda: run_script('java'))
button_java.pack(pady=10)

button_python = tk.Button(app, text="Python", command=lambda: run_script('python'))
button_python.pack(pady=10)

button_js = tk.Button(app, text="JavaScript", command=lambda: run_script('js'))
button_js.pack(pady=10)

app.mainloop()
