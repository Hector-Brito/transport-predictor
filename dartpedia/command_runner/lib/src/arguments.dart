import 'package:command_runner/src/command_runner_base.dart';
import 'dart:collection';
import 'dart:async';

enum OptionType { flag, option }

abstract class Argument {
  String get name;
  String? get help;
  Object? get defaultValue;
  String? get valueHelp;
  String get usage;
}

abstract class Command extends Argument {
  @override
  String get name;
  @override
  String? get help;
  @override
  Object? defaultValue;
  @override
  String? valueHelp;

  bool get requiresArgument => false;
  String get description;
  late CommandRunner runner;

  final List<Option> _options = [];

  UnmodifiableSetView<Option> get options =>
      UnmodifiableSetView(_options.toSet());

  void addFlag(
    int idNumber,
    String name, {
    String? help,
    String? abbr,
    String? valueHelp,
  }) {
    _options.add(
      Option(
        idNumber,
        name,
        help: help,
        abbr: abbr,
        valueHelp: valueHelp,
        type: OptionType.flag,
      ),
    );
  }

  void addOption(
    int idNumber,
    String name, {
    String? help,
    String? abbr,
    String? defaultValue,
    String? valueHelp,
  }) {
    _options.add(
      Option(
        idNumber,
        name,
        help: help,
        abbr: abbr,
        defaultValue: defaultValue,
        valueHelp: valueHelp,
        type: OptionType.option,
      ),
    );
  }

  FutureOr<Object?> run(ArgResults args);

  @override
  String get usage {
    return '$name: $description';
  }
}

class Option extends Argument {
  Option(
    this.idNumber,
    this.name, {
    required this.type,
    this.help,
    this.abbr,
    this.defaultValue,
    this.valueHelp,
  });

  @override
  final String name;

  @override
  final String? help;

  @override
  final Object? defaultValue;

  @override
  final String? valueHelp;

  @override
  String get usage {
    if (abbr != null) {
      return '-$abbr, --$name: $help';
    }
    return '--$name: $help';
  }

  final int idNumber;
  final OptionType type;
  final String? abbr;
}

class ArgResults {
  Command? command;
  String? commandArg;
  Map<Option, Object?> options = {};

  bool flag(String name) {
    for (var option in options.keys.where(
      (option) => option.type == OptionType.flag,
    )) {
      if (option.name == name) {
        return options[option] as bool;
      }
    }
    return false;
  }

  bool hasOption(String name) {
    return options.keys.any((option) => option.name == name);
  }

  ({Option option, Object? input}) getOption(String name) {
    var mapEntry = options.entries.firstWhere(
      (entry) => entry.key.name == name || entry.key.abbr == name,
    );

    return (option: mapEntry.key, input: mapEntry.value);
  }
}
