import 'dart:io';

class index {
  String name = 'index';

  void GET() {
  }
}

// simulate urls in web.py
// Set literal is not supported in Dart, so use Map instead of Set
// @see https://code.google.com/p/dart/issues/detail?id=3792
final Map<String, Object> urls = {
  '/': new index()
};

class application {
  // simulate web.application in web.py
  final String HOST = 'localhost';
  final int PORT = 8080;
  Map<String, Object> urls;

  // constructor
  application(Map<String, Object> urls) {
    this.urls = urls;
  }

  void _httpRequestHandler(HttpRequest req) {
    print('${req.method} ${req.uri.path}');
    req.response.write('Hello, world');
    req.response.close();
  }

  void _onBindSuccess(HttpServer server) {
    info();
    server.listen(_httpRequestHandler);
  }

  void info() {
    print('server running at http://${HOST}:${PORT}/');
    print('HOST: ${HOST}');
    print('PORT: ${PORT}');
    this.urls.forEach((key, value) {
      print('${key} => ${value}');
    });
  }

  void run() {
    HttpServer.bind(HOST, PORT).then(_onBindSuccess);
  }
}


void main() {
  application app = new application(urls);
  app.run();
  print(app.runtimeType.toString());
}
