from pathlib import Path
from tkinter import Canvas, Frame, Button, PhotoImage

OUTPUT_PATH = Path(__file__).parent
ASSETS_PATH = OUTPUT_PATH / Path(r"assets\frame1")

def relative_to_assets(path: str) -> Path:
    return ASSETS_PATH / Path(path)

class Screen1(Frame):
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
        self.canvas.place(x=0, y=0, width=1280, height=832)

        self.canvas.create_text(
            229.0,
            47.0,
            anchor="nw",
            justify="center",
            text="Welcome Brave Adventurer",
            fill="#002199",
            font=("Inter", 64 * -1, "bold")
        )

        self.canvas.create_text(
            39.0,
            147.0,
            anchor="nw",
            text="In a cyber-attacked Rubik's Cube Tower, you're the IT expert tasked with fixing \n the scrambled core using the RoboSolver. \n\n However, the challenge is to do it using foreign technology module. \n\n Thankfully you can do it with Javonet, unlock the cube’s power, and win the \n prestigious and uniqueJavonet Cube.",
            fill="#394560",
            font=("Poppins Regular", 32 * -1)
        )

        self.button_image_1 = PhotoImage(
            file=relative_to_assets("button_1.png"))
        self.button_1 = Button(
            self,
            image=self.button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: controller.show_frame("Screen2"),
            relief="flat"
        )
        self.button_1.place(
            x=505.0,
            y=597.0,
            width=270.0,
            height=82.0
        )

        self.image_image_1 = PhotoImage(
            file=relative_to_assets("image_1.png"))
        self.image_1 = self.canvas.create_image(
            1170.0,
            35.0,
            image=self.image_image_1
        )
