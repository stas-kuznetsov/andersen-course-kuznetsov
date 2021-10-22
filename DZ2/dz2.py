from flask import Flask
from flask import request
from flask import url_for
from flask import json
import emoji

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def dz2():
	if request.method == 'GET':
		return 'Hello. This site created by Stanislav Kuznetsov 😀 '
	else:
		request_data = request.get_json(force=True)

		if "animal" in request_data:
			animal = request_data["animal"]

		if 'sound' in request_data:
			sound = request_data["sound"]

		if  'count' in request_data:
			count = request_data["count"]

		i=0
		list = ''

		while i < count:
			list = list + animal + ' says ' + sound + '\n'
			i += 1

		pict = emoji.emojize(f':{animal}:')
		list = list + 'Made with ' + pict  + ' by Stanislav Kuznetsov \n'

		return list




if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)




