import bs4
import requests
import smtplib

#find price of silver
res = requests.get('http://www.kitco.com/charts/livesilver.html')
res.raise_for_status()
exampleSoup = bs4.BeautifulSoup(res.text)
elems = exampleSoup.select('#sp-bid')
priceslvr = elems[0].getText()
print(priceslvr)

#find price of gold
res2 = requests.get('http://www.kitco.com/charts/livegold.html')
res2.raise_for_status()
exampleSoup = bs4.BeautifulSoup(res2.text)
elems = exampleSoup.select('#sp-bid')
pricegld = elems[0].getText()
print(pricegld)

#send email
smtpObj = smtplib.SMTP('smtp.gmail.com', 587)
print('1')
smtpObj.ehlo()
print('2')
smtpObj.starttls()
smtpObj.login('upandcomming88@gmail.com', 'PASSWORD')
print('3')
smtpObj.sendmail('upandcomming88@gmail.com','mabiesen@outlook.com','Subject: Python Automated email silver ' + priceslvr + ' and gold ' + pricegld)
smtpObj.quit()
print('done')
