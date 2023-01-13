from flask import Flask, render_template, request, flash, redirect
import config, csv 
from binance.client import Client
from binance.enums import *

app = Flask(__name__)
app.secret_key= b'sdkjfhjsdk"F4Q8zec]/'

client = Client(config.API_KEY, config.API_SECRET, tld='com')


@app.route('/')
def index():
    title ='CoinView'

    account = client.get_account()
    balances= account['balances']
    exchange_info = client.get_exchange_info()
    symbols = exchange_info['symbols']
    return render_template('index.html', title=title, my_balances=balances, symbols = symbols)

@app.route('/buy', methods=['POST'])
def buy():
    print(request.form)
    try:
        order = client.create_order(
            symbol=request.form['symbol'],
            side=SIDE_BUY,
            type=ORDER_TYPE_MARKET,
            # type=ORDER_TYPE_SALE,
            timeInFroce=TIME_IN_FORCE_GTC,
            quantity=request.form['quantity'])
    except Exception as e:
        flash(e.message, "error")
    return redirect('/')

# @app.route('/buy2')
# def buy():
#     return 'buy'

@app.route('/sell')
def sell():
    return 'sell'

@app.route('/settings')
def settings():
    return 'settings'
