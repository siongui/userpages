library web;

import 'dart:io';
import 'dart:async';
import 'package:path/path.dart' as path;

part 'template.dart';


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
    print('${new DateTime.now()}  '
          '${req.method} ${req.uri.path}');

    for (final String urlRegex in this.urls.keys) {
      RegExp exp = new RegExp('^${urlRegex}\$');
      Iterable<Match> matches = exp.allMatches(req.uri.path);

      //for (Match m in matches) {
      //  String match = m.group(0);
      //  print('matched: ${match}');
      //}
      if (matches.length > 0) {
        try {
          if (req.method == 'GET')
            this.urls[urlRegex].GET(req);
          if (req.method == 'POST')
            this.urls[urlRegex].POST(req);
          return;
        } on NoSuchMethodError catch(e) {
          print(e);
        }
      }
    }
    // 404 not found
    print('not found: ${req.uri.path}');
    req.response.statusCode = HttpStatus.NOT_FOUND;
    req.response.close();
  }

  void _onBindSuccess(HttpServer server) {
    info();
    server.listen(_httpRequestHandler);
  }

  void info() {
    // print information (for debugging)
    print('${new DateTime.now()}  ');
    print('server running at http://${HOST}:${PORT}/');
    print('HOST: ${HOST}');
    print('PORT: ${PORT}');
    this.urls.forEach((key, value) {
      print('${key} => ${value}');
    });
    print(this.runtimeType.toString());
    print('--- end of starting server ---');
  }

  void run() {
    HttpServer.bind(HOST, PORT).then(_onBindSuccess);
  }
}
