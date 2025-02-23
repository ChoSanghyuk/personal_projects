import 'dart:convert';
import 'package:flutter/services.dart';

class ConfigLoader {
  static Map<String, dynamic>? _config;

  // Load the config file only once
  static Future<void> loadConfig() async {
    if (_config == null) {
      final String data = await rootBundle.loadString('assets/config.json');
      _config = json.decode(data);
    }
  }

  // Get a value from the config file
  static String get(String key, {String defaultValue = ''}) {
    return _config?[key] ?? defaultValue;
  }

  static String getUrl() {
    return _config?["url"] ?? "";
  }
}
