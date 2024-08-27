from flask import Flask, Response, request, jsonify
import time

from flask_cors import CORS

app = Flask(__name__)
CORS(app)
clients = {}


def event_stream(client_id):
    while True:
        if client_id not in clients:
            break
        yield f"data: Message for client {client_id} at {time.strftime('%Y-%m-%d %H:%M:%S')}\n\n"
        time.sleep(5)


@app.route('/event/subscribe', methods=['GET'])
def subscribe():
    client_id = request.args.get('clientId')
    if not client_id:
        return jsonify({"error": "clientId is required"}), 400

    def generate():
        yield from event_stream(client_id)

    response = Response(generate(), content_type='text/event-stream')
    response.headers['Cache-Control'] = 'no-cache'
    response.headers['Connection'] = 'keep-alive'
    response.headers['Access-Control-Allow-Origin'] = '*'
    clients[client_id] = response
    return response


@app.route('/event/unsubscribe', methods=['GET'])
def unsubscribe():
    client_id = request.args.get('clientId')
    if not client_id:
        return jsonify({"error": "clientId is required"}), 400

    if client_id in clients:
        del clients[client_id]
        return jsonify({"message": f"Client {client_id} unsubscribed"}), 200
    else:
        return jsonify({"error": "Client not found"}), 404


if __name__ == '__main__':
    app.run(debug=True, threaded=True, port=8080)
