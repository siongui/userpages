import 'dart:io';

final String HOST = 'localhost';
final int PORT = 8080;

void _httpRequestHandler(HttpRequest req) {
  print('${req.method} ${req.uri.path}');
  req.response.write('Hello, world');
  req.response.close();
}

void _onBindSuccess(HttpServer server) {
  server.listen(_httpRequestHandler);
}

void main() {
  HttpServer.bind(HOST, PORT).then(_onBindSuccess);
}
