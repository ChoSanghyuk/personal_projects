import pyautogui
import sys
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


# 기본 폴더, 이미지를 저장할 폴더, 책의 이름을 지정합니다.
my_dir = "./book/"
img_dir = f'{my_dir}img/'
book_name = "토비의스프링2"


# 마우스 트레이서라는 프로그램을 통해, 전자책의 사각형 꼭지점의 좌표들을 기록합니다.
left, top, right, bottom = 190-30, 40, 1660+60, 1020
width, height = right-left, bottom - top


# 전자책의 넘기는 버튼의 좌표를 기록합니다.
btn_x, btn_y = 1865, 510

# 전자책의 시작, 끝 페이지 기록
start_page, end_page = 1, 838

# 전자책을 넘기고 스크롤 하는 함수를 작성합니다.
# 전자책을 최대한 확대후 캡처하기 때문에, 옆 페이지로 넘긴 후에 한 스크롤 내리는 모션 추가합니다.
def move_click_scroll(x, y):
    pyautogui.moveTo(x, y)
    time.sleep(0.5)
    pyautogui.click(x, y)
    time.sleep(0.5)
    pyautogui.moveTo(x-100, y)
    time.sleep(0.5)
    pyautogui.scroll(-100)
    time.sleep(0.5)

# area를 캡처하고 이미지 파일로 저장하는 함수를 작성합니다.
def screenshot_save(page:int, area:tuple):
    screenshot = pyautogui.screenshot(region=area)
    pil_image = Image.frombytes("RGB", screenshot.size, screenshot.tobytes())
    pil_image.save(f'{img_dir}{page:03d}.png')


# 최적의 캡처 위치를 찾습니다.
# 보통 가로의 폭의 지정이 어렵기 때문에, for 문을 돌면서 left와 right를 조정해가며 캡처합니다.
# 캡처의 결과물을 보고 left와 right를 재정한 후 전체 캡처를 진행합니다.
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


# 프로그램 시작 버튼 누른 후, ebook을 화면에 띄울 때까지의 시간을 벌기 위함
time.sleep(1)

# find_best_point(45)

# 캡처를 진행합니다.
# 교보 ebook의 경우, 단일 페이지로 할 경우, 프레임 안에서 페이지들의 위치가 달라지기 때문에 일정한 형태로 캡처하기 어렵습니다.
# 이에 양쪽 페이지로 펼친 후, 한쪽면씩 캡처합니다.
page = start_page
while page <= end_page: #878:
    # Capture the screen region as a PIL image
    left_area = (left, top, width/2, height)
    screenshot_save(page, left_area)

    right_area = (left+width/2, top, width/2, height)
    screenshot_save(page+1, right_area)

    page+=2
    move_click_scroll(btn_x, btn_y)


# 저장한 이미지들을 pdf 파일로 저장합니다.
# 여러장을 할 경우, 중간에 blank page가 들어가는 현상이 존재합니다.
# 이로 인해 현재 아래 기능은 사용하지 않고, https://tools.pdf24.org/ko/images-to-pdf 페이지에서 변환 작업을 진행합니다.
pdf_writer = PyPDF2.PdfWriter()


image_list = os.listdir(img_dir)
for image_file in image_list:
    img = ImageReader(f'{img_dir}{image_file}')

    pdf_canvas = canvas.Canvas(f'{my_dir}temp.pdf', pagesize=letter)
    pdf_canvas.drawImage(img, x=0, y=0, width=letter[0], height=letter[1])
    pdf_canvas.save()

    pdf_temp = open(f'{my_dir}temp.pdf', 'rb')
    print(f"The size of the {image_file} file is {os.path.getsize(f'{my_dir}temp.pdf')} bytes.")
    pdf_temp_reader = PyPDF2.PdfReader(pdf_temp)
    pdf_temp_reader.pages[0].scale
    print(f"The size of the {image_file} file is f'{sys.getsizeof(pdf_temp_reader.pages[0])} bytes.")
    
    pdf_writer.add_page(pdf_temp_reader.pages[0])
    pdf_temp.close()


with open(f'{my_dir}{book_name}.pdf', 'wb') as output_file:
    pdf_writer.write(output_file)


