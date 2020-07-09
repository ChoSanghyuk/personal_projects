import csv
import matplotlib.pyplot as plt
from matplotlib import font_manager, rc
import platform

f = open('2020년 06월  교통카드 통계자료.csv', encoding='euc-kr')
data = csv.reader(f)

paid = []
unpaid = []
station = []
legend_list = ["유임승차", "무임승차"]
for row in data:
    if data.line_num ==1:
        continue
    if row[1] != '1호선':
        continue
    for i in range(4,8):
        row[i] = int(row[i].replace(',',''))
    paid.append(row[4])
    unpaid.append(row[6])
    station.append(row[3])


print(station)

font_name = font_manager.FontProperties(fname="c:/Windows/Fonts/malgun.ttf").\
    get_name()
rc('font', family = font_name)

plt.bar(range(len(station)), paid, width=-0.4,align='edge')
plt.bar(range(len(station)), unpaid, width=0.4,align='edge')

plt.ylim(0, 1000000)
plt.xticks(range(len(station)), station, rotation=10)
plt.legend(legend_list)
plt.show()
