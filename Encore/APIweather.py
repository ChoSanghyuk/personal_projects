from urllib.request import urlopen
import json
from xml.dom import minidom
import matplotlib.pyplot as plt

serviceKey ="B9%2ByG3Vqt3zhs2etzqCOJKCuKyo7LhENtjC3FMXBWqTj713QypVnIzs4pN2G%2BClL0JfSufwoshjeaj79N1%2BR2g%3D%3D"
startDt = "20100101"
endDt = "20101231"
stnIds_s = "108"
stnIds_p = "99"

url = "http://apis.data.go.kr/1360000/AsosDalyInfoService/getWthrDataList?serviceKey="+serviceKey+\
    "&numOfRows=365&pageNo=1&dataCd=ASOS&dateCd=DAY&startDt="+startDt+"&endDt="+endDt+"&stnIds="+stnIds_p

response = urlopen(url)

dom = minidom.parse(response)

f=open("pajuWeather.xml","w")
dom.writexml(f, indent='\t', addindent='\t',newl='\n')
f.close()

items = dom.getElementsByTagName("item")
print(type(items))

p_list=[]
p_graph=[]

for item in items:
    i = item.childNodes
    p_dict= {"date" :"",
    "max_wind" : "",
    "min_temperature" : ""  }
    if int(i[1].childNodes[0].data.split('-')[1]) < 3 or int(i[1].childNodes[0].data.split('-')[1])>10 :
        p_dict["date"] = i[1].childNodes[0].data
        p_dict["max_wind"] = i[14].childNodes[0].data
        p_dict["min_temperature"] = i[5].childNodes[0].data
        p_list.append(p_dict)
        p_graph.append(float(i[14].childNodes[0].data))

p_json = json.dumps(p_list, ensure_ascii=False, indent=4).encode('utf-8')

f= open("paju_json.json","wb")
f.write(p_json)
f.close()


url2 = "http://apis.data.go.kr/1360000/AsosDalyInfoService/getWthrDataList?serviceKey="+serviceKey+\
    "&numOfRows=365&pageNo=1&dataCd=ASOS&dateCd=DAY&startDt="+startDt+"&endDt="+endDt+"&stnIds="+stnIds_s


response = urlopen(url2)

dom = minidom.parse(response)

items = dom.getElementsByTagName("item")

s_list=[]
s_graph=[]

for item in items:
    i = item.childNodes
    p_dict= {"date" :"",
    "max_wind" : "",
    "min_temperature" : ""  }
    if int(i[1].childNodes[0].data.split('-')[1]) < 3 or int(i[1].childNodes[0].data.split('-')[1])>10 :       
        s_graph.append(float(i[14].childNodes[0].data))


plt.plot(p_graph,"blue")
plt.plot(s_graph,"red")
plt.show()
