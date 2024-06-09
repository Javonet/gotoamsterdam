from pathlib import Path
from tkinter import Canvas, Frame, Button, PhotoImage

OUTPUT_PATH = Path(__file__).parent
ASSETS_PATH = OUTPUT_PATH / Path(r"C:\Local\GoToAmsterdam\player\gui\new-ui\assets\frame3")

def relative_to_assets(path: str) -> Path:
    return ASSETS_PATH / Path(path)

class Screen3(Frame):
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
        self.image_image_1 = PhotoImage(
            file=relative_to_assets("image_1.png"))
        canvas.create_image(
            923.0,
            416.0,
            image=self.image_image_1
        )

        canvas.create_text(
            68.0,
            299.0,
            anchor="nw",
            text="Using your code\nconnect to \nrobot’s dll library\nand call \nRobot.Solve()",
            fill="#000000",
            font=("Inter", 40 * -1)
        )

        self.button_image_1 = PhotoImage(
            file=relative_to_assets("button_1.png"))
        button_1 = Button(
            self,
            image=self.button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: print("button_1 clicked"),
            relief="flat"
        )
        button_1.place(
            x=205.0,
            y=640.0,
            width=291.2352294921875,
            height=174.161865234375
        )

        self.button_image_2 = PhotoImage(
            file=relative_to_assets("button_2.png"))
        button_2 = Button(
            self,
            image=self.button_image_2,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: print("button_2 clicked"),
            relief="flat"
        )
        button_2.place(
            x=640.0,
            y=640.0,
            width=291.2352294921875,
            height=174.161865234375
        )
