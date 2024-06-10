from pathlib import Path
from tkinter import Canvas, Frame, Button, PhotoImage, Text

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

        # Image 1
        self.image_image_1 = PhotoImage(
            file=relative_to_assets("image_1.png"))
        self.image_1 = self.canvas.create_image(
            1080.0,
            531.0,
            image=self.image_image_1
        )

        # Image 2
        self.image_image_2 = PhotoImage(
            file=relative_to_assets("image_2.png"))
        self.image_2 = self.canvas.create_image(
            654.0,
            121.0,
            image=self.image_image_2
        )

        # Text
        self.canvas.create_text(
            166.0,
            218.0,
            anchor="nw",
            justify="center",
            text="Speed of Integration\nChallenge",
            fill="#002199",
            font=("Poppins Bold", 96 * -1, "bold")
        )

        # Button 1
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
            x=430.0,
            y=568.0,
            width=387.0,
            height=118.0
        )