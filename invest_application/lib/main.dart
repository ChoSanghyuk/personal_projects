import 'package:flutter/material.dart';
import 'presentation/home.dart'; // Import your home screen
import 'data/config_loader.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized(); // Ensures async code runs before app starts
  await ConfigLoader.loadConfig(); // Load config before app runs


  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Your App Name',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const HomeScreen(), // Use your HomeScreen widget here
    );
  }
}