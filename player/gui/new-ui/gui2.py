from pathlib import Path
from tkinter import Canvas, Frame, Button, PhotoImage

OUTPUT_PATH = Path(__file__).parent
ASSETS_PATH = OUTPUT_PATH / Path(r"C:\Local\GoToAmsterdam\player\gui\new-ui\assets\frame2")

def relative_to_assets(path: str) -> Path:
    return ASSETS_PATH / Path(path)

class Screen2(Frame):
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
            133.0,
            96.0,
            anchor="nw",
            text="Choose your preferred language",
            fill="#002199",
            font=("Poppins Bold", 64 * -1, "bold")
        )

        self.image_image_1 = PhotoImage(
            file=relative_to_assets("image_1.png"))
        self.image_1 = self.canvas.create_image(
            1170.0,
            35.0,
            image=self.image_image_1
        )

        self.button_image_1 = PhotoImage(
            file=relative_to_assets("button_1.png"))
        self.button_1 = Button(
            self,
            image=self.button_image_1,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("dotnet"),
            relief="flat"
        )
        self.button_1.place(
            x=39.0,
            y=357.0,
            width=263.0,
            height=87.0
        )

        self.button_image_2 = PhotoImage(
            file=relative_to_assets("button_2.png"))
        self.button_2 = Button(
            self,
            image=self.button_image_2,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("java"),
            relief="flat"
        )
        self.button_2.place(
            x=356.0,
            y=357.0,
            width=263.0,
            height=87.0
        )

        self.button_image_3 = PhotoImage(
            file=relative_to_assets("button_3.png"))
        self.button_3 = Button(
            self,
            image=self.button_image_3,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("python"),
            relief="flat"
        )
        self.button_3.place(x=356.0, y=520.0, width=263.0, height=87.0)

        self.button_image_4 = PhotoImage(file=relative_to_assets("button_4.png"))
        self.button_4 = Button(
            self,
            image=self.button_image_4,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("js"),
            relief="flat"
        )
        self.button_4.place(x=39.0, y=520.0, width=263.0, height=87.0)

        self.button_image_5 = PhotoImage(file=relative_to_assets("button_5.png"))
        self.button_5 = Button(
            self,
            image=self.button_image_5,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("ruby"),
            relief="flat"
        )
        self.button_5.place(x=982.0, y=357.0, width=263.0, height=87.0)

        self.button_image_6 = PhotoImage(file=relative_to_assets("button_6.png"))
        self.button_6 = Button(
            self,
            image=self.button_image_6,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("golang"),
            relief="flat"
        )
        self.button_6.place(x=982.0, y=520.0, width=263.0, height=87.0)

        self.button_image_7 = PhotoImage(file=relative_to_assets("button_7.png"))
        self.button_7 = Button(
            self,
            image=self.button_image_7,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("perl"),
            relief="flat"
        )
        self.button_7.place(x=673.0, y=357.0, width=263.0, height=87.0)

        self.button_image_8 = PhotoImage(file=relative_to_assets("button_8.png"))
        self.button_8 = Button(
            self,
            image=self.button_image_8,
            borderwidth=0,
            highlightthickness=0,
            command=lambda: self.set_language_and_proceed("c++"),
            relief="flat"
        )
        self.button_8.place(x=673.0, y=520.0, width=263.0, height=87.0)
        
    def set_language_and_proceed(self, language):
        self.controller.selected_language = language
        self.controller.show_frame("Screen3")