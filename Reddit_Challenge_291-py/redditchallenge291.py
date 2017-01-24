goldilocks = open('locationoffilegoeshere') #for windows users, format is C:\\folder\\...\\file.extension
goldicontent = goldilocks.read()
goldiray = goldicontent.split()
optionlen = len(goldiray)
goldilist = []

# ctr starts at 2 due to goldilocks, incremented by 2 due to one chair
#having two values
#chairctr starts at 1 due to goldilocks
#two items of the array (index items 0 and 1)
chairctr = 1
ctr = 2

while ctr < len(goldiray):
    goldiweight = int(goldiray[0])
    golditemp = int(goldiray[1])
    weight = int(goldiray[ctr])
    temp = int(goldiray[ctr+1])
    if (weight > goldiweight) and (temp < golditemp):
        goldilist.append(chairctr)
    ctr += 2
    chairctr += 1

print(goldilist)
