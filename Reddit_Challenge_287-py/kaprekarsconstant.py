
#acquires list of numbers from external file, removes spaces
def getlist():
    x = 0
    numlist = open('C:\\Users\\Matt\\Documents\\folder\\kaprekarlist.txt')
    listcontent = numlist.read()
    listarry = listcontent.split()
    for x in listarry:
        sortstuff(x)

#sorts numbers into ascending and descending for kaprekarfunction
def sortstuff(currentitem):
    z = 0
    y = 0
    fournum = []
    for y in currentitem:
            fournum.append(y)
    fournum.sort()
    revstring = fournum[0]+ fournum[1] + fournum[2] + fournum[3]
    string = fournum[3] + fournum[2] + fournum[1] + fournum[0]
    kaprekarfunction(string, revstring)

#tests whether a kaprekar solution is found
#if not, return the product back to sortstuff for another evaluation
def kaprekarfunction(big, small):
    v = 0
    print(big, small)
    if big == small:
        print('There is no solution for number ',big)
        return
    product = int(big) - int(small)
    productstring = str(product)
    if len(productstring) < 4:
        numzeros = 4 - len(productstring)
        while v < numzeros:
            productstring = '0' + str(productstring)
            v += 1
    if product == 6174:
        print('found solution')
    else:
        sortstuff(productstring)
            
getlist()
