import pyautogui
import pytesseract
import cv2 
import numpy as np
import PyPDF2
from PIL import Image
import time
import os
from reportlab.lib.pagesizes import letter
from reportlab.pdfgen import canvas
from reportlab.lib.utils import ImageReader

my_dir = "./book/"
img_dir = f'{my_dir}img/'
book_name = "토비의스프링1"


# Define the top-left and bottom-right coordinates of the region to capture
left, top, right, bottom = 550, 90, 1350, 1070
width, height = right-left, bottom - top
right, bottom = left + width, top + height

screenshot = pyautogui.screenshot(region=(left, top, width, height))
screenshot = np.array(screenshot)
screenshot = screenshot[top:top+height, left:left+width, :]

# Convert the screenshot to grayscale and apply thresholding
gray = cv2.cvtColor(screenshot, cv2.COLOR_BGR2GRAY)
thresh = cv2.threshold(gray, 0, 255, cv2.THRESH_BINARY_INV | cv2.THRESH_OTSU)[1]

# Perform OCR on the thresholded image
text = pytesseract.image_to_string(thresh)

# Get the coordinates of the bounding box around the text
boxes = pytesseract.image_to_boxes(thresh)
coords = boxes.split('\n')
x1, y1, x2, y2 = map(int, coords[0].split()[1:5])

# Calculate the coordinates of the rectangular part of the screen around the text
text_area_left = left + x1
text_area_top = top + (height - y2)
text_area_width = x2 - x1
text_area_height = y2

# Get the rectangular part of the screen around the text
text_area_image = pyautogui.screenshot(region=(text_area_left, text_area_top, text_area_width, text_area_height))

# Save the image
text_area_image.save(f'{img_dir}/text_area.png')



# Define the coordinates of the point to click
x, y = 1866, 510


# Save the captured image to a file
# page = 20
# while page <= 878:
#     time.sleep(1)
#     # Capture the screen region as a PIL image
#     screenshot = pyautogui.screenshot(region=(left, top, width, height))
#     pil_image = Image.frombytes("RGB", screenshot.size, screenshot.tobytes())
#     pil_image.save(f'{img_dir}{page:03d}.png')

#     # pyautogui.moveTo(x, y)
#     pyautogui.click(x, y)
#     page+=1




# pdf_writer = PyPDF2.PdfWriter()


# image_list = os.listdir(img_dir)
# for image_file in image_list:
#     image = ImageReader(f'{img_dir}{image_file}')
#     pdf_canvas = canvas.Canvas(f'{my_dir}temp.pdf', pagesize=letter)
#     pdf_canvas.drawImage(image, x=0, y=0, width=letter[0], height=letter[1])
#     pdf_canvas.save()

#     pdf_temp = open(f'{my_dir}temp.pdf', 'rb')
#     pdf_temp_reader = PyPDF2.PdfReader(pdf_temp)
#     pdf_writer.add_page(pdf_temp_reader.pages[0])
#     pdf_temp.close()

# with open(f'{my_dir}{book_name}', 'wb') as output_file:
#     pdf_writer.write(output_file)

