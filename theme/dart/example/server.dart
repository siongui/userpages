import 'dart:io';
import 'package:webdart/web.dart';
import 'package:path/path.dart' as path;


// simulate jinja2
template tmpl_env = new template(
  path.join(
    path.dirname(Platform.script.toFilePath()),
    'views')
);


class index {
  void GET(HttpRequest req) {
    Map<String, Object> data = {
      'name': 'siongui',
      'time': new DateTime.now()
    };
    tmpl_env.render('index.html', data).then((String result) {
      req.response.write(result);
      req.response.close();
    });
  }

  void GETSync(HttpRequest req) {
    Map<String, Object> data = {
      'name': 'siongui',
      'time': new DateTime.now()
    };
    req.response.write(
        tmpl_env.renderSync('index.html', data)
    );
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
