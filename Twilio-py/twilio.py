from twilio.rest import TwilioRestClient
accountSID = 'SecretSID'
authToken = 'SecretAuthTok'
twilioCli = TwilioRestClient(accountSID, authToken)
myTwilioNumber = 'SecretTwil#'
myCellPhone = 'SecretPersonal#'
message = twilioCli.messages.create(body='Are you okay matt, can I help you?', from_=myTwilioNumber, to_=myCellPhone)
message.to
message.from_
message.body
updatedMessage = twilioCli.messages.get(message.sid)
print(updatedMessage.status)
updatedMessage.date_sent
print('done')
