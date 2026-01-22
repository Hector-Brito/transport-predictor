// import 'package:cli/cli.dart' as cli;
import 'dart:io';

const version = '0.0.1';

void printUsage() {
  print("""
    Only the following commands are valid: 'help', 'version', 'search <ARTICLE-TITLE>'.
    """);
}

void searchArticle(List<String>? arguments) {
  final String articleTitle;
  if (arguments == null || arguments.isEmpty) {
    print('Please provide an article title.');
    articleTitle = stdin.readLineSync() ?? '';
  } else {
    articleTitle = arguments.join(' ');
  }
  print('Searching $articleTitle article on Wikipedia.');
  print('Ready!');
}

void main(List<String> arguments) {
  //arguments es una lista de parametros
  print(
    'Welcome to Dart CLI',
  ); //que se muestran o ingresan cuando ejecutamos 'cli.dart'

  if (arguments.isEmpty || arguments.first == 'help') {
    print('Hello this is Dart CLI.');
  } else if (arguments.first == 'version') {
    print('You has enter argument "version". version of this cli is $version');
  } else if (arguments.first == 'search') {
    //la keyword 'final' es una constante cuyo valor es asignado en tiempo de ejecucion
    // Se diferencia de 'const' ya que final no ocupa espacio en memoria
    // 'final' ocupa espacio en memoria y cuando se le asigne el valor en tiempo de ejecucion.
    // 'final' usa lazy initialization
    // Es perfecta para datos que provienen de una API
    final articleTitle = arguments.length > 1 ? arguments.sublist(1) : null;
    searchArticle(articleTitle);
  } else {
    printUsage();
  }
}
