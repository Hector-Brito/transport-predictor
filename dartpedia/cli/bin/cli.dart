// import 'package:cli/cli.dart' as cli;
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:command_runner/command_runner.dart';

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

Future<void> main(List<String> arguments) async {
  var commandRunner = CommandRunner()..addCommand(HelpCommand());
  commandRunner.run(arguments);
}
