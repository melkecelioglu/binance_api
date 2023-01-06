import numpy
import talib
from numpy import genfromtxt

my_data = genfromtxt('15minutes.csv', delimiter=',')
print(my_data)
close= my_data[:,4]
print(close)
# close=numpy.random.random(100)

# print(close)

# moving_average = talib.SMA(close, timeperiod=10) #30 timeperiod first elements
# print(moving_average)

#relatice strenght test
rsi = talib.RSI(close)
print(rsi)


#unixtimestamp.com/index.php