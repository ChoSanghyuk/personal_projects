import 'package:flutter/material.dart';
import '../routes/app_routes.dart';
import 'funds.dart'; 


class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  String _currentRoute = AppRoutes.funds;

  void _onTabTapped(String route) {
    setState(() {
      _currentRoute = route;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          Navigator(
            initialRoute: _currentRoute,
            key: ValueKey(_currentRoute),
            onGenerateRoute: (settings) {
              Widget screen;
              switch (settings.name ?? _currentRoute) {
                case AppRoutes.screen1:
                  screen = const Screen1();
                  break;
                case AppRoutes.screen2:
                  screen = const Screen2();
                  break;
                case AppRoutes.funds:
                  screen = const Funds();
                  break;
                case AppRoutes.screen4:
                  screen = const Screen4();
                  break;
                case AppRoutes.screen5:
                  screen = const Screen5();
                  break;
                default:
                  screen = const Screen1();
              }
              return MaterialPageRoute(builder: (_) => screen, settings: settings);
            },
          ),
          Align(
            alignment: Alignment.bottomCenter,
            child: Container(
              decoration: BoxDecoration(
                border: Border(
                  top: BorderSide(color: Colors.blue, width: 1),
                ),
              ),
              child: BottomNavigationBar(
                type: BottomNavigationBarType.fixed,
                selectedItemColor: Colors.blue,
                unselectedItemColor: Colors.grey,
                items: [
                  BottomNavigationBarItem(
                    icon: Padding(
                      padding: EdgeInsets.all(8.0),
                      child: Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.blue, width: 1),
                          borderRadius: BorderRadius.all(Radius.circular(8)),
                        ),
                        child: Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.looks_one),
                        ),
                      ),
                    ),
                    label: 'Screen 1',
                  ),
                  BottomNavigationBarItem(
                    icon: Padding(
                      padding: EdgeInsets.all(8.0),
                      child: Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.blue, width: 1),
                          borderRadius: BorderRadius.all(Radius.circular(8)),
                        ),
                        child: Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.looks_two),
                        ),
                      ),
                    ),
                    label: 'Screen 2',
                  ),
                  BottomNavigationBarItem(
                    icon: Padding(
                      padding: EdgeInsets.all(8.0),
                      child: Container(
                        decoration: BoxDecoration(
                          color: Colors.blue.withOpacity(0.1),
                          border: Border.all(color: Colors.blue, width: 1),
                          borderRadius: BorderRadius.all(Radius.circular(8)),
                        ),
                        child: Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.home, color: Colors.blue),
                        ),
                      ),
                    ),
                    label: 'Home',
                  ),
                  BottomNavigationBarItem(
                    icon: Padding(
                      padding: EdgeInsets.all(8.0),
                      child: Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.blue, width: 1),
                          borderRadius: BorderRadius.all(Radius.circular(8)),
                        ),
                        child: Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.looks_4),
                        ),
                      ),
                    ),
                    label: 'Screen 4',
                  ),
                  BottomNavigationBarItem(
                    icon: Padding(
                      padding: EdgeInsets.all(8.0),
                      child: Container(
                        decoration: BoxDecoration(
                          border: Border.all(color: Colors.blue, width: 1),
                          borderRadius: BorderRadius.all(Radius.circular(8)),
                        ),
                        child: Padding(
                          padding: EdgeInsets.all(8.0),
                          child: Icon(Icons.looks_5),
                        ),
                      ),
                    ),
                    label: 'Screen 5',
                  ),
                ],
                onTap: (index) {
                  _onTabTapped([
                    AppRoutes.screen1,
                    AppRoutes.screen2,
                    AppRoutes.funds,
                    AppRoutes.screen4,
                    AppRoutes.screen5
                  ][index]);
                },
                currentIndex: [
                  AppRoutes.screen1,
                  AppRoutes.screen2,
                  AppRoutes.funds,
                  AppRoutes.screen4,
                  AppRoutes.screen5
                ].indexOf(_currentRoute),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class Screen1 extends StatelessWidget {
  const Screen1({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(child: Text('Screen 1'));
  }
}

class Screen2 extends StatelessWidget {
  const Screen2({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(child: Text('Screen 2'));
  }
}

class Screen4 extends StatelessWidget {
  const Screen4({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(child: Text('Screen 4'));
  }
}

class Screen5 extends StatelessWidget {
  const Screen5({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(child: Text('Screen 5'));
  }
}
