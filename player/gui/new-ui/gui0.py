from pathlib import Path
from tkinter import Tk, Canvas, Button, PhotoImage, Frame

OUTPUT_PATH = Path(__file__).parent
ASSETS_PATH_0 = OUTPUT_PATH / Path(r"C:\Local\GoToAmsterdam\player\gui\new-ui\build\assets\frame0")
ASSETS_PATH_1 = OUTPUT_PATH / Path(r"C:\Local\GoToAmsterdam\player\gui\new-ui\build\assets\frame1")

def relative_to_assets(path: str, frame: int) -> Path:
    if frame == 0:
        return ASSETS_PATH_0 / Path(path)
    else:
        return ASSETS_PATH_1 / Path(path)

class App(Tk):
    def __init__(self):
        super().__init__()
        self.title("Multi-Screen App")
        self.geometry("1280x832")
        self.configure(bg="#FFFFFF")
        self.frames = {}
        
        for F in (Screen1, Screen2):
            page_name = F.__name__
            frame = F(parent=self, controller=self)
            self.frames[page_name] = frame
            frame.grid(row=0, column=0, sticky="nsew")
        
        self.show_frame("Screen1")
    
    def show_frame(self, page_name):
        frame = self.frames[page_name]
        frame.tkraise()

class Screen1(Frame):
    def __init__(self, parent, controller):
        super().__init__(parent, bg="#FFFFFF")
        self.controller = controller

        canvas = Canvas(
            self,
            bg="#FFFFFF",
            height=832,
            width=1280,
            bd=0,
            highlightthickness=0,
            relief="ridge"
        )
        canvas.place(x=0, y=0)
        button_image_1 = PhotoImage(file=relative_to_assets("button_1.png", 0))
        button_1 = Button(
            self,
            image=button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: controller.show_frame("Screen2"),
            relief="flat"
        )
        button_1.image = button_image_1  # Keep a reference to avoid garbage collection
        button_1.place(x=449.0, y=481.0, width=381.9984130859375, height=188.0)

        canvas.create_text(
            175.0,
            83.0,
            anchor="nw",
            text="Speed of Integration\nChallenge",
            fill="#000000",
            font=("Inter", 96 * -1)
        )

class Screen2(Frame):
    def __init__(self, parent, controller):
        super().__init__(parent, bg="#FFFFFF")
        self.controller = controller

        canvas = Canvas(
            self,
            bg="#FFFFFF",
            height=832,
            width=1280,
            bd=0,
            highlightthickness=0,
            relief="ridge"
        )
        canvas.place(x=0, y=0)
        canvas.create_text(
            229.0,
            114.0,
            anchor="nw",
            text="Welcome Brave Adventurer",
            fill="#000000",
            font=("Inter", 64 * -1)
        )

        button_image_1 = PhotoImage(file=relative_to_assets("button_1.png", 1))
        button_1 = Button(
            self,
            image=button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: controller.show_frame("Screen1"),
            relief="flat"
        )
        button_1.image = button_image_1  # Keep a reference to avoid garbage collection
        button_1.place(x=494.0, y=613.0, width=291.2352294921875, height=77.33251953125)

if __name__ == "__main__":
    app = App()
    app.mainloop()
