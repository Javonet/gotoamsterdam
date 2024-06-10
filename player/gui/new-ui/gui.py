from tkinter import Tk
from gui0 import Screen0
from gui1 import Screen1
from gui2 import Screen2
from gui3 import Screen3
from gui3_python import Screen3Python

class App(Tk):
    def __init__(self):
        super().__init__()
        self.title("Javonet - Speed of Integration")
        #self.attributes("-fullscreen", True)
        self.geometry("1920x1080")
        self.configure(bg="#AAAAAA")
        self.frames = {}
        self.selected_language = "default" 
        
        for F in (Screen0, Screen1, Screen2, Screen3, Screen3Python):
            page_name = F.__name__
            frame = F(parent=self, controller=self)
            self.frames[page_name] = frame
            frame.grid(row=0, column=0, sticky="nsew")
        
        self.show_frame("Screen0")
    
    def show_frame(self, page_name):
        frame = self.frames[page_name]
        frame.tkraise()

if __name__ == "__main__":
    app = App()
    app.mainloop()
