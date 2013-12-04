part of web;


class template {
  String templates_dir;

  // constructor
  template(String templates_dir) {
    this.templates_dir = templates_dir;
  }

  Future<String> render(String tempalte, Map<String, Object> data) {
    // render template in async way
    final File file = new File(
        path.join(this.templates_dir, tempalte));
    return file.readAsString().then((String content) {
      return processTemplateSyntax(content, data);
    });
  }

  String renderSync(String tempalte, Map<String, Object> data) {
    // render template in sync way
    final File file = new File(
        path.join(this.templates_dir, tempalte));
    return processTemplateSyntax(file.readAsStringSync(), data);
  }

  String processTemplateSyntax(String template_content,
      Map<String, Object> data) {
    // input: content in template file

    var variables = new RegExp(r'{{([^}}]+)?}}');

    /*
    for (Match match in variables.allMatches(template_content)) {
      print(match.group(0));
    }
    */

    return template_content.replaceAllMapped(variables, (Match m) {
      String var_name = m.group(1).trim();
      if (data.containsKey(var_name)) {
        return data[var_name];
      }
      return '';
    });
  }
}
