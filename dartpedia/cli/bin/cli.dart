// import 'package:cli/cli.dart' as cli;
import 'dart:io';
import 'package:http/http.dart' as http;

const version = '0.0.1';

void printUsage() {
  print("""
    Only the following commands are valid: 'help', 'version', 'search <ARTICLE-TITLE>'.
    """);
}

Future<String> getWikipediaArticle(String articleTitle) async {
  final client = http.Client();
  final url = Uri.https(
    'en.wikipedia.org',
    'api/rest_v1/page/summary/$articleTitle',
  );
  final response = await client.get(url);
  print('Ready!');
  if (response.statusCode == 200) {
    return response.body;
  }
  return 'Error: Failed to fetch article "$articleTitle". Status code: ${response.statusCode}';
}

void searchArticle(List<String>? arguments) async {
  late String? articleTitle;
  if (arguments == null || arguments.isEmpty) {
    print('Please provide an article title.');
    final inputFromStdin = stdin.readLineSync();
    if (inputFromStdin == null || inputFromStdin.isEmpty) {
      print('No article title provided.Exiting');
      return;
    }
    articleTitle = inputFromStdin;
  } else {
    articleTitle = arguments.join(' ');
  }
  print('Searching $articleTitle article on Wikipedia.');
  var articleContent = await getWikipediaArticle(articleTitle);
  print(articleContent);
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
  } else if (arguments.first == 'wikipedia') {
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
