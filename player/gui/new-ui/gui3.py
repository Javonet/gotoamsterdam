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

        self.canvas = Canvas(
            self,
            bg="#FFFFFF",
            height=832,
            width=1280,
            bd=0,
            highlightthickness=0,
            relief="ridge"
        )

        #Image1
        self.canvas.place(x=0, y=0, width=1280, height=832)
        self.image_image_1 = PhotoImage(
            file=relative_to_assets("image_1.png"))
        self.canvas.create_image(
            889.0,
            305.0,
            image=self.image_image_1
        )

        # Javonet Logo
        self.image_image_2 = PhotoImage(
            file=relative_to_assets("image_2.png"))
        self.image_2 = self.canvas.create_image(
            1170.0,
            35.0,
            image=self.image_image_2
        )

        self.canvas.create_text(
            467.0,
            66.0,
            anchor="nw",
            text="Use Your Code",
            fill="#002199",
            font=("Poppins Bold", 64 * -1, "bold")
        )

        self.canvas.create_text(
            68.0,
            229.0,
            anchor="nw",
            text="Using your code connect to\nrobot’s dll library\nCall Solve static method\non a Robot class",
            fill="#000000",
            font=("Inter", 40 * -1)
        )

        #Show me Javonet magic
        self.button_image_1 = PhotoImage(
            file=relative_to_assets("button_1.png"))
        self.button_1 = Button(
            self,
            image=self.button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: print("button_1 clicked"),
            relief="flat"
        )
        self.button_1.place(
            x=215.0,
            y=507.0,
            width=391.0,
            height=161.0
        )

        #I'm Pro
        self.button_image_2 = PhotoImage(
            file=relative_to_assets("button_2.png"))
        self.button_2 = Button(
            self,
            image=self.button_image_2,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: print("button_2 clicked"),
            relief="flat"
        )
        self.button_2.place(
            x=700.0,
            y=507.0,
            width=430.0,
            height=161.0
        )
