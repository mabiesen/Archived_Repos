import bs4
import requests
import smtplib

#find price of silver
res = requests.get('http://www.kitco.com/charts/livesilver.html')
res.raise_for_status()
exampleSoup = bs4.BeautifulSoup(res.text, features="html.parser")
elems = exampleSoup.select('#sp-bid')
price_of_silver = elems[0].getText()
silver_price_string = "The price of silver is: " + price_of_silver
print(silver_price_string)

#find price of gold
res2 = requests.get('http://www.kitco.com/charts/livegold.html')
res2.raise_for_status()
exampleSoup = bs4.BeautifulSoup(res2.text, features="html.parser")
elems = exampleSoup.select('#sp-bid')
price_of_gold = elems[0].getText()
gold_price_string = "The price of gold is: " + price_of_gold
print(gold_price_string)

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
