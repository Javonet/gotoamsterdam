import tkinter as tk
from tkinter import messagebox
from PIL import Image, ImageTk
import subprocess

def run_bat_script():
    subprocess.run(["path_to_your_script.bat"])

def show_screen2():
    screen1.pack_forget()
    screen2.pack()

def show_screen3():
    screen2.pack_forget()
    screen3.pack()

def show_screen4():
    screen3.pack_forget()
    screen4.pack()

# Create main window
root = tk.Tk()
root.title("Javonet - Speed of Integration")
root.geometry("1920x1080")

# Screen 1
screen1 = tk.Frame(root)
title = tk.Label(screen1, text="Welcome to Speed of Integration Challenge", font=("Arial", 24))
title.pack(pady=20)

# Load and display image for screen 1
original_image1 = Image.open("simple-arch-dll.png")
resized_image1 = original_image1.resize((960, 540))  # Resizing to 50% of the screen width and height
photo1 = ImageTk.PhotoImage(resized_image1)

image_label1 = tk.Label(screen1, image=photo1)
image_label1.pack(pady=20)

button1 = tk.Button(screen1, text="Start", command=show_screen2)
button1.pack(pady=20)

# Screen 2
screen2 = tk.Frame(root)
label2 = tk.Label(screen2, text="Welcome Brave Adventurer", font=("Arial", 18))
label2.pack(pady=20)

button2 = tk.Button(screen2, text="Next", command=show_screen3)
button2.pack(pady=20)

# Screen 3
screen3 = tk.Frame(root)
label3 = tk.Label(screen3, text="Choose your option", font=("Arial", 18))
label3.pack(pady=20, side=tk.RIGHT)

# Load and display image for screen 3
original_image3 = Image.open("simple-arch-dll.png")
resized_image3 = original_image3.resize((960, 540))  # Resizing to 50% of the screen width and height
photo3 = ImageTk.PhotoImage(resized_image3)

image_label3 = tk.Label(screen3, image=photo3)
image_label3.pack(pady=20, side=tk.LEFT)

button3_1 = tk.Button(screen3, text="I'm PRO, I'll do it my way from scratch", command=run_bat_script)
button3_1.pack(pady=100, side=tk.LEFT)

button3_2 = tk.Button(screen3, text="I will use Javonet magic", command=show_screen4)
button3_2.pack(pady=70, side=tk.RIGHT)

# Screen 4
screen4 = tk.Frame(root)
for i in range(8):
    btn = tk.Button(screen4, text=f"Button {i+1}")
    btn.grid(row=i//4, column=i%4, padx=10, pady=10)

# Show initial screen
screen1.pack()

root.mainloop()
