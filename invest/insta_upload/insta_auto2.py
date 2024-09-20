from instabot import Bot
from PIL import Image
from PIL.ExifTags import TAGS
import cv2
import numpy as np
import os

# 이미지 사진에서 날짜 데이터 불러옴
def get_info(img_name):
    img = Image.open(path_ori + img_name )
    exifdata = img.getexif()
    info_target = 'DateTime'
    for tag_id in exifdata:
        tag = TAGS.get(tag_id, tag_id)
        if tag != info_target:
            continue
        info = exifdata.get(tag_id)
        if isinstance(info, bytes):
            info = info.decode()
        return info

# 이미지 불러옴(한글 경로 가능)
def imread(filename, flags=cv2.IMREAD_COLOR, dtype=np.uint8):
    try: 
        n = np.fromfile(filename, dtype) 
        img = cv2.imdecode(n, flags) 
        return img 
    except Exception as e: 
        print(e) 
        return None

# 변경 후 이미지 저장(한글 경로 가능)
def imwrite(filename, img, params=None): 
    try: 
        ext = os.path.splitext(filename)[1] 
        result, n = cv2.imencode(ext, img, params) 
        if result: 
            with open(filename, mode='w+b') as f: 
                n.tofile(f) 
            return True 
        else: 
            return False 
    except Exception as e: 
        print(e) 
        return False

# 세로 사진의 경우 90도 회전
def rotate_image(img, angle):
    rows = img.shape[0]
    cols = img.shape[1]
    img_center = (cols / 2, rows / 2)
    M = cv2.getRotationMatrix2D(img_center, angle, 1)
    rotated_image = cv2.warpAffine(img, M, (cols, rows), borderValue=(255,255,255))
    return rotated_image

def resize(img_name, scale_percent):
    path = path_ori + img_name

    src = imread(path)
    width = int(src.shape[1]*scale_percent /100 )
    height = int(src.shape[0] * scale_percent /100)
    dsize = (width, height)

    output = cv2.resize(src, dsize)
    # 이미지를 회전 시키려면 이미지 이름이 r로 시작해야 함
    if img_name[0] == 'r':
        output = rotate_image(output, -90)
    imwrite(path_resize + img_name, output)


path_ori = 'C:/Users/Cho sanghyuk/Pictures/insta/original/'
path_resize = 'C:/Users/Cho sanghyuk/Pictures/insta/resize/'
temp = os.listdir(path_ori)

listing = [[i , get_info(i) ] for i in temp ]
listing = [ i for i in listing if isinstance(i[1],str) and i[1].strip(' ') != '' ]
listing = sorted(listing , key= lambda x : x[1])


bot = Bot()
# 인스타 아이디와 패스워드 입력
bot.login(username = '*****', password = '*****')

count = 122
except_li = []
for i in listing[122:]:
    print(count)
    resize(i[0], 50)
    # 사진 업로드 및 사진의 날짜를 caption으로
    bot.upload_photo(path_resize+i[0] , caption = i[1][:10].replace(':', '-'))
    
    count +=1
