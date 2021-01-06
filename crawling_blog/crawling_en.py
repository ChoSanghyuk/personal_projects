from urllib.request import urlopen
from bs4 import BeautifulSoup

url = ""
response = urlopen(url)
encoding = response.info().get_content_charset(failobj="utf-8")
text = response.read().decode(encoding)

soup = BeautifulSoup(text, 'html.parser')
my_data = soup.find_all("p", {"class" : "HStyle0"})

for i in my_data:
    print(i.get_text())