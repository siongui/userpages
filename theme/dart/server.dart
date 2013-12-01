import 'dart:io';

// simulate urls in web.py
// no set literal in Dart, so use Map instead of Set
// @see https://code.google.com/p/dart/issues/detail?id=3792
final Map<String, String> urls = {
  '/': 'index'
};

class application {
  // simulate web.application in web.py
  final String HOST = 'localhost';
  final int PORT = 8080;
  Map<String, String> urls;

  application(Map<String, String> urls) {
    this.urls = urls;
  }

  void _httpRequestHandler(HttpRequest req) {
    print('${req.method} ${req.uri.path}');
    req.response.write('Hello, world');
    req.response.close();
  }

  void _onBindSuccess(HttpServer server) {
    server.listen(_httpRequestHandler);
  }

  void info() {
    print('HOST: ${HOST}');
    print('PORT: ${PORT}');
    print('urls: ${urls}');
  }

  void run() {
    HttpServer.bind(HOST, PORT).then(_onBindSuccess);
  }
}


void main() {
  application app = new application(urls);
  app.info();
  app.run();
}
