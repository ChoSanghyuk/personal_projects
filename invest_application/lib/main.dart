import 'package:flutter/material.dart';
import 'data/config_loader.dart';
import 'presentation/auth.dart'; // Import your home screen // check login 제거
import 'presentation/home.dart'; // Import your home screen
// import 'services/notifications_service.dart';

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
      initialRoute: '/', 
      
      // Define named routes
      routes: {
        '/': (context) => const AuthWrapper(),
        '/home': (context) => const HomeScreen(),
        '/login': (context) => const LoginScreen(),
      },

      // Fallback route handler
      onUnknownRoute: (settings) {
        return MaterialPageRoute(
          builder: (context) => Scaffold(
            appBar: AppBar(title: const Text('Not Found')),
            body: const Center(child: Text('Page not found')),
          ),
        );
      },
    );
  }
}

/*
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Popup Alarm Example',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: HomeScreen(),
    );
  }
}

class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Popup Alarm Example")),
      body: Center(
        child: ElevatedButton(
          onPressed: () {
            showPopupMessage(context); // Show dialog
          },
          child: Text("Show Alarm"),
        ),
      ),
    );
  }
}

void showPopupMessage(BuildContext context) {
  showDialog(
    context: context,
    builder: (BuildContext context) {
      return AlertDialog(
        title: Text("Alarm"),
        content: Text("This is your alarm message!"),
        actions: [
          TextButton(
            onPressed: () {
              Navigator.of(context).pop(); // Close dialog
            },
            child: Text("OK"),
          ),
        ],
      );
    },
  );
}
*/


