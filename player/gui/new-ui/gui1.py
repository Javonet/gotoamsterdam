from pathlib import Path
from tkinter import Canvas, Frame, Button, PhotoImage

OUTPUT_PATH = Path(__file__).parent
ASSETS_PATH = OUTPUT_PATH / Path(r"C:\Local\GoToAmsterdam\player\gui\new-ui\assets\frame1")

def relative_to_assets(path: str) -> Path:
    return ASSETS_PATH / Path(path)

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

        canvas.place(x=0, y=0, width=1280, height=832)
        canvas.create_text(
            229.0,
            114.0,
            anchor="nw",
            text="Welcome Brave Adventurer",
            fill="#000000",
            font=("Inter", 64 * -1)
        )

        self.button_image_1 = PhotoImage(
            file=relative_to_assets("button_1.png"))
        button_1 = Button(
            self,
            image=self.button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: controller.show_frame("Screen2"),
            relief="flat"
        )
        button_1.place(
            x=494.0,
            y=413.0,
            width=291.2352294921875,
            height=77.33251953125
        )
