from urllib.request import urlopen
import sys
from bs4 import BeautifulSoup
import re
import json

f = urlopen('https://www.weather.go.kr/weather/forecast/mid-term_01.jsp')

print(f.getheader('Content-Type'))
encoding = f.info().get_content_charset(failobj="utf-8")
text= f.read().decode(encoding)


# print(text)

soup = BeautifulSoup(text, 'html.parser')
my_data = soup.select(
   '#content_weather  table:nth-child(9)  tbody  tr:nth-child(1)'
)

# data_cl = my_data[0].group()
# data_cl = re.sub('(<([^>]+)>)', '', str(my_data[0]) )
data_cl = my_data[0].text.strip()

# print(data_cl)
# print(type(data_cl))

data_li = data_cl.split('\n')

print(type(data_li))

data_li.remove('서울ㆍ인천ㆍ경기도')
data_li.remove('그래프')


while '' in data_li:
    data_li.remove('')

print(data_li)

my_data2 = soup.select(
   '#content_weather > table:nth-child(9) > thead'
)
data_cl2 = my_data2[0].text.strip()
data2_li = data_cl2.split('\n')
print(data2_li)

weath = {}

for i in range(0,8) :
       weath[data2_li[i+2]] = data_li[i+1]
print(weath)

seo_weath={}
seo_weath[data_li[0]] = weath

print(seo_weath)

output_name = "weather\seo_weather.json"
output = open(output_name, "w", encoding = "utf-8", newline="")


print(json.dumps(seo_weath, ensure_ascii=False, indent=4), file = output)