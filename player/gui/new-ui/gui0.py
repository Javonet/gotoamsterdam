from pathlib import Path
from tkinter import Canvas, Frame, Button, PhotoImage

OUTPUT_PATH = Path(__file__).parent
ASSETS_PATH = OUTPUT_PATH / Path(r"C:\Local\GoToAmsterdam\player\gui\new-ui\assets\frame0")

def relative_to_assets(path: str) -> Path:
    return ASSETS_PATH / Path(path)

class Screen0(Frame):
    def __init__(self, parent, controller):
        super().__init__(parent, bg="#FFFFFF")
        self.controller = controller

        self.canvas = Canvas(
            self,
            bg="#FFFFFF",
            height=832,
            width=1280,
            bd=0,
            highlightthickness=0,
            relief="ridge"
        )

        self.canvas.pack()
        self.button_image_1 = PhotoImage(
            file=relative_to_assets("button_1.png"))
        self.button_1 = Button(
            self,
            image=self.button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: controller.show_frame("Screen1"),
            relief="flat"
        )
        self.button_1.place(
            x=449.0,
            y=481.0,
            width=381.9984130859375,
            height=188.0
        )

        self.canvas.create_text(
            175.0,
            83.0,
            anchor="nw",
            text="Speed of Integration\nChallenge",
            fill="#000000",
            font=("Inter", 96 * -1)
        )