import pyautogui
import pytesseract
import cv2 
import numpy as np
import PyPDF2
from PIL import Image
import time
import os
from reportlab.lib.pagesizes import letter,LEDGER,ELEVENSEVENTEEN
from reportlab.pdfgen import canvas
from reportlab.lib.utils import ImageReader


my_dir = "./book/"
img_dir = f'{my_dir}img/'
book_name = "토비의스프링1"


# Define the top-left and bottom-right coordinates of the region to capture
left, top, right, bottom = 190-30, 40, 1660+60, 1020
width, height = right-left, bottom - top
# right, bottom = left + width, top + height

# Define the coordinates of the point to click
btn_x, btn_y = 1865, 510

def move_click_scroll(x, y):
    pyautogui.moveTo(x, y)
    time.sleep(0.5)
    pyautogui.click(x, y)
    time.sleep(0.5)
    pyautogui.moveTo(x-100, y)
    time.sleep(0.5)
    pyautogui.scroll(-100)
    time.sleep(0.5)

def screenshot_save(page:int, area:tuple):
    screenshot = pyautogui.screenshot(region=area)
    pil_image = Image.frombytes("RGB", screenshot.size, screenshot.tobytes())
    pil_image.save(f'{img_dir}{page:03d}.png')


# -30, 60
def find_best_point(page:int):
    test_cases = [-30, -20, -10, -5, 0, 5, 10, 20, 30,40,50, 60]
    for t1 in test_cases:
        left, top = 190 -30, 40
        for t2 in test_cases:
            right, bottom = 1660+t2, 1020
            width, height = right-left, bottom - top

            left_area = (left, top, width/2, height)
            screenshot_save(page, left_area)

            right_area = (left+width/2, top, width/2, height)
            screenshot_save(page+1, right_area)

            page+=2
            move_click_scroll(btn_x, btn_y)





time.sleep(1)

find_best_point(45)

# Save the captured image to a file
page = 19
while page <= 40: #878:
    # Capture the screen region as a PIL image
    left_area = (left, top, width/2, height)
    screenshot_save(page, left_area)

    right_area = (left+width/2, top, width/2, height)
    screenshot_save(page+1, right_area)

    page+=2
    move_click_scroll(btn_x, btn_y)

    
    
pdf_writer = PyPDF2.PdfWriter()

image_list = os.listdir(img_dir)
for image_file in image_list:
    img = ImageReader(f'{img_dir}{image_file}')

    pdf_canvas = canvas.Canvas(f'{my_dir}temp.pdf', pagesize=letter)
    pdf_canvas.drawImage(img, x=0, y=0, width=letter[0], height=letter[1])
    pdf_canvas.save()

    pdf_temp = open(f'{my_dir}temp.pdf', 'rb')
    pdf_temp_reader = PyPDF2.PdfReader(pdf_temp)
    pdf_writer.add_page(pdf_temp_reader.pages[0])
    pdf_temp.close()
    
    

with open(f'{my_dir}{book_name}.pdf', 'wb') as output_file:
    pdf_writer.write(output_file)


