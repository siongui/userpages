import 'dart:io';
import 'lib/web.dart';

class index {
  void GET(HttpRequest req) {
    req.response.write('<!doctype html>'
        '<html><head></head>'
        '<body>Hello World!</body></html>');
    req.response.close();
  }
}

// simulate urls in web.py
// Set literal is not supported in Dart, so use Map instead of Set
// @see https://code.google.com/p/dart/issues/detail?id=3792
final Map<String, Object> urls = {
  '/': new index()
};


void main() {
  application app = new application(urls);
  app.run();
}
